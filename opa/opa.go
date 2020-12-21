// Package opa provides client make requests to the Open Policy Agent API
package opa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type OpaClient struct {
	logger *zap.Logger
	Host   string
}

type OpaEvalutePolicyRequest struct {
	Input json.RawMessage `json:"input"`
}

type OpaEvaluatePolicyResponse struct {
	Result *OpaEvaluatePolicyResult `json:result`
}
type OpaEvaluatePolicyResult struct {
	Pass       bool `json:pass`
	Violations []*OpaEvaluatePolicyViolation
}

type OpaEvaluatePolicyViolation struct {
	Message string `json:message`
}

func NewOPA(logger *zap.Logger, host string) *OpaClient {
	client := &OpaClient{
		logger: logger,
		Host:   host,
	}
	return client
}

func (opa *OpaClient) InitializePolicy(policy string) error {
	return nil
}

func (opa *OpaClient) EvaluatePolicy(policy string, input string) (*OpaEvaluatePolicyResult, error) {
	log := opa.logger.Named("Evalute Policy")
	request, err := json.Marshal(&OpaEvalutePolicyRequest{Input: json.RawMessage(input)})
	if err != nil {
		log.Error("failed to encode OPA input", zap.Error(err), zap.String("input", input))
		return nil, fmt.Errorf("failed to encode OPA input: %s", err)
	}
	response, err := http.Post(opa.getURL(fmt.Sprintf("v1/data/%s", policy)), "application/json", bytes.NewReader(request))
	if err != nil {
		log.Error("http request to OPA failed", zap.Error(err))
		return nil, fmt.Errorf("http request to OPA failed: %s", err)
	}

	res := &OpaEvaluatePolicyResponse{}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Error("failed to decode OPA result", zap.Error(err))
		return nil, fmt.Errorf("failed to decode OPA result: %s", err)
	}

	return res.Result, nil
}

func (opa *OpaClient) getURL(path string) string {
	return fmt.Sprintf("%s/%s", opa.Host, path)
}
