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
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering/filteringfakes"
	"github.com/rode/rode/pkg/policy/policyfakes"
	"github.com/rode/rode/pkg/resource/resourcefakes"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	immocks "github.com/rode/es-index-manager/mocks"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/esutil"
	"github.com/rode/grafeas-elasticsearch/go/v1beta1/storage/filtering"
	"github.com/rode/rode/config"
	"github.com/rode/rode/mocks"
	pb "github.com/rode/rode/proto/v1alpha1"
	"github.com/rode/rode/protodeps/grafeas/proto/v1beta1/build_go_proto"
	grafeas_common_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/common_go_proto"
	"github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"
	"github.com/rode/rode/protodeps/grafeas/proto/v1beta1/provenance_go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ = Describe("rode server", func() {
	var (
		server                pb.RodeServer
		grafeasClient         *mocks.FakeGrafeasV1Beta1Client
		grafeasProjectsClient *mocks.FakeProjectsClient
		esClient              *elasticsearch.Client
		esTransport           *mockEsTransport
		filterer              *filteringfakes.FakeFilterer
		elasticsearchConfig   *config.ElasticsearchConfig
		resourceManager       *resourcefakes.FakeManager
		policyManager         *policyfakes.FakeManager
		indexManager          *immocks.FakeIndexManager
		ctx                   context.Context

		expectedPoliciesIndex        string
		expectedPoliciesAlias        string
		expectedGenericResourceIndex string
		expectedGenericResourceAlias string
	)

	BeforeEach(func() {
		grafeasClient = &mocks.FakeGrafeasV1Beta1Client{}
		grafeasProjectsClient = &mocks.FakeProjectsClient{}
		resourceManager = &resourcefakes.FakeManager{}
		elasticsearchConfig = &config.ElasticsearchConfig{
			Refresh: "true",
		}
		filterer = &filteringfakes.FakeFilterer{}

		esTransport = &mockEsTransport{}
		esClient = &elasticsearch.Client{
			Transport: esTransport,
			API:       esapi.New(esTransport),
		}

		expectedPoliciesIndex = gofakeit.LetterN(10)
		expectedPoliciesAlias = gofakeit.LetterN(10)
		expectedGenericResourceIndex = gofakeit.LetterN(10)
		expectedGenericResourceAlias = gofakeit.LetterN(10)
		indexManager = &immocks.FakeIndexManager{}

		indexManager.AliasNameStub = func(documentKind, _ string) string {
			return map[string]string{
				genericResourcesDocumentKind: expectedGenericResourceAlias,
				policiesDocumentKind:         expectedPoliciesAlias,
			}[documentKind]
		}

		indexManager.IndexNameStub = func(documentKind, _ string) string {
			return map[string]string{
				genericResourcesDocumentKind: expectedGenericResourceIndex,
				policiesDocumentKind:         expectedPoliciesIndex,
			}[documentKind]
		}

		ctx = context.Background()

		// not using the constructor as it has side effects. side effects are tested under the "initialize" context
		server = &rodeServer{
			logger:              logger,
			grafeasCommon:       grafeasClient,
			grafeasProjects:     grafeasProjectsClient,
			esClient:            esClient,
			filterer:            filterer,
			elasticsearchConfig: elasticsearchConfig,
			resourceManager:     resourceManager,
			indexManager:        indexManager,
		}
	})

	Context("initialize", func() {
		var (
			actualRodeServer pb.RodeServer
			actualError      error

			expectedProject         *grafeas_project_proto.Project
			expectedGetProjectError error

			expectedCreateProjectError error
		)

		BeforeEach(func() {
			expectedProject = &grafeas_project_proto.Project{
				Name: fmt.Sprintf("projects/%s", gofakeit.LetterN(10)),
			}
			expectedGetProjectError = nil
			expectedCreateProjectError = nil

		})

		JustBeforeEach(func() {
			grafeasProjectsClient.GetProjectReturns(expectedProject, expectedGetProjectError)
			grafeasProjectsClient.CreateProjectReturns(expectedProject, expectedCreateProjectError)

			actualRodeServer, actualError = NewRodeServer(logger, grafeasClient, grafeasProjectsClient, esClient, filterer, elasticsearchConfig, resourceManager, indexManager, policyManager)
		})

		It("should check if the rode project exists", func() {
			Expect(grafeasProjectsClient.GetProjectCallCount()).To(Equal(1))

			_, getProjectRequest, _ := grafeasProjectsClient.GetProjectArgsForCall(0)
			Expect(getProjectRequest.Name).To(Equal(rodeProjectSlug))
		})

		// happy path: project already exists
		It("should not create a project", func() {
			Expect(grafeasProjectsClient.CreateProjectCallCount()).To(Equal(0))
		})

		It("should initialize the index manager", func() {
			Expect(indexManager.InitializeCallCount()).To(Equal(1))
		})

		It("should create an index for policies", func() {
			Expect(indexManager.CreateIndexCallCount()).To(Equal(2))

			_, actualIndexName, actualAliasName, documentKind := indexManager.CreateIndexArgsForCall(0)

			Expect(actualIndexName).To(Equal(expectedPoliciesIndex))
			Expect(actualAliasName).To(Equal(expectedPoliciesAlias))
			Expect(documentKind).To(Equal(policiesDocumentKind))
		})

		It("should create an index for generic resources", func() {
			Expect(indexManager.CreateIndexCallCount()).To(Equal(2))

			_, actualIndexName, actualAliasName, documentKind := indexManager.CreateIndexArgsForCall(1)

			Expect(actualIndexName).To(Equal(expectedGenericResourceIndex))
			Expect(actualAliasName).To(Equal(expectedGenericResourceAlias))
			Expect(documentKind).To(Equal(genericResourcesDocumentKind))
		})

		It("should return the initialized rode server", func() {
			Expect(actualRodeServer).ToNot(BeNil())
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("getting the rode project fails", func() {
			BeforeEach(func() {
				expectedGetProjectError = status.Error(codes.Internal, "getting project failed")
			})

			It("should return an error", func() {
				Expect(actualRodeServer).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})

			It("should not create a project", func() {
				Expect(grafeasProjectsClient.CreateProjectCallCount()).To(Equal(0))
			})

			It("should not attempt to create indices", func() {
				Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
			})
		})

		When("the rode project does not exist", func() {
			BeforeEach(func() {
				expectedGetProjectError = status.Error(codes.NotFound, "not found")
			})

			It("should create the rode project", func() {
				Expect(grafeasProjectsClient.CreateProjectCallCount()).To(Equal(1))

				_, createProjectRequest, _ := grafeasProjectsClient.CreateProjectArgsForCall(0)
				Expect(createProjectRequest.Project.Name).To(Equal(rodeProjectSlug))
			})

			When("creating the rode project fails", func() {
				BeforeEach(func() {
					expectedCreateProjectError = errors.New("create project failed")
				})

				It("should return an error", func() {
					Expect(actualRodeServer).To(BeNil())
					Expect(actualError).To(HaveOccurred())
				})

				It("should not attempt to create indices", func() {
					Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
				})
			})
		})

		When("initializing the index manager fails", func() {
			BeforeEach(func() {
				indexManager.InitializeReturns(errors.New(gofakeit.Word()))
			})

			It("should return an error", func() {
				Expect(actualRodeServer).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})

			It("should not create the application indices", func() {
				Expect(indexManager.CreateIndexCallCount()).To(Equal(0))
			})
		})

		When("creating the first index fails", func() {
			BeforeEach(func() {
				indexManager.CreateIndexReturns(errors.New(gofakeit.Word()))
			})

			It("should return an error", func() {
				Expect(actualRodeServer).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})

			It("should not attempt to create another index", func() {
				Expect(indexManager.CreateIndexCallCount()).To(Equal(1))
			})
		})

		When("creating the second index fails", func() {
			BeforeEach(func() {
				indexManager.CreateIndexReturnsOnCall(1, errors.New(gofakeit.Word()))
			})

			It("should return an error", func() {
				Expect(actualRodeServer).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})
	})

	Context("BatchCreateOccurrences", func() {
		var (
			actualRodeBatchCreateOccurrencesResponse *pb.BatchCreateOccurrencesResponse
			actualError                              error

			expectedRodeBatchCreateOccurrencesRequest *pb.BatchCreateOccurrencesRequest

			expectedOccurrence *grafeas_proto.Occurrence

			expectedGrafeasBatchCreateOccurrencesResponse *grafeas_proto.BatchCreateOccurrencesResponse
			expectedGrafeasBatchCreateOccurrencesError    error

			expectedBatchCreateResourcesError        error
			expectedBatchCreateResourceVersionsError error

			expectedResourceName string
		)

		BeforeEach(func() {
			expectedOccurrence = createRandomOccurrence(grafeas_common_proto.NoteKind_NOTE_KIND_UNSPECIFIED)
			expectedResourceName = gofakeit.URL()
			expectedOccurrence.Resource.Uri = fmt.Sprintf("%s@sha256:%s", expectedResourceName, gofakeit.LetterN(10))

			expectedGrafeasBatchCreateOccurrencesResponse = &grafeas_proto.BatchCreateOccurrencesResponse{
				Occurrences: []*grafeas_proto.Occurrence{
					expectedOccurrence,
				},
			}
			expectedGrafeasBatchCreateOccurrencesError = nil

			expectedRodeBatchCreateOccurrencesRequest = &pb.BatchCreateOccurrencesRequest{
				Occurrences: []*grafeas_proto.Occurrence{
					expectedOccurrence,
				},
			}

			expectedBatchCreateResourcesError = nil
			expectedBatchCreateResourceVersionsError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.BatchCreateOccurrencesReturns(expectedGrafeasBatchCreateOccurrencesResponse, expectedGrafeasBatchCreateOccurrencesError)
			resourceManager.BatchCreateGenericResourcesReturns(expectedBatchCreateResourcesError)
			resourceManager.BatchCreateGenericResourceVersionsReturns(expectedBatchCreateResourceVersionsError)

			actualRodeBatchCreateOccurrencesResponse, actualError = server.BatchCreateOccurrences(ctx, expectedRodeBatchCreateOccurrencesRequest)
		})

		It("should send occurrences to Grafeas", func() {
			Expect(grafeasClient.BatchCreateOccurrencesCallCount()).To(Equal(1))

			_, batchCreateOccurrencesRequest, _ := grafeasClient.BatchCreateOccurrencesArgsForCall(0)
			Expect(batchCreateOccurrencesRequest.Occurrences).To(HaveLen(1))
			Expect(batchCreateOccurrencesRequest.Occurrences[0]).To(BeEquivalentTo(expectedOccurrence))
		})

		It("should create generic resources from the received occurrences", func() {
			Expect(resourceManager.BatchCreateGenericResourcesCallCount()).To(Equal(1))

			_, occurrences := resourceManager.BatchCreateGenericResourcesArgsForCall(0)
			Expect(occurrences).To(BeEquivalentTo(expectedRodeBatchCreateOccurrencesRequest.Occurrences))
		})

		It("should create generic resource versions from the received occurrences", func() {
			Expect(resourceManager.BatchCreateGenericResourceVersionsCallCount()).To(Equal(1))

			_, occurrences := resourceManager.BatchCreateGenericResourceVersionsArgsForCall(0)
			Expect(occurrences).To(BeEquivalentTo(expectedRodeBatchCreateOccurrencesRequest.Occurrences))
		})

		It("should return the created occurrences", func() {
			Expect(actualRodeBatchCreateOccurrencesResponse.Occurrences).To(HaveLen(1))
			Expect(actualRodeBatchCreateOccurrencesResponse.Occurrences[0]).To(BeEquivalentTo(expectedOccurrence))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("an error occurs while creating occurrences", func() {
			BeforeEach(func() {
				expectedGrafeasBatchCreateOccurrencesError = errors.New("error batch creating occurrences")
			})

			It("should return an error", func() {
				Expect(actualRodeBatchCreateOccurrencesResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})

			It("should not attempt to create generic resources", func() {
				Expect(resourceManager.BatchCreateGenericResourcesCallCount()).To(Equal(0))
			})

			It("should not attempt to create generic resource versions", func() {
				Expect(resourceManager.BatchCreateGenericResourceVersionsCallCount()).To(Equal(0))
			})
		})

		When("an error occurs while creating generic resources", func() {
			BeforeEach(func() {
				expectedBatchCreateResourcesError = errors.New("error batch creating generic resources")
			})

			It("should not attempt to create generic resource versions", func() {
				Expect(resourceManager.BatchCreateGenericResourceVersionsCallCount()).To(Equal(0))
			})

			It("should return an error", func() {
				Expect(actualRodeBatchCreateOccurrencesResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})

		When("an error occurs while creating generic resource versions", func() {
			BeforeEach(func() {
				expectedBatchCreateResourceVersionsError = errors.New("error creating generic resource versions")
			})

			It("should return an error", func() {
				Expect(actualRodeBatchCreateOccurrencesResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})
	})

	Context("ListGenericResources", func() {
		var (
			expectedListGenericResourcesRequest *pb.ListGenericResourcesRequest

			expectedListGenericResourcesResponse *pb.ListGenericResourcesResponse
			expectedListGenericResourcesError    error

			actualListGenericResourcesResponse *pb.ListGenericResourcesResponse
			actualError                        error
		)

		BeforeEach(func() {
			expectedListGenericResourcesRequest = &pb.ListGenericResourcesRequest{
				Filter: gofakeit.LetterN(10),
			}

			expectedListGenericResourcesResponse = &pb.ListGenericResourcesResponse{
				GenericResources: []*pb.GenericResource{
					{
						Name: gofakeit.LetterN(10),
						Type: pb.ResourceType(gofakeit.Number(0, 6)),
					},
				},
			}
			expectedListGenericResourcesError = nil
		})

		JustBeforeEach(func() {
			resourceManager.ListGenericResourcesReturns(expectedListGenericResourcesResponse, expectedListGenericResourcesError)

			actualListGenericResourcesResponse, actualError = server.ListGenericResources(ctx, expectedListGenericResourcesRequest)
		})

		It("should return the result from the resource manager", func() {
			Expect(resourceManager.ListGenericResourcesCallCount()).To(Equal(1))

			_, listGenericResourcesRequest := resourceManager.ListGenericResourcesArgsForCall(0)
			Expect(listGenericResourcesRequest).To(Equal(expectedListGenericResourcesRequest))

			Expect(actualListGenericResourcesResponse).To(Equal(expectedListGenericResourcesResponse))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("the resource manager returns an error", func() {
			BeforeEach(func() {
				expectedListGenericResourcesResponse = nil
				expectedListGenericResourcesError = errors.New("error listing generic resources")
			})

			It("should return an error", func() {
				Expect(actualListGenericResourcesResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})
	})

	Context("ListGenericResourceVersions", func() {
		var (
			actualListGenericResourceVersionsResponse *pb.ListGenericResourceVersionsResponse
			actualError                               error

			expectedResourceId              string
			expectedGenericResource         *pb.GenericResource
			expectedGetGenericResourceError error

			expectedListGenericResourceVersionsRequest  *pb.ListGenericResourceVersionsRequest
			expectedListGenericResourceVersionsResponse *pb.ListGenericResourceVersionsResponse
			expectedListGenericResourceVersionsError    error
		)

		BeforeEach(func() {
			expectedResourceId = gofakeit.LetterN(10)
			expectedGenericResource = &pb.GenericResource{
				Id:   expectedResourceId,
				Name: gofakeit.LetterN(10),
				Type: pb.ResourceType(gofakeit.Number(0, 7)),
			}

			expectedListGenericResourceVersionsRequest = &pb.ListGenericResourceVersionsRequest{
				Id: expectedResourceId,
			}
			expectedListGenericResourceVersionsResponse = &pb.ListGenericResourceVersionsResponse{
				Versions: []*pb.GenericResourceVersion{
					{
						Version: gofakeit.LetterN(10),
						Names:   []string{gofakeit.LetterN(10)},
						Created: timestamppb.Now(),
					},
				},
			}

			expectedGetGenericResourceError = nil
			expectedListGenericResourceVersionsError = nil
		})

		JustBeforeEach(func() {
			resourceManager.GetGenericResourceReturns(expectedGenericResource, expectedGetGenericResourceError)
			resourceManager.ListGenericResourceVersionsReturns(expectedListGenericResourceVersionsResponse, expectedListGenericResourceVersionsError)

			actualListGenericResourceVersionsResponse, actualError = server.ListGenericResourceVersions(ctx, expectedListGenericResourceVersionsRequest)
		})

		It("should fetch the generic resource with the provided id", func() {
			Expect(resourceManager.GetGenericResourceCallCount()).To(Equal(1))

			_, id := resourceManager.GetGenericResourceArgsForCall(0)
			Expect(id).To(Equal(expectedResourceId))
		})

		It("should list generic resource versions", func() {
			Expect(resourceManager.ListGenericResourceVersionsCallCount()).To(Equal(1))

			_, request := resourceManager.ListGenericResourceVersionsArgsForCall(0)
			Expect(request).To(Equal(expectedListGenericResourceVersionsRequest))
		})

		It("should return the result and no error", func() {
			Expect(actualListGenericResourceVersionsResponse).To(Equal(expectedListGenericResourceVersionsResponse))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("the generic resource id is not specified", func() {
			BeforeEach(func() {
				expectedListGenericResourceVersionsRequest.Id = ""
			})

			It("should not list generic resource versions", func() {
				Expect(resourceManager.ListGenericResourceVersionsCallCount()).To(Equal(0))
			})

			It("should return an invalid argument error", func() {
				Expect(actualListGenericResourceVersionsResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.InvalidArgument))
			})
		})

		When("the generic resource is not found", func() {
			BeforeEach(func() {
				expectedGenericResource = nil
			})

			It("should return a not found error", func() {
				Expect(actualListGenericResourceVersionsResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.NotFound))
			})

			It("should not list generic resource versions", func() {
				Expect(resourceManager.ListGenericResourceVersionsCallCount()).To(Equal(0))
			})
		})

		When("getting the generic resource fails", func() {
			BeforeEach(func() {
				expectedGetGenericResourceError = errors.New("getting generic resource failed")
			})

			It("should not list generic resource versions", func() {
				Expect(resourceManager.ListGenericResourceVersionsCallCount()).To(Equal(0))
			})

			It("should return an error", func() {
				Expect(actualListGenericResourceVersionsResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})

		When("listing generic resource versions fails", func() {
			BeforeEach(func() {
				expectedListGenericResourceVersionsError = errors.New("list generic resource versions failed")
				expectedListGenericResourceVersionsResponse = nil
			})

			It("should return an error", func() {
				Expect(actualListGenericResourceVersionsResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})
	})

	Context("ListVersionedResourceOccurrences", func() {
		var (
			listBuildOccurrencesResponse *grafeas_proto.ListOccurrencesResponse
			listBuildOccurrencesError    error

			listAllOccurrencesResponse *grafeas_proto.ListOccurrencesResponse
			listAllOccurrencesError    error

			gitResourceUri string

			nextPageToken    string
			currentPageToken string
			pageSize         int32
			ctx              context.Context
			resourceUri      string
			request          *pb.ListVersionedResourceOccurrencesRequest
			actualResponse   *pb.ListVersionedResourceOccurrencesResponse
			actualError      error
		)

		BeforeEach(func() {
			ctx = context.Background()
			resourceUri = gofakeit.URL()
			nextPageToken = gofakeit.Word()
			currentPageToken = gofakeit.Word()
			pageSize = gofakeit.Int32()

			request = &pb.ListVersionedResourceOccurrencesRequest{
				ResourceUri: resourceUri,
				PageToken:   currentPageToken,
				PageSize:    pageSize,
			}

			gitResourceUri = fmt.Sprintf("git://%s", gofakeit.DomainName())

			listBuildOccurrencesResponse = &grafeas_proto.ListOccurrencesResponse{
				Occurrences: []*grafeas_proto.Occurrence{
					{
						Resource: &grafeas_proto.Resource{
							Uri: gitResourceUri,
						},
						Kind: grafeas_common_proto.NoteKind_BUILD,
						Details: &grafeas_proto.Occurrence_Build{
							Build: &build_go_proto.Details{
								Provenance: &provenance_go_proto.BuildProvenance{
									BuiltArtifacts: []*provenance_go_proto.Artifact{
										{
											Id: resourceUri,
										},
									},
								},
							},
						},
					},
				},
			}
			listBuildOccurrencesError = nil

			listAllOccurrencesResponse = &grafeas_proto.ListOccurrencesResponse{
				Occurrences: []*grafeas_proto.Occurrence{
					createRandomOccurrence(grafeas_common_proto.NoteKind_VULNERABILITY),
					createRandomOccurrence(grafeas_common_proto.NoteKind_BUILD),
				},
				NextPageToken: nextPageToken,
			}
			listAllOccurrencesError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.ListOccurrencesReturnsOnCall(0, listBuildOccurrencesResponse, listBuildOccurrencesError)
			grafeasClient.ListOccurrencesReturnsOnCall(1, listAllOccurrencesResponse, listAllOccurrencesError)

			actualResponse, actualError = server.ListVersionedResourceOccurrences(ctx, request)
		})

		It("should list build occurrences for the resource uri", func() {
			_, buildOccurrencesRequest, _ := grafeasClient.ListOccurrencesArgsForCall(0)

			Expect(buildOccurrencesRequest).NotTo(BeNil())
			Expect(buildOccurrencesRequest.Parent).To(Equal("projects/rode"))
			Expect(buildOccurrencesRequest.Filter).To(ContainSubstring(fmt.Sprintf(`build.provenance.builtArtifacts.nestedFilter(id == "%s")`, resourceUri)))
			Expect(buildOccurrencesRequest.Filter).To(ContainSubstring(fmt.Sprintf(`resource.uri == "%s"`, resourceUri)))
			Expect(buildOccurrencesRequest.PageSize).To(Equal(int32(1000)))
		})

		It("should use the build occurrence to find all occurrences", func() {
			expectedFilter := []string{
				fmt.Sprintf(`resource.uri == "%s"`, resourceUri),
				fmt.Sprintf(`resource.uri == "%s"`, gitResourceUri),
			}

			_, allOccurrencesRequest, _ := grafeasClient.ListOccurrencesArgsForCall(1)

			Expect(allOccurrencesRequest).NotTo(BeNil())
			Expect(allOccurrencesRequest.Parent).To(Equal("projects/rode"))
			Expect(allOccurrencesRequest.PageSize).To(Equal(pageSize))
			Expect(allOccurrencesRequest.PageToken).To(Equal(currentPageToken))

			filterParts := strings.Split(allOccurrencesRequest.Filter, " || ")
			Expect(filterParts).To(ConsistOf(expectedFilter))
		})

		It("should return the occurrences and page token from the call to list all occurrences", func() {
			Expect(actualResponse.Occurrences).To(BeEquivalentTo(listAllOccurrencesResponse.Occurrences))
			Expect(actualResponse.NextPageToken).To(BeEquivalentTo(listAllOccurrencesResponse.NextPageToken))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("there are no build occurrences", func() {
			BeforeEach(func() {
				listBuildOccurrencesResponse.Occurrences = []*grafeas_proto.Occurrence{}
			})

			It("should list occurrences for the resource uri", func() {
				_, allOccurrencesRequest, _ := grafeasClient.ListOccurrencesArgsForCall(1)

				Expect(allOccurrencesRequest.Filter).To(Equal(fmt.Sprintf(`resource.uri == "%s"`, resourceUri)))
			})
		})

		When("the resource uri is not specified", func() {
			BeforeEach(func() {
				request.ResourceUri = ""
			})

			It("should return an error", func() {
				Expect(actualResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.InvalidArgument))
			})
		})

		When("an error occurs listing build occurrences", func() {
			BeforeEach(func() {
				listBuildOccurrencesError = errors.New("error listing build occurrences")
			})

			It("should return an error", func() {
				Expect(actualResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
			})
		})

		When("an error occurs listing all occurrences", func() {
			BeforeEach(func() {
				listAllOccurrencesError = errors.New("error listing all occurrences")
			})

			It("should return an error", func() {
				Expect(actualResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
			})
		})
	})

	Context("ListOccurrences", func() {
		var (
			actualResponse *pb.ListOccurrencesResponse
			actualError    error

			expectedOccurrence *grafeas_proto.Occurrence
			expectedPageToken  string
			expectedPageSize   int32
			expectedFilter     string

			expectedListOccurrencesRequest *pb.ListOccurrencesRequest

			expectedGrafeasListOccurrencesResponse *grafeas_proto.ListOccurrencesResponse
			expectedGrafeasListOccurrencesError    error
		)

		BeforeEach(func() {
			expectedOccurrence = createRandomOccurrence(grafeas_common_proto.NoteKind_NOTE_KIND_UNSPECIFIED)

			expectedPageToken = gofakeit.Word()
			expectedPageSize = gofakeit.Int32()
			expectedFilter = fmt.Sprintf(`"resource.uri" == "%s"`, expectedOccurrence.Resource.Uri)

			expectedListOccurrencesRequest = &pb.ListOccurrencesRequest{
				Filter:    expectedFilter,
				PageToken: expectedPageToken,
				PageSize:  expectedPageSize,
			}

			expectedGrafeasListOccurrencesResponse = &grafeas_proto.ListOccurrencesResponse{
				Occurrences: []*grafeas_proto.Occurrence{
					expectedOccurrence,
				},
				NextPageToken: gofakeit.Word(),
			}

			expectedGrafeasListOccurrencesError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.ListOccurrencesReturns(expectedGrafeasListOccurrencesResponse, expectedGrafeasListOccurrencesError)

			actualResponse, actualError = server.ListOccurrences(context.Background(), expectedListOccurrencesRequest)
		})

		It("should list occurrences from grafeas", func() {
			Expect(grafeasClient.ListOccurrencesCallCount()).To(Equal(1))
			_, listOccurrencesRequest, _ := grafeasClient.ListOccurrencesArgsForCall(0)

			Expect(listOccurrencesRequest.Parent).To(Equal(rodeProjectSlug))
			Expect(listOccurrencesRequest.Filter).To(Equal(expectedFilter))
			Expect(listOccurrencesRequest.PageToken).To(Equal(expectedPageToken))
			Expect(listOccurrencesRequest.PageSize).To(Equal(expectedPageSize))
		})

		It("should return the results from grafeas", func() {
			Expect(actualResponse.Occurrences).To(BeEquivalentTo(expectedGrafeasListOccurrencesResponse.Occurrences))
			Expect(actualResponse.NextPageToken).To(Equal(expectedGrafeasListOccurrencesResponse.NextPageToken))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("Grafeas returns an error", func() {
			BeforeEach(func() {
				expectedGrafeasListOccurrencesError = errors.New("error listing occurrences")
			})

			It("should return an error", func() {
				Expect(actualResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
			})
		})
	})

	Context("UpdateOccurrence", func() {
		var (
			actualError    error
			actualResponse *grafeas_go_proto.Occurrence

			expectedOccurrence              *grafeas_proto.Occurrence
			expectedUpdateOccurrenceRequest *pb.UpdateOccurrenceRequest

			expectedGrafeasUpdateOccurrenceResponse *grafeas_proto.Occurrence
			expectedGrafeasUpdateOccurrenceError    error
		)

		BeforeEach(func() {
			expectedOccurrence = createRandomOccurrence(grafeas_common_proto.NoteKind_NOTE_KIND_UNSPECIFIED)
			occurrenceId := gofakeit.UUID()
			occurrenceName := fmt.Sprintf("projects/rode/occurrences/%s", occurrenceId)
			expectedOccurrence.Name = occurrenceName
			expectedUpdateOccurrenceRequest = &pb.UpdateOccurrenceRequest{
				Id:         occurrenceId,
				Occurrence: expectedOccurrence,
				UpdateMask: &fieldmaskpb.FieldMask{
					Paths: []string{gofakeit.Word()},
				},
			}

			expectedGrafeasUpdateOccurrenceResponse = createRandomOccurrence(grafeas_common_proto.NoteKind_NOTE_KIND_UNSPECIFIED)
			expectedGrafeasUpdateOccurrenceError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.UpdateOccurrenceReturns(expectedGrafeasUpdateOccurrenceResponse, expectedGrafeasUpdateOccurrenceError)

			actualResponse, actualError = server.UpdateOccurrence(context.Background(), expectedUpdateOccurrenceRequest)
		})

		It("should update the occurrence in grafeas", func() {
			Expect(grafeasClient.UpdateOccurrenceCallCount()).To(Equal(1))

			_, updateOccurrenceRequest, _ := grafeasClient.UpdateOccurrenceArgsForCall(0)
			Expect(updateOccurrenceRequest.Name).To(Equal(expectedOccurrence.Name))
			Expect(updateOccurrenceRequest.Occurrence).To(Equal(expectedOccurrence))
			Expect(updateOccurrenceRequest.UpdateMask).To(Equal(expectedUpdateOccurrenceRequest.UpdateMask))
		})

		It("should return the updated occurrence", func() {
			Expect(actualError).ToNot(HaveOccurred())
			Expect(actualResponse).To(Equal(expectedGrafeasUpdateOccurrenceResponse))
		})

		When("Grafeas returns an error", func() {
			BeforeEach(func() {
				expectedGrafeasUpdateOccurrenceError = errors.New("error updating occurrence")
			})

			It("should return an error", func() {
				Expect(actualError).To(HaveOccurred())
				Expect(actualResponse).To(BeNil())
			})
		})

		When("the occurrence name doesn't contain the occurrence id", func() {
			BeforeEach(func() {
				expectedUpdateOccurrenceRequest.Id = gofakeit.UUID()
			})

			It("should return an error", func() {
				Expect(actualError).ToNot(BeNil())
			})

			It("should return a status code of invalid argument", func() {
				s, ok := status.FromError(actualError)
				Expect(ok).To(BeTrue(), "Expected error to be a gRPC status")

				Expect(s.Code()).To(Equal(codes.InvalidArgument))
				Expect(s.Message()).To(ContainSubstring("occurrence name does not contain the occurrence id"))
			})

			It("should not attempt to update the occurrence", func() {
				Expect(grafeasClient.UpdateOccurrenceCallCount()).To(Equal(0))
			})
		})
	})

	Context("ListResources", func() {
		var (
			occurrences              []*grafeas_proto.Occurrence
			request                  *pb.ListResourcesRequest
			listResourcesResponse    *pb.ListResourcesResponse
			listResourcesResponseErr error
		)

		BeforeEach(func() {
			request = &pb.ListResourcesRequest{}
			occurrences = []*grafeas_proto.Occurrence{
				createRandomOccurrence(grafeas_common_proto.NoteKind_VULNERABILITY),
				createRandomOccurrence(grafeas_common_proto.NoteKind_ATTESTATION),
			}
			esTransport.preparedHttpResponses = []*http.Response{
				{
					StatusCode: http.StatusOK,
					Body:       createEsSearchResponse(occurrences),
				},
			}
		})

		When("querying for resources without a filter", func() {
			BeforeEach(func() {
				Expect(filterer.ParseExpressionCallCount()).To(Equal(0))

				listResourcesResponse, listResourcesResponseErr = server.ListResources(context.Background(), request)
			})

			It("should query the Rode occurrences index", func() {
				actualRequest := esTransport.receivedHttpRequests[0]

				Expect(actualRequest.URL.Path).To(Equal("/grafeas-rode-occurrences/_search"))
			})

			It("should take the first 1000 matches", func() {
				actualRequest := esTransport.receivedHttpRequests[0]
				query := actualRequest.URL.Query()

				Expect(query.Get("size")).To(Equal("1000"))
			})

			It("should collapse fields on resource.uri", func() {
				actualRequest := esTransport.receivedHttpRequests[0]
				search := readEsSearchResponse(actualRequest)

				Expect(search.Collapse.Field).To(Equal("resource.uri"))
			})

			It("should not return an error", func() {
				Expect(listResourcesResponseErr).To(BeNil())
			})

			It("should return all resources from the query", func() {
				var expectedResourceUris []string

				for _, occurrence := range occurrences {
					expectedResourceUris = append(expectedResourceUris, occurrence.Resource.Uri)
				}

				var actualResourceUris []string
				for _, resource := range listResourcesResponse.Resources {
					actualResourceUris = append(actualResourceUris, resource.Uri)
				}

				Expect(actualResourceUris).To(ConsistOf(expectedResourceUris))
			})
		})

		When("querying for resources with a filter", func() {
			BeforeEach(func() {
				request.Filter = gofakeit.UUID()
			})

			It("should pass the filter to the filterer", func() {
				_, _ = server.ListResources(context.Background(), request)

				Expect(filterer.ParseExpressionCallCount()).To(Equal(1))
				actualFilter := filterer.ParseExpressionArgsForCall(0)

				Expect(actualFilter).To(Equal(request.Filter))
			})

			It("should include the Elasticsearch query in the response", func() {
				expectedQuery := &filtering.Query{
					Term: &filtering.Term{
						gofakeit.UUID(): gofakeit.UUID(),
					},
				}
				filterer.ParseExpressionReturns(expectedQuery, nil)

				_, err := server.ListResources(context.Background(), request)
				Expect(err).To(BeNil())

				actualRequest := esTransport.receivedHttpRequests[0]
				search := readEsSearchResponse(actualRequest)

				Expect(search.Query).To(Equal(expectedQuery))
			})
		})

		When("elasticsearch returns with an error", func() {
			BeforeEach(func() {
				esTransport.preparedHttpResponses[0] = &http.Response{
					StatusCode: http.StatusInternalServerError,
				}

				listResourcesResponse, listResourcesResponseErr = server.ListResources(context.Background(), request)
			})

			It("should return the error", func() {
				Expect(listResourcesResponseErr).ToNot(BeNil())
			})

			It("should not return a protobuf response", func() {
				Expect(listResourcesResponse).To(BeNil())
			})
		})

		When("listing resources with pagination", func() {
			var (
				actualResponse *pb.ListResourcesResponse
				actualError    error

				expectedPageToken string
				expectedPageSize  int32
				expectedPitId     string
				expectedFrom      int
			)

			BeforeEach(func() {
				expectedPageSize = int32(gofakeit.Number(5, 20))
				expectedPitId = gofakeit.LetterN(20)
				expectedFrom = gofakeit.Number(int(expectedPageSize), 100)

				request.PageSize = expectedPageSize

				esTransport.preparedHttpResponses[0].Body = createPaginatedEsSearchResponse(occurrences, gofakeit.Number(1000, 2000))
			})

			JustBeforeEach(func() {
				actualResponse, actualError = server.ListResources(context.Background(), request)
			})

			When("a page token is not specified", func() {
				BeforeEach(func() {
					esTransport.preparedHttpResponses = append([]*http.Response{
						{
							StatusCode: http.StatusOK,
							Body: structToJsonBody(&esutil.ESPitResponse{
								Id: expectedPitId,
							}),
						},
					}, esTransport.preparedHttpResponses...)
				})

				It("should create a PIT in Elasticsearch", func() {
					Expect(esTransport.receivedHttpRequests[0].URL.Path).To(Equal(fmt.Sprintf("/%s/_pit", rodeElasticsearchOccurrencesAlias)))
					Expect(esTransport.receivedHttpRequests[0].Method).To(Equal(http.MethodPost))
					Expect(esTransport.receivedHttpRequests[0].URL.Query().Get("keep_alive")).To(Equal("5m"))
				})

				It("should query using the PIT", func() {
					Expect(esTransport.receivedHttpRequests[1].URL.Path).To(Equal("/_search"))
					Expect(esTransport.receivedHttpRequests[1].Method).To(Equal(http.MethodGet))
					request := readEsSearchResponse(esTransport.receivedHttpRequests[1])
					Expect(request.Pit.Id).To(Equal(expectedPitId))
					Expect(request.Pit.KeepAlive).To(Equal("5m"))
				})

				It("should not return an error", func() {
					Expect(actualError).To(BeNil())
				})

				It("should return the next page token", func() {
					nextPitId, nextFrom, err := esutil.ParsePageToken(actualResponse.NextPageToken)

					Expect(err).ToNot(HaveOccurred())
					Expect(nextPitId).To(Equal(expectedPitId))
					Expect(nextFrom).To(BeEquivalentTo(expectedPageSize))
				})
			})

			When("a valid token is specified", func() {
				BeforeEach(func() {
					expectedPageToken = esutil.CreatePageToken(expectedPitId, expectedFrom)

					request.PageToken = expectedPageToken
				})

				It("should query Elasticsearch using the PIT", func() {
					Expect(esTransport.receivedHttpRequests[0].URL.Path).To(Equal("/_search"))
					Expect(esTransport.receivedHttpRequests[0].Method).To(Equal(http.MethodGet))
					request := readEsSearchResponse(esTransport.receivedHttpRequests[0])
					Expect(request.Pit.Id).To(Equal(expectedPitId))
					Expect(request.Pit.KeepAlive).To(Equal("5m"))
				})

				It("should return the next page token", func() {
					nextPitId, nextFrom, err := esutil.ParsePageToken(actualResponse.NextPageToken)

					Expect(err).ToNot(HaveOccurred())
					Expect(nextPitId).To(Equal(expectedPitId))
					Expect(nextFrom).To(BeEquivalentTo(expectedPageSize + int32(expectedFrom)))
				})
			})

			When("an invalid token is passed (bad format)", func() {
				BeforeEach(func() {
					request.PageToken = gofakeit.LetterN(10)
				})

				It("should not make any further Elasticsearch queries", func() {
					Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
				})

				It("should return an error", func() {
					Expect(actualError).To(HaveOccurred())
					Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
					Expect(actualResponse).To(BeNil())
				})
			})

			When("an invalid token is passed (bad from)", func() {
				BeforeEach(func() {
					request.PageToken = esutil.CreatePageToken(expectedPitId, expectedFrom) + gofakeit.LetterN(5)
				})

				It("should not make any further Elasticsearch queries", func() {
					Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
				})

				It("should return an error", func() {
					Expect(actualError).To(HaveOccurred())
					Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
					Expect(actualResponse).To(BeNil())
				})
			})
		})
	})

	Context("RegisterCollector", func() {
		var (
			actualRegisterCollectorResponse *pb.RegisterCollectorResponse
			actualRegisterCollectorError    error

			expectedCollectorId              string
			expectedRegisterCollectorRequest *pb.RegisterCollectorRequest

			expectedListNotesResponse *grafeas_proto.ListNotesResponse
			expectedListNotesError    error

			expectedBatchCreateNotesResponse *grafeas_proto.BatchCreateNotesResponse
			expectedBatchCreateNotesError    error

			expectedNotes []*grafeas_proto.Note
		)

		BeforeEach(func() {
			expectedDiscoveryNote := &grafeas_proto.Note{
				ShortDescription: "Harbor Image Scan",
				Kind:             grafeas_common_proto.NoteKind_DISCOVERY,
			}
			expectedAttestationNote := &grafeas_proto.Note{
				ShortDescription: "Harbor Attestation",
				Kind:             grafeas_common_proto.NoteKind_ATTESTATION,
			}
			expectedNotes = []*grafeas_proto.Note{
				expectedDiscoveryNote,
				expectedAttestationNote,
			}
			expectedCollectorId = gofakeit.LetterN(10)
			expectedRegisterCollectorRequest = &pb.RegisterCollectorRequest{
				Id:    expectedCollectorId,
				Notes: expectedNotes,
			}

			// happy path: notes do not already exist
			expectedListNotesResponse = &grafeas_proto.ListNotesResponse{
				Notes: []*grafeas_proto.Note{},
			}
			expectedListNotesError = nil

			// when notes are returned, their name should not be empty
			expectedCreatedDiscoveryNote := deepCopyNote(expectedDiscoveryNote)
			expectedCreatedAttestationNote := deepCopyNote(expectedAttestationNote)

			expectedCreatedDiscoveryNote.Name = fmt.Sprintf("%s/notes/%s", rodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedCreatedDiscoveryNote))
			expectedCreatedAttestationNote.Name = fmt.Sprintf("%s/notes/%s", rodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedCreatedAttestationNote))

			expectedBatchCreateNotesResponse = &grafeas_proto.BatchCreateNotesResponse{
				Notes: []*grafeas_proto.Note{
					expectedCreatedDiscoveryNote,
					expectedCreatedAttestationNote,
				},
			}
			expectedBatchCreateNotesError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.ListNotesReturns(expectedListNotesResponse, expectedListNotesError)
			grafeasClient.BatchCreateNotesReturns(expectedBatchCreateNotesResponse, expectedBatchCreateNotesError)

			actualRegisterCollectorResponse, actualRegisterCollectorError = server.RegisterCollector(ctx, expectedRegisterCollectorRequest)
		})

		It("should search grafeas for the notes", func() {
			Expect(grafeasClient.ListNotesCallCount()).To(Equal(1))

			_, listNotesRequest, _ := grafeasClient.ListNotesArgsForCall(0)
			Expect(listNotesRequest.Parent).To(Equal(rodeProjectSlug))
			Expect(listNotesRequest.Filter).To(Equal(fmt.Sprintf(`name.startsWith("%s/notes/%s-")`, rodeProjectSlug, expectedCollectorId)))
		})

		It("should create the missing notes", func() {
			Expect(grafeasClient.BatchCreateNotesCallCount()).To(Equal(1))

			_, batchCreateNotesRequest, _ := grafeasClient.BatchCreateNotesArgsForCall(0)
			Expect(batchCreateNotesRequest.Parent).To(Equal(rodeProjectSlug))
			Expect(batchCreateNotesRequest.Notes).To(ConsistOf(expectedNotes))
		})

		It("should return the collector's notes", func() {
			Expect(actualRegisterCollectorResponse).ToNot(BeNil())
			Expect(actualRegisterCollectorResponse.Notes).To(HaveLen(len(expectedNotes)))
			for _, note := range expectedNotes {
				note.Name = buildNoteIdFromCollectorId(expectedCollectorId, note)
				Expect(actualRegisterCollectorResponse.Notes).To(ContainElement(note))
			}

			Expect(actualRegisterCollectorError).ToNot(HaveOccurred())
		})

		When("a note already exists", func() {
			BeforeEach(func() {
				expectedNoteThatAlreadyExists := deepCopyNote(expectedNotes[0])
				expectedNoteThatAlreadyExists.Name = fmt.Sprintf("%s/notes/%s", rodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedNoteThatAlreadyExists))

				expectedListNotesResponse.Notes = []*grafeas_proto.Note{
					expectedNoteThatAlreadyExists,
				}
			})

			It("should not attempt to create that note", func() {
				Expect(grafeasClient.BatchCreateNotesCallCount()).To(Equal(1))

				_, batchCreateNotesRequest, _ := grafeasClient.BatchCreateNotesArgsForCall(0)
				Expect(batchCreateNotesRequest.Parent).To(Equal(rodeProjectSlug))
				Expect(batchCreateNotesRequest.Notes).To(HaveLen(1))
				Expect(batchCreateNotesRequest.Notes).To(ContainElement(expectedNotes[1]))
			})
		})

		When("both notes already exist", func() {
			BeforeEach(func() {
				var notesThatAlreadyExist []*grafeas_proto.Note
				for _, note := range expectedNotes {
					noteThatAlreadyExists := deepCopyNote(note)
					noteThatAlreadyExists.Name = fmt.Sprintf("%s/notes/%s", rodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, noteThatAlreadyExists))

					notesThatAlreadyExist = append(notesThatAlreadyExist, noteThatAlreadyExists)
				}

				expectedListNotesResponse.Notes = notesThatAlreadyExist
			})

			It("should not attempt to create any notes", func() {
				Expect(grafeasClient.BatchCreateNotesCallCount()).To(Equal(0))
			})
		})
	})

	Context("CreateNote", func() {
		var (
			actualNote  *grafeas_proto.Note
			actualError error

			expectedNote   *grafeas_proto.Note
			expectedNoteId string

			expectedCreateNoteRequest *pb.CreateNoteRequest
			expectedCreateNoteError   error
		)

		BeforeEach(func() {
			expectedNote = &grafeas_proto.Note{
				ShortDescription: gofakeit.LetterN(10),
				Kind:             grafeas_common_proto.NoteKind_DISCOVERY,
			}
			expectedNoteId = gofakeit.LetterN(10)

			expectedCreateNoteRequest = &pb.CreateNoteRequest{
				NoteId: expectedNoteId,
				Note:   expectedNote,
			}

			expectedCreateNoteError = nil
		})

		JustBeforeEach(func() {
			grafeasClient.CreateNoteReturns(expectedNote, expectedCreateNoteError)

			actualNote, actualError = server.CreateNote(ctx, expectedCreateNoteRequest)
		})

		It("should invoke the grafeas CreateNote rpc and return its result", func() {
			Expect(grafeasClient.CreateNoteCallCount()).To(Equal(1))

			_, createNoteRequest, _ := grafeasClient.CreateNoteArgsForCall(0)
			Expect(createNoteRequest.NoteId).To(Equal(expectedNoteId))
			Expect(createNoteRequest.Note).To(Equal(expectedNote))
			Expect(createNoteRequest.Parent).To(Equal(rodeProjectSlug))

			Expect(actualNote).To(Equal(expectedNote))
			Expect(actualError).ToNot(HaveOccurred())
		})

		When("the grafeas CreateNote rpc fails", func() {
			BeforeEach(func() {
				expectedCreateNoteError = errors.New("error creating note")
				expectedNote = nil
			})

			It("should return an error", func() {
				Expect(actualError).To(HaveOccurred())
				Expect(actualNote).To(BeNil())
			})
		})
	})

	//When("creating multiple policies sequentially", func() {
	//	var (
	//		policyEntityOne   *pb.PolicyEntity
	//		policyEntityTwo   *pb.PolicyEntity
	//		policyResponseOne *pb.Policy
	//		policyResponseTwo *pb.Policy
	//		err               error
	//	)
	//
	//	BeforeEach(func() {
	//		policyEntityOne = createRandomPolicyEntity(goodPolicy)
	//		policyEntityTwo = createRandomPolicyEntity(goodPolicy)
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//		}
	//		policyResponseOne, err = server.CreatePolicy(context.Background(), policyEntityOne)
	//		policyResponseTwo, err = server.CreatePolicy(context.Background(), policyEntityTwo)
	//	})
	//
	//	When("attempting to list the policies", func() {
	//		var (
	//			listRequest   *pb.ListPoliciesRequest
	//			listResponse  *pb.ListPoliciesResponse
	//			policiesList  []*pb.Policy
	//			err           error
	//			filter        string
	//			expectedQuery *filtering.Query
	//		)
	//
	//		BeforeEach(func() {
	//			policiesList = append(policiesList, policyResponseOne)
	//			policiesList = append(policiesList, policyResponseOne)
	//
	//			filter = `name=="abc"`
	//			expectedQuery := &filtering.Query{
	//				Term: &filtering.Term{
	//					"name": "abc",
	//				},
	//			}
	//
	//			listRequest = &pb.ListPoliciesRequest{Filter: filter}
	//			esTransport.preparedHttpResponses = []*http.Response{
	//				{
	//					StatusCode: http.StatusOK,
	//					Body:       createEsSearchResponseForPolicy(policiesList),
	//				},
	//				{
	//					StatusCode: http.StatusOK,
	//					Body:       createEsSearchResponseForPolicy(policiesList),
	//				},
	//			}
	//
	//			mockFilterer.EXPECT().ParseExpression(gomock.Any()).Return(expectedQuery, nil)
	//			listResponse, err = server.ListPolicies(context.Background(), listRequest)
	//		})
	//
	//		It("should not return an error", func() {
	//			Expect(err).To(Not(HaveOccurred()))
	//		})
	//
	//		It("should have listed 4 different policies", func() {
	//			Expect(listResponse.Policies).To(HaveLen(4))
	//		})
	//
	//		It("should have generated a filter query", func() {
	//			actualRequest := esTransport.receivedHttpRequests[0]
	//			search := readEsSearchResponse(actualRequest)
	//
	//			Expect(search.Query).To(Equal(expectedQuery))
	//		})
	//
	//		It("should have generated a filter query", func() {
	//			actualRequest := esTransport.receivedHttpRequests[0]
	//			search := readEsSearchResponse(actualRequest)
	//
	//			Expect(search.Query).To(Equal(expectedQuery))
	//		})
	//	})
	//})

	//When("listing policies with pagination", func() {
	//	var (
	//		listRequest  *pb.ListPoliciesRequest
	//		listResponse *pb.ListPoliciesResponse
	//		actualError  error
	//
	//		expectedPageToken string
	//		expectedPageSize  int32
	//		expectedPitId     string
	//		expectedFrom      int
	//	)
	//
	//	BeforeEach(func() {
	//		expectedPageSize = int32(gofakeit.Number(5, 20))
	//		expectedPitId = gofakeit.LetterN(20)
	//		expectedFrom = gofakeit.Number(int(expectedPageSize), 100)
	//
	//		listRequest = &pb.ListPoliciesRequest{
	//			PageSize: expectedPageSize,
	//		}
	//
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//				Body: createPaginatedEsSearchResponseForPolicy([]*pb.Policy{
	//					{
	//						Id:     gofakeit.UUID(),
	//						Policy: createRandomPolicyEntity(goodPolicy),
	//					},
	//				}, gofakeit.Number(1000, 10000)),
	//			},
	//		}
	//	})
	//
	//	JustBeforeEach(func() {
	//		listResponse, actualError = server.ListPolicies(context.Background(), listRequest)
	//	})
	//
	//	When("a page token is not specified", func() {
	//		BeforeEach(func() {
	//			esTransport.preparedHttpResponses = append([]*http.Response{
	//				{
	//					StatusCode: http.StatusOK,
	//					Body: structToJsonBody(&esutil.ESPitResponse{
	//						Id: expectedPitId,
	//					}),
	//				},
	//			}, esTransport.preparedHttpResponses...)
	//		})
	//
	//		It("should create a PIT in Elasticsearch", func() {
	//			Expect(esTransport.receivedHttpRequests[0].URL.Path).To(Equal(fmt.Sprintf("/%s/_pit", expectedPoliciesAlias)))
	//			Expect(esTransport.receivedHttpRequests[0].Method).To(Equal(http.MethodPost))
	//			Expect(esTransport.receivedHttpRequests[0].URL.Query().Get("keep_alive")).To(Equal("5m"))
	//		})
	//
	//		It("should query using the PIT", func() {
	//			Expect(esTransport.receivedHttpRequests[1].URL.Path).To(Equal("/_search"))
	//			Expect(esTransport.receivedHttpRequests[1].Method).To(Equal(http.MethodGet))
	//			request := readEsSearchResponse(esTransport.receivedHttpRequests[1])
	//			Expect(request.Pit.Id).To(Equal(expectedPitId))
	//			Expect(request.Pit.KeepAlive).To(Equal("5m"))
	//		})
	//
	//		It("should not return an error", func() {
	//			Expect(actualError).To(BeNil())
	//		})
	//
	//		It("should return the next page token", func() {
	//			nextPitId, nextFrom, err := esutil.ParsePageToken(listResponse.NextPageToken)
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(nextPitId).To(Equal(expectedPitId))
	//			Expect(nextFrom).To(BeEquivalentTo(expectedPageSize))
	//		})
	//	})
	//
	//	When("a valid token is specified", func() {
	//		BeforeEach(func() {
	//			expectedPageToken = esutil.CreatePageToken(expectedPitId, expectedFrom)
	//
	//			listRequest.PageToken = expectedPageToken
	//		})
	//
	//		It("should query Elasticsearch using the PIT", func() {
	//			Expect(esTransport.receivedHttpRequests[0].URL.Path).To(Equal("/_search"))
	//			Expect(esTransport.receivedHttpRequests[0].Method).To(Equal(http.MethodGet))
	//			request := readEsSearchResponse(esTransport.receivedHttpRequests[0])
	//			Expect(request.Pit.Id).To(Equal(expectedPitId))
	//			Expect(request.Pit.KeepAlive).To(Equal("5m"))
	//		})
	//
	//		It("should return the next page token", func() {
	//			nextPitId, nextFrom, err := esutil.ParsePageToken(listResponse.NextPageToken)
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(nextPitId).To(Equal(expectedPitId))
	//			Expect(nextFrom).To(BeEquivalentTo(int(expectedPageSize) + expectedFrom))
	//		})
	//
	//		When("the user reaches the last page of results", func() {
	//			BeforeEach(func() {
	//				esTransport.preparedHttpResponses[0].Body = createPaginatedEsSearchResponseForPolicy([]*pb.Policy{
	//					{
	//						Id:     gofakeit.UUID(),
	//						Policy: createRandomPolicyEntity(goodPolicy),
	//					},
	//				}, gofakeit.Number(1, int(expectedPageSize)+expectedFrom-1))
	//			})
	//
	//			It("should return an empty next page token", func() {
	//				Expect(listResponse.NextPageToken).To(Equal(""))
	//			})
	//		})
	//	})
	//
	//	When("an invalid token is passed (bad format)", func() {
	//		BeforeEach(func() {
	//			listRequest.PageToken = gofakeit.LetterN(10)
	//		})
	//
	//		It("should not make any further Elasticsearch queries", func() {
	//			Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
	//		})
	//
	//		It("should return an error", func() {
	//			Expect(actualError).To(HaveOccurred())
	//			Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
	//			Expect(listResponse).To(BeNil())
	//		})
	//	})
	//
	//	When("an invalid token is passed (bad from)", func() {
	//		BeforeEach(func() {
	//			listRequest.PageToken = esutil.CreatePageToken(expectedPitId, expectedFrom) + gofakeit.LetterN(5)
	//		})
	//
	//		It("should not make any further Elasticsearch queries", func() {
	//			Expect(esTransport.receivedHttpRequests).To(HaveLen(0))
	//		})
	//
	//		It("should return an error", func() {
	//			Expect(actualError).To(HaveOccurred())
	//			Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
	//			Expect(listResponse).To(BeNil())
	//		})
	//	})
	//})

	//When("attempting to list an empty policy index", func() {
	//	var (
	//		listRequest          *pb.ListPoliciesRequest
	//		policiesList         []*pb.Policy
	//		listPoliciesResponse *pb.ListPoliciesResponse
	//		err                  error
	//	)
	//	BeforeEach(func() {
	//		listRequest = &pb.ListPoliciesRequest{}
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//				Body:       createEsSearchResponseForPolicy(policiesList),
	//			},
	//		}
	//		listPoliciesResponse, err = server.ListPolicies(context.Background(), listRequest)
	//	})
	//	It("should return an error", func() {
	//		Expect(err).To(Not(HaveOccurred()))
	//	})
	//	It("should return an empty list", func() {
	//		Expect(len(listPoliciesResponse.Policies)).To(Equal(0))
	//	})
	//
	//})

	//When("attempting to list policies and elasticsearch is unreachable", func() {
	//	var (
	//		listRequest  *pb.ListPoliciesRequest
	//		policiesList []*pb.Policy
	//		err          error
	//	)
	//	BeforeEach(func() {
	//		listRequest = &pb.ListPoliciesRequest{}
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusInternalServerError,
	//				Body:       createEsSearchResponseForPolicy(policiesList),
	//			},
	//		}
	//		_, err = server.ListPolicies(context.Background(), listRequest)
	//	})
	//	It("should return an error", func() {
	//		Expect(err).To(HaveOccurred())
	//	})
	//
	//})

	//When("creating an unparseable policy", func() {
	//	var (
	//		policyEntity   *pb.PolicyEntity
	//		policyResponse *pb.Policy
	//		err            error
	//	)
	//
	//	BeforeEach(func() {
	//		policyEntity = createRandomPolicyEntity(unparseablePolicy)
	//		policyResponse, err = server.CreatePolicy(context.Background(), policyEntity)
	//	})
	//
	//	It("should throw a compilation error", func() {
	//		Expect(policyResponse).To(BeNil())
	//		Expect(err).To(HaveOccurred())
	//	})
	//})

	//When("creating an compilable policy with missing Rode Requirements", func() {
	//	var (
	//		policyEntity   *pb.PolicyEntity
	//		policyResponse *pb.Policy
	//		err            error
	//	)
	//
	//	BeforeEach(func() {
	//		policyEntity = createRandomPolicyEntity(compilablePolicyMissingRodeFields)
	//		policyResponse, err = server.CreatePolicy(context.Background(), policyEntity)
	//	})
	//
	//	It("should throw a compilation error", func() {
	//		Expect(policyResponse).To(BeNil())
	//		Expect(err).To(HaveOccurred())
	//	})
	//})

	//When("creating an compilable policy with missing results return", func() {
	//	var (
	//		policyEntity   *pb.PolicyEntity
	//		policyResponse *pb.Policy
	//		err            error
	//	)
	//
	//	BeforeEach(func() {
	//		policyEntity = createRandomPolicyEntity(compilablePolicyMissingResultsReturn)
	//		policyResponse, err = server.CreatePolicy(context.Background(), policyEntity)
	//	})
	//
	//	It("should throw a compilation error", func() {
	//		Expect(policyResponse).To(BeNil())
	//		Expect(err).To(HaveOccurred())
	//	})
	//})

	//When("creating an compilable policy with missing required fields in the result object", func() {
	//	var (
	//		policyEntity   *pb.PolicyEntity
	//		policyResponse *pb.Policy
	//		err            error
	//	)
	//
	//	BeforeEach(func() {
	//		policyEntity = createRandomPolicyEntity(compilablePolicyMissingResultsFields)
	//		policyResponse, err = server.CreatePolicy(context.Background(), policyEntity)
	//	})
	//
	//	It("should throw a compilation error", func() {
	//		Expect(policyResponse).To(BeNil())
	//		Expect(err).To(HaveOccurred())
	//	})
	//})
	//
	//When("validating a good policy", func() {
	//	var (
	//		validatePolicyRequest  *pb.ValidatePolicyRequest
	//		validatePolicyResponse *pb.ValidatePolicyResponse
	//		err                    error
	//	)
	//
	//	BeforeEach(func() {
	//		validatePolicyRequest = &pb.ValidatePolicyRequest{Policy: goodPolicy}
	//		validatePolicyResponse, err = server.ValidatePolicy(context.Background(), validatePolicyRequest)
	//	})
	//
	//	It("should not throw an error", func() {
	//		Expect(err).To(Not(HaveOccurred()))
	//	})
	//	It("should return a successful compilation", func() {
	//		Expect(validatePolicyResponse.Compile).To(BeTrue())
	//	})
	//	It("should return an empty error array", func() {
	//		Expect(validatePolicyResponse.Errors).To(BeEmpty())
	//	})
	//})
	//
	//When("validating an empty policy", func() {
	//	var (
	//		validatePolicyRequest *pb.ValidatePolicyRequest
	//		err                   error
	//	)
	//
	//	BeforeEach(func() {
	//		validatePolicyRequest = &pb.ValidatePolicyRequest{Policy: ""}
	//		_, err = server.ValidatePolicy(context.Background(), validatePolicyRequest)
	//	})
	//
	//	It("should throw an error", func() {
	//		Expect(err).To(HaveOccurred())
	//	})
	//})
	//
	//When("validating an uncompilable policy", func() {
	//	var (
	//		validatePolicyRequest  *pb.ValidatePolicyRequest
	//		validatePolicyResponse *pb.ValidatePolicyResponse
	//		err                    error
	//	)
	//
	//	BeforeEach(func() {
	//		validatePolicyRequest = &pb.ValidatePolicyRequest{Policy: uncompilablePolicy}
	//		validatePolicyResponse, err = server.ValidatePolicy(context.Background(), validatePolicyRequest)
	//	})
	//
	//	It("should throw an error", func() {
	//		Expect(err).To(HaveOccurred())
	//	})
	//	It("should return an unsuccessful compilation", func() {
	//		Expect(validatePolicyResponse.Compile).To(BeFalse())
	//	})
	//	It("should not return an empty error array", func() {
	//		Expect(len(validatePolicyResponse.Errors)).To(Not(Equal(0)))
	//	})
	//})

	//When("updating the name of a policy", func() {
	//	var (
	//		createPolicyRequest  *pb.PolicyEntity
	//		createPolicyResponse *pb.Policy
	//		updatePolicyRequest  *pb.UpdatePolicyRequest
	//		updatePolicyResponse *pb.Policy
	//		err                  error
	//		initialPolicyName    string
	//	)
	//
	//	BeforeEach(func() {
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//		}
	//
	//		createPolicyRequest = createRandomPolicyEntity(goodPolicy)
	//		createPolicyResponse, _ = server.CreatePolicy(context.Background(), createPolicyRequest)
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//				Body:       createEsSearchResponseForPolicy([]*pb.Policy{createPolicyResponse}),
	//			},
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//		}
	//
	//		initialPolicyName = createPolicyResponse.Policy.Name
	//		updatePolicyRequest = &pb.UpdatePolicyRequest{
	//			Id: createPolicyResponse.Id,
	//			Policy: &pb.PolicyEntity{
	//				Name: "random name",
	//			},
	//			UpdateMask: &fieldmaskpb.FieldMask{
	//				Paths: []string{"name"},
	//			},
	//		}
	//		updatePolicyResponse, err = server.UpdatePolicy(context.Background(), updatePolicyRequest)
	//	})
	//
	//	It("should not throw an error", func() {
	//		Expect(err).To(Not(HaveOccurred()))
	//	})
	//	It("should now have a new policy name", func() {
	//		Expect(initialPolicyName).To(Not(Equal(updatePolicyResponse.Policy.Name)))
	//	})
	//	When("the original policy does not exist", func() {
	//		BeforeEach(func() {
	//			esTransport.preparedHttpResponses = []*http.Response{
	//				{
	//					StatusCode: http.StatusOK,
	//					Body:       createEsSearchResponseForPolicy([]*pb.Policy{}),
	//				},
	//				{
	//					StatusCode: http.StatusOK,
	//				},
	//			}
	//			updatePolicyRequest = &pb.UpdatePolicyRequest{
	//				Id: gofakeit.LetterN(10),
	//				Policy: &pb.PolicyEntity{
	//					Name: "random name",
	//				},
	//				UpdateMask: &fieldmaskpb.FieldMask{
	//					Paths: []string{"name"},
	//				},
	//			}
	//			updatePolicyResponse, err = server.UpdatePolicy(context.Background(), updatePolicyRequest)
	//		})
	//		It("should throw an error", func() {
	//			Expect(err).To(HaveOccurred())
	//		})
	//	})
	//})

	//When("updating the rego content of a policy", func() {
	//	var (
	//		createPolicyRequest  *pb.PolicyEntity
	//		createPolicyResponse *pb.Policy
	//		updatePolicyRequest  *pb.UpdatePolicyRequest
	//		updatePolicyResponse *pb.Policy
	//		err                  error
	//	)
	//
	//	BeforeEach(func() {
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//		}
	//
	//		createPolicyRequest = createRandomPolicyEntity(goodPolicy)
	//		createPolicyResponse, _ = server.CreatePolicy(context.Background(), createPolicyRequest)
	//		esTransport.preparedHttpResponses = []*http.Response{
	//			{
	//				StatusCode: http.StatusOK,
	//				Body:       createEsSearchResponseForPolicy([]*pb.Policy{createPolicyResponse}),
	//			},
	//			{
	//				StatusCode: http.StatusOK,
	//			},
	//		}
	//
	//	})
	//	When("the policy does not compile", func() {
	//		BeforeEach(func() {
	//			updatePolicyRequest = &pb.UpdatePolicyRequest{
	//				Id: createPolicyResponse.Id,
	//				Policy: &pb.PolicyEntity{
	//					RegoContent: uncompilablePolicy,
	//				},
	//				UpdateMask: &fieldmaskpb.FieldMask{
	//					Paths: []string{"rego_content"},
	//				},
	//			}
	//			updatePolicyResponse, err = server.UpdatePolicy(context.Background(), updatePolicyRequest)
	//		})
	//		It("should throw an error ", func() {
	//			Expect(err).To(HaveOccurred())
	//			Expect(updatePolicyResponse).To(BeNil())
	//		})
	//	})
	//
	//})
})

func createRandomOccurrence(kind grafeas_common_proto.NoteKind) *grafeas_proto.Occurrence {
	return &grafeas_proto.Occurrence{
		Name: gofakeit.LetterN(10),
		Resource: &grafeas_proto.Resource{
			Uri: fmt.Sprintf("%s@sha256:%s", gofakeit.URL(), gofakeit.LetterN(10)),
		},
		NoteName:    gofakeit.LetterN(10),
		Kind:        kind,
		Remediation: gofakeit.LetterN(10),
		CreateTime:  timestamppb.New(gofakeit.Date()),
		UpdateTime:  timestamppb.New(gofakeit.Date()),
		Details:     nil,
	}
}

func structToJsonBody(i interface{}) io.ReadCloser {
	b, err := json.Marshal(i)
	Expect(err).ToNot(HaveOccurred())

	return io.NopCloser(strings.NewReader(string(b)))
}

func createEsSearchResponse(occurrences []*grafeas_proto.Occurrence) io.ReadCloser {
	return createPaginatedEsSearchResponse(occurrences, len(occurrences))
}

func createPaginatedEsSearchResponse(occurrences []*grafeas_proto.Occurrence, totalResults int) io.ReadCloser {
	var occurrenceHits []*esutil.EsSearchResponseHit

	for _, occurrence := range occurrences {
		source, err := protojson.Marshal(proto.MessageV2(occurrence))
		Expect(err).To(BeNil())

		response := &esutil.EsSearchResponseHit{
			ID:     gofakeit.UUID(),
			Source: source,
		}

		occurrenceHits = append(occurrenceHits, response)
	}

	response := &esutil.EsSearchResponse{
		Hits: &esutil.EsSearchResponseHits{
			Total: &esutil.EsSearchResponseTotal{
				Value: totalResults,
			},
			Hits: occurrenceHits,
		},
		Took: gofakeit.Number(1, 10),
	}

	responseBody, err := json.Marshal(response)
	Expect(err).To(BeNil())

	return io.NopCloser(bytes.NewReader(responseBody))
}

func createPaginatedEsSearchResponseForPolicy(policies []*pb.Policy, totalValue int) io.ReadCloser {
	var policyHits []*esutil.EsSearchResponseHit

	for _, occurrence := range policies {
		source, err := protojson.Marshal(proto.MessageV2(occurrence))
		Expect(err).To(BeNil())

		response := &esutil.EsSearchResponseHit{
			ID:     gofakeit.UUID(),
			Source: source,
		}

		policyHits = append(policyHits, response)
	}

	response := &esutil.EsSearchResponse{
		Hits: &esutil.EsSearchResponseHits{
			Total: &esutil.EsSearchResponseTotal{
				Value: totalValue,
			},
			Hits: policyHits,
		},
		Took: gofakeit.Number(1, 10),
	}

	responseBody, err := json.Marshal(response)
	Expect(err).To(BeNil())

	return io.NopCloser(bytes.NewReader(responseBody))
}

func readEsSearchResponse(request *http.Request) *esutil.EsSearch {
	search := &esutil.EsSearch{}
	readResponseBody(request, search)

	return search
}

func readResponseBody(request *http.Request, v interface{}) {
	requestBody, err := io.ReadAll(request.Body)
	Expect(err).To(BeNil())

	err = json.Unmarshal(requestBody, v)
	Expect(err).To(BeNil())
}

func getGRPCStatusFromError(err error) *status.Status {
	s, ok := status.FromError(err)
	Expect(ok).To(BeTrue(), "Expected error to be a gRPC status")

	return s
}

func deepCopyNote(note *grafeas_proto.Note) *grafeas_proto.Note {
	return &grafeas_proto.Note{
		Name:             note.Name,
		ShortDescription: note.ShortDescription,
		LongDescription:  note.LongDescription,
		Kind:             note.Kind,
	}
}
