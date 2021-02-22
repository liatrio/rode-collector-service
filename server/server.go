package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	"github.com/rode/rode/opa"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	rodeElasticsearchOccurrencesIndex = "grafeas-v1beta1-rode-occurrences"
)

// NewRodeServer constructor for rodeServer
func NewRodeServer(
	logger *zap.Logger,
	grafeasCommon grafeas_proto.GrafeasV1Beta1Client,
	grafeasProjects grafeas_project_proto.ProjectsClient,
	opa opa.Client,
	esClient *elasticsearch.Client,
	filterer filtering.Filterer,
) (pb.RodeServer, error) {
	rodeServer := &rodeServer{
		logger:          logger,
		grafeasCommon:   grafeasCommon,
		grafeasProjects: grafeasProjects,
		opa:             opa,
		esClient:        esClient,
		filterer:        filterer,
	}
	if err := rodeServer.initialize(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to initialize rode server: %s", err)
	}
	return rodeServer, nil
}

type rodeServer struct {
	pb.UnimplementedRodeServer
	logger          *zap.Logger
	esClient        *elasticsearch.Client
	filterer        filtering.Filterer
	grafeasCommon   grafeas_proto.GrafeasV1Beta1Client
	grafeasProjects grafeas_project_proto.ProjectsClient
	opa             opa.Client
}

func (r *rodeServer) BatchCreateOccurrences(ctx context.Context, occurrenceRequest *pb.BatchCreateOccurrencesRequest) (*pb.BatchCreateOccurrencesResponse, error) {
	log := r.logger.Named("BatchCreateOccurrences")
	log.Debug("received request", zap.Any("BatchCreateOccurrencesRequest", occurrenceRequest))

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
		r.esClient.Search.WithIndex(rodeElasticsearchOccurrencesIndex),
		r.esClient.Search.WithBody(encodedBody),
		r.esClient.Search.WithSize(1000),
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

	return nil
}
