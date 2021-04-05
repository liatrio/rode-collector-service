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

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/google/uuid"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/esutil"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	"github.com/rode/rode/opa"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"

	"github.com/golang/protobuf/proto"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"github.com/open-policy-agent/opa/ast"
	"github.com/rode/rode/config"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	rodeElasticsearchOccurrencesAlias   = "grafeas-rode-occurrences"
	rodeElasticsearchPoliciesIndex      = "rode-v1alpha1-policies"
	rodeElasticsearchResourceNamesIndex = "rode-v1alpha1-resource-names"
	maxPageSize                         = 1000
)

// NewRodeServer constructor for rodeServer
func NewRodeServer(
	logger *zap.Logger,
	grafeasCommon grafeas_proto.GrafeasV1Beta1Client,
	grafeasProjects grafeas_project_proto.ProjectsClient,
	opa opa.Client,
	esClient *elasticsearch.Client,
	filterer filtering.Filterer,
	elasticsearchConfig *config.ElasticsearchConfig,
) (pb.RodeServer, error) {
	rodeServer := &rodeServer{
		logger:              logger,
		grafeasCommon:       grafeasCommon,
		grafeasProjects:     grafeasProjects,
		opa:                 opa,
		esClient:            esClient,
		filterer:            filterer,
		elasticsearchConfig: elasticsearchConfig,
	}
	if err := rodeServer.initialize(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to initialize rode server: %s", err)
	}

	return rodeServer, nil
}

type rodeServer struct {
	pb.UnimplementedRodeServer
	logger              *zap.Logger
	esClient            *elasticsearch.Client
	filterer            filtering.Filterer
	grafeasCommon       grafeas_proto.GrafeasV1Beta1Client
	grafeasProjects     grafeas_project_proto.ProjectsClient
	opa                 opa.Client
	elasticsearchConfig *config.ElasticsearchConfig
}

type docsResponse struct {
	Docs []esMGetResponse `json:"docs"`
}

type esBulkQueryFragment struct {
	Create *esBulkQueryCreateFragment `json:"create"`
}

type esBulkQueryCreateFragment struct {
	Id string `json:"_id"`
}

type esBulkResponse struct {
	Items  []*esBulkResponseItem `json:"items"`
	Errors bool				     `json:"errors"`
}

type esBulkResponseItem struct {
	Create *esCreateDocResponse `json:"create,omitempty"`
}

type esCreateDocResponse struct {
	Id      string           `json:"_id"`
	Result  string           `json:"result"`
	Version int              `json:"_version"`
	Status  int              `json:"status"`
	Error   *esCreateDocError `json:"error,omitempty"`
}

type esCreateDocError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

func (r *rodeServer) BatchCreateOccurrences(ctx context.Context, occurrenceRequest *pb.BatchCreateOccurrencesRequest) (*pb.BatchCreateOccurrencesResponse, error) {
	log := r.logger.Named("BatchCreateOccurrences")
	log.Debug("received request", zap.Any("BatchCreateOccurrencesRequest", occurrenceRequest))
	uriMap := make(map[string]bool)
	for _, x := range occurrenceRequest.Occurrences {
		resourceName := strings.Split(x.Resource.Uri, "@")[0]
		uriMap[resourceName] = true
	}
	var resNames []string
	for resName := range uriMap {
		resNames = append(resNames, resName)
	}

	docsReqBody, value := encodeRequest(map[string]interface{}{
		"ids": resNames,
	})

	log.Info("printing value of encodeRequest", zap.String("value", string(value)))
	getDocsRes, err := r.esClient.Mget(docsReqBody, r.esClient.Mget.WithContext(ctx), r.esClient.Mget.WithIndex(rodeElasticsearchResourceNamesIndex))

	if err != nil {
		log.Error("failed to get documents", zap.NamedError("error", err))
		return nil, err
	}
	if getDocsRes.IsError() {
		return nil, fmt.Errorf("Unexpected status code while getting documents: %d", getDocsRes.StatusCode)
	}
	docs := docsResponse{}
	if err := decodeResponse(getDocsRes.Body, &docs); err != nil {
		return nil, err
	}
	fmt.Printf("docs: %v", docs)



	var body bytes.Buffer
	for resName := range uriMap {
		docIdCheck := false

		resBody := map[string]string{
			"resourceName": resName,
		}
		fmt.Printf("JSON string: %s", value)
		for _, docIds := range docs.Docs {

			if docIds.Found && docIds.ID == resName {
				log.Info("in comparison ", zap.String("resName", string(resName)))
				log.Info("in comparison ", zap.String("docId.ID", string(docIds.ID)))
				docIdCheck = true
			}
		}
		if docIdCheck {
			log.Info("skipping resource", zap.String("resName", string(resName)))
			continue
		}

		createMetadata := &esBulkQueryFragment{
			Create: &esBulkQueryCreateFragment{
				Id: resName,
			},
		}

		metadata, _ := json.Marshal(createMetadata)
		metadata = append(metadata, "\n"...)

		name, _ := json.Marshal(resBody)
		dataBytes := append(name, "\n"...)

		body.Grow(len(metadata) + len(dataBytes))
		body.Write(metadata)
		body.Write(dataBytes)
	}
	if body.Len() > 0 {
		createDocRes, createDocErr := r.esClient.Bulk(
			bytes.NewReader(body.Bytes()),
			r.esClient.Bulk.WithContext(ctx),
			r.esClient.Bulk.WithRefresh(string(r.elasticsearchConfig.Refresh.String())),
			r.esClient.Bulk.WithIndex(rodeElasticsearchResourceNamesIndex))

		if createDocErr != nil {
			log.Error("failed to create documents", zap.NamedError("error", createDocErr))
			return nil, createDocErr
		}

		if createDocRes.IsError() {
			body, err := ioutil.ReadAll(createDocRes.Body)
			if err != nil {
				return nil, err
			}
			log.Info("The response body", zap.String("body", string(body)))
			return nil, fmt.Errorf("Unexpected status code while creating documents: %d", createDocRes.StatusCode)
		}

		response := &esBulkResponse{}
		err = esutil.DecodeResponse(createDocRes.Body, response)
		if err != nil {
			return nil, err
		}
		var errors []error
		if response.Errors {
			for _, item := range response.Items {
				if item.Create.Error != nil && item.Create.Status != http.StatusConflict {
					itemError := fmt.Errorf( "[%d] %s: %s", item.Create.Status, item.Create.Error.Type, item.Create.Error.Reason)
					log.Error("Error creating resources", zap.Error(itemError))
					errors = append(errors, itemError)
				}
			}
			if len(errors) > 0 {
				return nil, status.Errorf(codes.Internal, "Failed to create all resources: %v", errors)
			}
		}

		//fmt.Printf("bulk response: %v", response)
	}
	//Forward to grafeas to create occurrence
	occurrenceResponse, err := r.grafeasCommon.BatchCreateOccurrences(ctx, &grafeas_proto.BatchCreateOccurrencesRequest{
		Parent:      "projects/rode",
		Occurrences: occurrenceRequest.GetOccurrences(),
	})
	if err != nil {
		log.Error("failed to create occurrences", zap.NamedError("error", err))
		return nil, err
	}

	return &pb.BatchCreateOccurrencesResponse{
		Occurrences: occurrenceResponse.GetOccurrences(),
	}, nil
}

func (r *rodeServer) EvaluatePolicy(ctx context.Context, request *pb.EvaluatePolicyRequest) (*pb.EvaluatePolicyResponse, error) {
	var err error
	log := r.logger.Named("EvaluatePolicy").With(zap.String("policy", request.Policy), zap.String("resource", request.ResourceURI))
	log.Debug("evaluate policy request received")

	// check OPA policy has been loaded
	err = r.opa.InitializePolicy(request.Policy)
	if err != nil {
		log.Error("error checking if policy exists", zap.Error(err))
		return nil, status.Error(codes.Internal, "check if policy exists failed")
	}

	// fetch occurrences from grafeas
	listOccurrencesResponse, err := r.grafeasCommon.ListOccurrences(ctx, &grafeas_proto.ListOccurrencesRequest{Parent: "projects/rode", Filter: fmt.Sprintf(`"resource.uri" == "%s"`, request.ResourceURI)})
	if err != nil {
		log.Error("list occurrences failed", zap.Error(err), zap.String("resource", request.ResourceURI))
		return nil, status.Error(codes.Internal, "list occurrences failed")
	}
	log.Debug("Occurrences found", zap.Any("occurrences", listOccurrencesResponse))

	// json encode occurrences. list occurrences response should not generate error
	input, _ := protojson.Marshal(proto.MessageV2(listOccurrencesResponse))

	// evalute OPA policy
	evaluatePolicyResponse, err := r.opa.EvaluatePolicy(request.Policy, input)
	if err != nil {
		log.Error("evaluate OPA policy failed")
		return nil, status.Error(codes.Internal, "evaluate OPA policy failed")
	}
	log.Debug("Evalute policy result", zap.Any("policy result", evaluatePolicyResponse))

	attestation := &pb.EvaluatePolicyResult{}
	attestation.Created = timestamppb.Now()
	attestation.Pass = evaluatePolicyResponse.Result.Pass
	for _, violation := range evaluatePolicyResponse.Result.Violations {
		attestation.Violations = append(attestation.Violations, &pb.EvaluatePolicyViolation{
			Id:          violation.ID,
			Name:        violation.Name,
			Description: violation.Description,
			Message:     violation.Message,
			Link:        violation.Link,
			Pass:        violation.Pass,
		})
	}

	return &pb.EvaluatePolicyResponse{
		Pass: evaluatePolicyResponse.Result.Pass,
		Result: []*pb.EvaluatePolicyResult{
			attestation,
		},
		Explanation: *evaluatePolicyResponse.Explanation,
	}, nil
}

func (r *rodeServer) ListResources(ctx context.Context, request *pb.ListResourcesRequest) (*pb.ListResourcesResponse, error) {
	log := r.logger.Named("ListResources")
	log.Debug("received request", zap.Any("ListResourcesRequest", request))

	searchQuery := esSearch{
		Collapse: &esCollapse{
			Field: "resource.uri",
		},
	}

	if request.Filter != "" {
		parsedQuery, err := r.filterer.ParseExpression(request.Filter)
		if err != nil {
			log.Error("failed to parse query", zap.Error(err))
			return nil, err
		}

		searchQuery.Query = parsedQuery
	}

	encodedBody, requestJSON := encodeRequest(searchQuery)
	log.Debug("es request payload", zap.Any("payload", requestJSON))
	res, err := r.esClient.Search(
		r.esClient.Search.WithContext(ctx),
		r.esClient.Search.WithIndex(rodeElasticsearchOccurrencesAlias),
		r.esClient.Search.WithBody(encodedBody),
		r.esClient.Search.WithSize(maxPageSize),
	)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("error occurred during ES query %v", res)
	}

	var searchResults esSearchResponse
	if err := decodeResponse(res.Body, &searchResults); err != nil {
		return nil, err
	}
	var resources []*grafeas_proto.Resource
	for _, hit := range searchResults.Hits.Hits {
		occurrence := &grafeas_proto.Occurrence{}
		err := protojson.Unmarshal(hit.Source, proto.MessageV2(occurrence))
		if err != nil {
			log.Error("failed to convert", zap.Error(err))
			return nil, err
		}

		resources = append(resources, occurrence.Resource)
	}

	return &pb.ListResourcesResponse{
		Resources:     resources,
		NextPageToken: "",
	}, nil
}

func (r *rodeServer) ListGenericResources(ctx context.Context, request *pb.ListGenericResourcesRequest) (*pb.ListGenericResourcesResponse, error) {
	log := r.logger.Named("ListGenericResources")
	log.Debug("received request", zap.Any("request", request))

	searchQuery := esSearch{}

	encodedBody, requestJSON := encodeRequest(searchQuery)
	log.Debug("es request payload", zap.Any("payload", requestJSON))
	res, err := r.esClient.Search(
		r.esClient.Search.WithContext(ctx),
		r.esClient.Search.WithIndex(rodeElasticsearchResourceNamesIndex),
		r.esClient.Search.WithBody(encodedBody),
		r.esClient.Search.WithSize(maxPageSize),
	)

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("error occurred during ES query %v", res)
	}

	var searchResults esSearchResponse
	if err := decodeResponse(res.Body, &searchResults); err != nil {
		return nil, err
	}
	var resourceNames []string
	for _, hit := range searchResults.Hits.Hits {
		resourceNames = append(resourceNames, hit.ID)
	}

	return &pb.ListGenericResourcesResponse{
		ResourceNames: resourceNames,
		NextPageToken: "",
	}, nil
}

func encodeRequest(body interface{}) (io.Reader, string) {
	b, err := json.Marshal(body)
	if err != nil {
		// we should know that `body` is a serializable struct before invoking `encodeRequest`
		panic(err)
	}

	return bytes.NewReader(b), string(b)
}

func decodeResponse(r io.ReadCloser, i interface{}) error {
	return json.NewDecoder(r).Decode(i)
}

func (r *rodeServer) initialize(ctx context.Context) error {
	log := r.logger.Named("initialize")

	_, err := r.grafeasProjects.GetProject(ctx, &grafeas_project_proto.GetProjectRequest{Name: "projects/rode"})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err := r.grafeasProjects.CreateProject(ctx, &grafeas_project_proto.CreateProjectRequest{Project: &grafeas_project_proto.Project{Name: "projects/rode"}})
			if err != nil {
				log.Error("failed to create rode project", zap.Error(err))
				return err
			}
			log.Info("created rode project")
		} else {
			log.Error("error checking if rode project exists", zap.Error(err))
			return err
		}
	}
	// Create an index for policy storage
	r.esClient.Indices.Create(rodeElasticsearchPoliciesIndex)
	// Create an index for resource names
	r.esClient.Indices.Create(rodeElasticsearchResourceNamesIndex)

	return nil
}

func (r *rodeServer) ListOccurrences(ctx context.Context, occurrenceRequest *pb.ListOccurrencesRequest) (*pb.ListOccurrencesResponse, error) {
	log := r.logger.Named("ListOccurrences")
	log.Debug("received request", zap.Any("ListOccurrencesRequest", occurrenceRequest))

	requestedFilter := occurrenceRequest.Filter

	listOccurrencesResponse, err := r.grafeasCommon.ListOccurrences(ctx, &grafeas_proto.ListOccurrencesRequest{Parent: "projects/rode", Filter: requestedFilter})
	if err != nil {
		log.Error("list occurrences failed", zap.Error(err), zap.String("filter", occurrenceRequest.Filter))
		return nil, status.Error(codes.Internal, "list occurrences failed")
	}
	return &pb.ListOccurrencesResponse{
		Occurrences:   listOccurrencesResponse.GetOccurrences(),
		NextPageToken: "",
	}, nil
}

func (r *rodeServer) UpdateOccurrence(ctx context.Context, occurrenceRequest *pb.UpdateOccurrenceRequest) (*grafeas_proto.Occurrence, error) {
	log := r.logger.Named("UpdateOccurrence")
	log.Debug("received request", zap.Any("UpdateOccurrenceRequest", occurrenceRequest))

	name := fmt.Sprintf("projects/rode/occurrences/%s", occurrenceRequest.Id)

	if occurrenceRequest.Occurrence.Name != name {
		log.Error("Occurrence name does not contain the occurrence id", zap.String("occurrenceName", occurrenceRequest.Occurrence.Name), zap.String("id", occurrenceRequest.Id))
		return nil, status.Error(codes.InvalidArgument, "Occurrence name does not contain the occurrence id")
	}

	updatedOccurrence, err := r.grafeasCommon.UpdateOccurrence(ctx, &grafeas_proto.UpdateOccurrenceRequest{
		Name:       name,
		Occurrence: occurrenceRequest.Occurrence,
		UpdateMask: occurrenceRequest.UpdateMask,
	})

	if err != nil {
		log.Error("update occurrences failed", zap.Error(err))
		return nil, status.Error(codes.Internal, "update occurrences failed")
	}

	return updatedOccurrence, nil
}

func (r *rodeServer) ValidatePolicy(ctx context.Context, policy *pb.ValidatePolicyRequest) (*pb.ValidatePolicyResponse, error) {
	log := r.logger.Named("ValidatePolicy")

	if len(policy.Policy) == 0 {
		return nil, createError(log, "empty policy passed in", nil)
	}

	// Generate the AST
	mod, err := ast.ParseModule("validate_module", string(policy.Policy))
	if err != nil {
		log.Debug("failed to parse the policy", zap.Any("policy", err))
		message := &pb.ValidatePolicyResponse{
			Policy:  policy.Policy,
			Compile: false,
			Errors:  []string{err.Error()},
		}
		s, _ := status.New(codes.InvalidArgument, "failed to parse the policy").WithDetails(message)
		return message, s.Err()
	}

	// Create a new compiler instance and compile the module
	c := ast.NewCompiler()

	mods := map[string]*ast.Module{
		"validate_module": mod,
	}

	if c.Compile(mods); c.Failed() {
		log.Debug("compilation error", zap.Any("payload", c.Errors))
		length := len(c.Errors)
		errorsList := make([]string, length)

		for i := range c.Errors {
			errorsList = append(errorsList, c.Errors[i].Error())
		}

		message := &pb.ValidatePolicyResponse{
			Policy:  policy.Policy,
			Compile: false,
			Errors:  errorsList,
		}
		s, _ := status.New(codes.InvalidArgument, "failed to compile the policy").WithDetails(message)
		return message, s.Err()

	}
	log.Debug("compilation successful")

	return &pb.ValidatePolicyResponse{
		Policy:  policy.Policy,
		Compile: true,
		Errors:  nil,
	}, nil
}

func (r *rodeServer) CreatePolicy(ctx context.Context, policyEntity *pb.PolicyEntity) (*pb.Policy, error) {
	// TODO maybe check if it already exists (if we think a unique name is required)

	log := r.logger.Named("CreatePolicy")
	// Name field is a requirement
	if len(policyEntity.Name) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "policy name not provided")
	}

	// CheckPolicy before writing to elastic
	result, err := r.ValidatePolicy(ctx, &pb.ValidatePolicyRequest{Policy: policyEntity.RegoContent})
	if (err != nil) || !result.Compile {
		message := &pb.ValidatePolicyResponse{
			Policy:  policyEntity.RegoContent,
			Compile: false,
			Errors:  result.Errors,
		}
		s, _ := status.New(codes.InvalidArgument, "failed to compile the provided policy").WithDetails(message)
		return nil, s.Err()
	}

	policy := &pb.Policy{
		Id:     uuid.New().String(),
		Policy: policyEntity,
	}
	str, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(proto.MessageV2(policy))
	if err != nil {
		return nil, createError(log, fmt.Sprintf("error marshalling %T to json", policy), err)
	}
	res, err := r.esClient.Index(
		rodeElasticsearchPoliciesIndex,
		bytes.NewReader(str),
		r.esClient.Index.WithContext(ctx),
		r.esClient.Index.WithRefresh(string(r.elasticsearchConfig.Refresh.String())),
	)

	if err != nil {
		return nil, createError(log, "error sending request to elasticsearch", err)
	}

	if res.IsError() {
		return nil, createError(log, "error indexing document in elasticsearch", nil, zap.String("response", res.String()), zap.Int("status", res.StatusCode))
	}
	return policy, nil
}

func (r *rodeServer) GetPolicy(ctx context.Context, getPolicyRequest *pb.GetPolicyRequest) (*pb.Policy, error) {
	log := r.logger.Named("GetPolicy")

	search := &esSearch{
		Query: &filtering.Query{
			Term: &filtering.Term{
				"id.keyword": getPolicyRequest.Id,
			},
		},
	}

	policy := &pb.Policy{}

	_, err := r.genericGet(ctx, log, search, rodeElasticsearchPoliciesIndex, policy)
	if err != nil {
		return nil, err
	}

	return policy, nil

}

func (r *rodeServer) DeletePolicy(ctx context.Context, deletePolicyRequest *pb.DeletePolicyRequest) (*emptypb.Empty, error) {
	log := r.logger.Named("DeletePolicy")

	search := &esSearch{
		Query: &filtering.Query{
			Term: &filtering.Term{
				"id.keyword": deletePolicyRequest.Id,
			},
		},
	}

	encodedBody, requestJSON := encodeRequest(search)
	log.Debug("es request payload", zap.Any("payload", requestJSON))

	res, err := r.esClient.DeleteByQuery(
		[]string{rodeElasticsearchPoliciesIndex},
		encodedBody,
		r.esClient.DeleteByQuery.WithContext(ctx),
		r.esClient.DeleteByQuery.WithRefresh(withRefreshBool(r.elasticsearchConfig.Refresh)),
	)
	if err != nil {
		return nil, createError(log, "error sending request to elasticsearch", err)
	}
	if res.IsError() {
		return nil, createError(log, "received unexpected response from elasticsearch", nil)
	}

	var deletedResults esDeleteResponse
	if err = decodeResponse(res.Body, &deletedResults); err != nil {
		return nil, createError(log, "error unmarshalling elasticsearch response", err)
	}

	if deletedResults.Deleted == 0 {
		return nil, createError(log, "elasticsearch returned zero deleted documents", nil, zap.Any("response", deletedResults))
	}

	return &emptypb.Empty{}, nil
}

func (r *rodeServer) ListPolicies(ctx context.Context, listPoliciesRequest *pb.ListPoliciesRequest) (*pb.ListPoliciesResponse, error) {
	log := r.logger.Named("List Policies")

	body := &esSearch{}
	if listPoliciesRequest.Filter != "" {
		log = log.With(zap.String("filter", listPoliciesRequest.Filter))
		filterQuery, err := r.filterer.ParseExpression(listPoliciesRequest.Filter)
		if err != nil {
			return nil, createError(log, "error while parsing filter expression", err)
		}

		body.Query = filterQuery
	}
	encodedBody, requestJson := esutil.EncodeRequest(body)
	log = log.With(zap.String("request", requestJson))
	log.Debug("performing search")

	var policies []*pb.Policy
	res, err := r.esClient.Search(
		r.esClient.Search.WithContext(ctx),
		r.esClient.Search.WithIndex(rodeElasticsearchPoliciesIndex),
		r.esClient.Search.WithBody(encodedBody),
		r.esClient.Search.WithSize(maxPageSize),
	)

	if err != nil {
		return nil, createError(log, "error sending request to elasticsearch", err)
	}
	if res.IsError() {
		return nil, createError(log, "error searching elasticsearch for document", nil, zap.String("response", res.String()), zap.Int("status", res.StatusCode))
	}

	var searchResults esSearchResponse
	if err := decodeResponse(res.Body, &searchResults); err != nil {
		return nil, createError(log, "error unmarshalling elasticsearch response", err)
	}

	if searchResults.Hits.Total.Value == 0 {
		log.Debug("document not found", zap.Any("search", "filter replace here"))
		return &pb.ListPoliciesResponse{}, nil
	}

	for _, hit := range searchResults.Hits.Hits {
		hitLogger := log.With(zap.String("policy raw", string(hit.Source)))

		policy := &pb.Policy{}
		err := protojson.Unmarshal(hit.Source, proto.MessageV2(policy))
		if err != nil {
			log.Error("failed to convert _doc to policy", zap.Error(err))
			return nil, createError(hitLogger, "error converting _doc to policy", err)
		}

		hitLogger.Debug("policy hit", zap.Any("policy", policy))

		policies = append(policies, policy)
	}

	return &pb.ListPoliciesResponse{Policies: policies}, nil
}

// UpdatePolicy will update only the fields provided by the user
func (r *rodeServer) UpdatePolicy(ctx context.Context, updatePolicyRequest *pb.UpdatePolicyRequest) (*pb.Policy, error) {
	log := r.logger.Named("Update Policy")

	// check if the policy exists
	search := &esSearch{
		Query: &filtering.Query{
			Term: &filtering.Term{
				"id.keyword": updatePolicyRequest.Id,
			},
		},
	}

	policy := &pb.Policy{}
	targetDocumentID, err := r.genericGet(ctx, log, search, rodeElasticsearchPoliciesIndex, policy)
	if err != nil {
		return nil, err
	}

	log.Debug("field masks", zap.Any("response", updatePolicyRequest.UpdateMask.Paths))
	// if one of the fields being updated is the rego policy, revalidate the policy
	if contains(updatePolicyRequest.UpdateMask.Paths, "rego_content") {
		_, err = r.ValidatePolicy(ctx, &pb.ValidatePolicyRequest{Policy: updatePolicyRequest.Policy.RegoContent})
	}
	if err != nil {
		return nil, err
	}

	m, err := fieldmask_utils.MaskFromPaths(updatePolicyRequest.UpdateMask.Paths, generator.CamelCase)
	if err != nil {
		log.Info("errors while mapping masks", zap.Any("errors", err))
		return policy, err
	}
	fieldmask_utils.StructToStruct(m, updatePolicyRequest.Policy, policy.Policy)

	str, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(proto.MessageV2(policy))
	if err != nil {
		return nil, createError(log, fmt.Sprintf("error marshalling %T to json", policy), err)
	}

	res, err := r.esClient.Index(
		rodeElasticsearchPoliciesIndex,
		bytes.NewReader(str),
		r.esClient.Index.WithDocumentID(targetDocumentID),
		r.esClient.Index.WithContext(ctx),
		r.esClient.Index.WithRefresh(r.elasticsearchConfig.Refresh.String()),
	)
	if err != nil {
		return nil, createError(log, "error sending request to elasticsearch", err)
	}

	if res.IsError() {
		return nil, createError(log, "error indexing document in elasticsearch", nil, zap.String("response", res.String()), zap.Int("status", res.StatusCode))
	}

	return policy, nil
}

func (r *rodeServer) genericGet(ctx context.Context, log *zap.Logger, search *esSearch, index string, protoMessage interface{}) (string, error) {
	encodedBody, requestJson := esutil.EncodeRequest(search)
	log = log.With(zap.String("request", requestJson))

	res, err := r.esClient.Search(
		r.esClient.Search.WithContext(ctx),
		r.esClient.Search.WithIndex(index),
		r.esClient.Search.WithBody(encodedBody),
	)
	if err != nil {
		return "", createError(log, "error sending request to elasticsearch", err)
	}
	if res.IsError() {
		return "", createError(log, "error searching elasticsearch for document", nil, zap.String("response", res.String()), zap.Int("status", res.StatusCode))
	}

	var searchResults esSearchResponse
	if err := esutil.DecodeResponse(res.Body, &searchResults); err != nil {
		return "", createError(log, "error unmarshalling elasticsearch response", err)
	}

	if searchResults.Hits.Total.Value == 0 {
		log.Debug("document not found", zap.Any("search", search))
		return "", status.Error(codes.NotFound, fmt.Sprintf("%T not found", protoMessage))
	}

	return searchResults.Hits.Hits[0].ID, protojson.Unmarshal(searchResults.Hits.Hits[0].Source, proto.MessageV2(protoMessage))
}

// createError is a helper function that allows you to easily log an error and return a gRPC formatted error.
func createError(log *zap.Logger, message string, err error, fields ...zap.Field) error {
	if err == nil {
		log.Error(message, fields...)
		return status.Errorf(codes.Internal, "%s", message)
	}

	log.Error(message, append(fields, zap.Error(err))...)
	return status.Errorf(codes.Internal, "%s: %s", message, err)
}

type esDeleteResponse struct {
	Deleted int `json:"deleted"`
}

type esIndexDocResponse struct {
	Id      string           `json:"_id"`
	Result  string           `json:"result"`
	Version int              `json:"_version"`
	Status  int              `json:"status"`
	Error   *esIndexDocError `json:"error,omitempty"`
}

type esIndexDocError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

// contains returns a boolean describing whether or not a string slice contains a particular string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func withRefreshBool(o config.RefreshOption) bool {
	if o == config.RefreshFalse {
		return false
	}
	return true
}
