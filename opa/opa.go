// Copyright 2021 The Rode Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package opa provides client make requests to the Open Policy Agent API
package opa

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "github.com/rode/rode/proto/v1alpha1"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

//go:generate counterfeiter -generate

// Client is an interface for sending requests to the OPA API
//counterfeiter:generate . Client
type Client interface {
	InitializePolicy(policy string, policyData string) ClientError
	EvaluatePolicy(policy string, input []byte) (*EvaluatePolicyResponse, error)
}

type client struct {
	logger       *zap.Logger
	Host         string
	ExplainQuery bool
	httpClient   *http.Client
}

// EvalutePolicyRequest OPA evalute policy request
type EvalutePolicyRequest struct {
	Input json.RawMessage `json:"input"`
}

// EvaluatePolicyResponse OPA evaluate policy response
type EvaluatePolicyResponse struct {
	Result      *EvaluatePolicyResult `json:"result"`
	Explanation *[]string             `json:"explanation"`
}

// EvaluatePolicyResult OPA evaluate policy result
type EvaluatePolicyResult struct {
	Pass       bool                          `json:"pass"`
	Violations []*pb.EvaluatePolicyViolation `json:"violations"`
}

// PolicyViolation Rego rule conditions
type PolicyViolation struct {
	Conditions []byte
}

// Write Rego rule to IO writer
func (v *PolicyViolation) Write(w io.Writer) {
	w.Write([]byte("violation[] {\n"))
	w.Write(v.Conditions)
	w.Write([]byte("\n}\n\n"))
}

// NewClient OpaClient constructor
func NewClient(logger *zap.Logger, host string, explainQuery bool) Client {
	// 5 second timeout on all http calls in opa
	httpclient := &http.Client{
		Timeout: 5 * time.Second,
	}
	client := &client{
		logger:       logger,
		Host:         host,
		ExplainQuery: explainQuery,
		httpClient:   httpclient,
	}
	return client
}

// InitializePolicy initializes OPA policy if it does not already exist
func (opa *client) InitializePolicy(policy string, policyData string) ClientError {
	_ = opa.logger.Named("Initialize Policy")

	exists, err := opa.policyExists(policy)
	if err != nil {
		return NewClientError("error checking if policy exists", OpaClientErrorTypeGetPolicy, err)
	}

	if !exists {
		err := opa.loadPolicy(policy, policyData)
		if err != nil {
			return NewClientError(err.Error(), OpaClientErrorTypeLoadPolicy, err)
		}
	}

	return nil
}

func (opa *client) loadPolicy(policy string, policyData string) error {
	log := opa.logger.Named("Load Policy")

	req, err := http.NewRequest(http.MethodPut, opa.getURL(fmt.Sprintf("v1/policies/%s", policy)), strings.NewReader(policyData))
	if err != nil {
		return fmt.Errorf("failed to generate policy request")
	}
	response, err := opa.httpClient.Do(req)
	if err != nil {
		log.Error("load policy error", zap.Error(err), zap.String("input", policyData))
		return fmt.Errorf("failed to load the policy in opa")
	}

	if response.StatusCode == http.StatusOK {
		log.Debug("successfully loaded policy in opa")
	} else {
		return fmt.Errorf("failed to load the policy in opa, got %d", response.StatusCode)
	}

	return nil
}

// EvaluatePolicy evaluates OPA policy against provided input
func (opa *client) EvaluatePolicy(policy string, input []byte) (*EvaluatePolicyResponse, error) {
	log := opa.logger.Named("Evalute Policy")

	request, err := json.Marshal(&EvalutePolicyRequest{Input: input})
	if err != nil {
		log.Error("failed to encode OPA input", zap.Error(err), zap.String("input", string(input)))
		return nil, fmt.Errorf("failed to encode OPA input: %s", err)
	}

	httpResponse, err := opa.httpClient.Post(opa.getDataQueryURL(getOpaPackagePath(policy)), "application/json", bytes.NewReader(request))
	if err != nil {
		log.Error("http request to OPA failed", zap.Error(err))
		return nil, fmt.Errorf("http request to OPA failed: %s", err)
	}
	if httpResponse.StatusCode != http.StatusOK {
		log.Error("http response status from OPA not OK", zap.Any("status", httpResponse.Status))
		return nil, fmt.Errorf("http response status not OK")
	}

	response := &EvaluatePolicyResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		log.Error("failed to decode OPA result", zap.Error(err))
		return nil, fmt.Errorf("failed to decode OPA result: %s", err)
	}

	return response, nil
}

// policyExists tests if OPA policy exists
func (opa *client) policyExists(policy string) (bool, error) {
	log := opa.logger.Named("Policy Exists")
	response, err := opa.httpClient.Get(opa.getURL(fmt.Sprintf("v1/policies/%s", policy)))
	if err != nil {
		log.Error("error sending get policy request to OPA", zap.Error(err))
		return false, NewClientError("error sending get policy request to OPA", OpaClientErrorTypeHTTP, err)
	}
	switch response.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		log.Error("unhandled get policy response from OPA", zap.Any("status", response.Status))
		return false, NewClientError(fmt.Sprintf("unhandled get policy response from OPA: %s", response.Status), OpaClientErrorTypeBadResponse, nil)
	}
}

// publishPolicy publishes attester violation rules to OPA policy
func (opa *client) publishPolicy(policy string, violations []PolicyViolation) error {
	log := opa.logger.Named("Publish Policy")
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("package %s\n\n", policy))
	buf.WriteString("pass = true { count(violation) == 0 }\n\n")
	for _, violation := range violations {
		violation.Write(buf)
	}
	request, err := http.NewRequest(http.MethodPut, opa.getURL(fmt.Sprintf("v1/policies/%s", policy)), buf)
	if err != nil {
		log.Error("error creating create policy request", zap.Error(err))
	}
	response, err := opa.httpClient.Do(request)
	if err != nil {
		log.Error("error sending create policy request", zap.Error(err))
		return NewClientError("error sending create policy request", OpaClientErrorTypeHTTP, err)
	}
	if response.StatusCode != http.StatusOK {
		log.Error("unhandled create policy response status", zap.Any("status", response.Status))
		return NewClientError("unhandled create policy resposne status", OpaClientErrorTypeBadResponse, nil)
	}
	return nil
}

// getURL for given OPA API path
func (opa *client) getURL(path string) string {
	return fmt.Sprintf("%s/%s", opa.Host, path)
}

func (opa *client) getDataQueryURL(path string) string {
	var query string
	if opa.ExplainQuery {
		query = "explain=full&pretty"
	} else {
		query = ""
	}
	return opa.getURL(fmt.Sprintf("v1/data/%s?%s", path, query))
}

func getOpaPackagePath(regoContent string) string {
	// TODO implement better regex in the event that the package line is not the first line
	// suggested regex: ^\s*package.*$
	// ex: package abc.def.ghi
	re := regexp.MustCompile(`(?m)^[[:blank:]]*(package.*$)`)
	packageLine := re.FindAllStringSubmatch(string(regoContent), -1)[0][1]

	// package name will be abc.def.ghi
	packageName := strings.Split(packageLine, " ")[1]

	// abc/def/ghi
	return strings.Replace(packageName, ".", "/", -1)
}
