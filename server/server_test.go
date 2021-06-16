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
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rode/rode/pkg/constants"
	"github.com/rode/rode/pkg/grafeas/grafeasfakes"

	"github.com/rode/rode/pkg/policy/policyfakes"
	"github.com/rode/rode/pkg/resource/resourcefakes"

	"github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	immocks "github.com/rode/es-index-manager/mocks"
	"github.com/rode/rode/mocks"
	pb "github.com/rode/rode/proto/v1alpha1"
	grafeas_common_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/common_go_proto"
	"github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grafeas_project_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/project_go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ = Describe("rode server", func() {
	var (
		server                pb.RodeServer
		grafeasClient         *mocks.FakeGrafeasV1Beta1Client
		grafeasProjectsClient *mocks.FakeProjectsClient
		grafeasExtensions     *grafeasfakes.FakeExtensions
		resourceManager       *resourcefakes.FakeManager
		policyManager         *policyfakes.FakeManager
		policyGroupManager    *policyfakes.FakePolicyGroupManager
		indexManager          *immocks.FakeIndexManager
		ctx                   context.Context

		expectedPoliciesIndex        string
		expectedPoliciesAlias        string
		expectedGenericResourceIndex string
		expectedGenericResourceAlias string
		expectedEnvironmentsIndex    string
		expectedEnvironmentsAlias    string
	)

	BeforeEach(func() {
		grafeasClient = &mocks.FakeGrafeasV1Beta1Client{}
		grafeasProjectsClient = &mocks.FakeProjectsClient{}
		grafeasExtensions = &grafeasfakes.FakeExtensions{}
		resourceManager = &resourcefakes.FakeManager{}
		policyGroupManager = &policyfakes.FakePolicyGroupManager{}

		expectedPoliciesIndex = gofakeit.LetterN(10)
		expectedPoliciesAlias = gofakeit.LetterN(10)
		expectedGenericResourceIndex = gofakeit.LetterN(10)
		expectedGenericResourceAlias = gofakeit.LetterN(10)
		expectedEnvironmentsIndex = gofakeit.LetterN(10)
		expectedEnvironmentsAlias = gofakeit.LetterN(10)
		indexManager = &immocks.FakeIndexManager{}

		indexManager.AliasNameStub = func(documentKind, _ string) string {
			return map[string]string{
				constants.GenericResourcesDocumentKind: expectedGenericResourceAlias,
				constants.PoliciesDocumentKind:         expectedPoliciesAlias,
				constants.PolicyGroupsDocumentKind:     expectedEnvironmentsAlias,
			}[documentKind]
		}

		indexManager.IndexNameStub = func(documentKind, _ string) string {
			return map[string]string{
				constants.GenericResourcesDocumentKind: expectedGenericResourceIndex,
				constants.PoliciesDocumentKind:         expectedPoliciesIndex,
				constants.PolicyGroupsDocumentKind:     expectedEnvironmentsIndex,
			}[documentKind]
		}

		ctx = context.Background()

		// not using the constructor as it has side effects. side effects are tested under the "initialize" context
		server = &rodeServer{
			logger:            logger,
			grafeasCommon:     grafeasClient,
			grafeasProjects:   grafeasProjectsClient,
			grafeasExtensions: grafeasExtensions,
			resourceManager:   resourceManager,
			indexManager:      indexManager,
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

			actualRodeServer, actualError = NewRodeServer(logger, grafeasClient, grafeasProjectsClient, grafeasExtensions, resourceManager, indexManager, policyManager, policyGroupManager)
		})

		It("should check if the rode project exists", func() {
			Expect(grafeasProjectsClient.GetProjectCallCount()).To(Equal(1))

			_, getProjectRequest, _ := grafeasProjectsClient.GetProjectArgsForCall(0)
			Expect(getProjectRequest.Name).To(Equal(constants.RodeProjectSlug))
		})

		// happy path: project already exists
		It("should not create a project", func() {
			Expect(grafeasProjectsClient.CreateProjectCallCount()).To(Equal(0))
		})

		It("should initialize the index manager", func() {
			Expect(indexManager.InitializeCallCount()).To(Equal(1))
		})

		It("should create the application indices", func() {
			Expect(indexManager.CreateIndexCallCount()).To(Equal(3))
		})

		It("should create an index for policies", func() {
			_, actualIndexName, actualAliasName, documentKind := indexManager.CreateIndexArgsForCall(0)

			Expect(actualIndexName).To(Equal(expectedPoliciesIndex))
			Expect(actualAliasName).To(Equal(expectedPoliciesAlias))
			Expect(documentKind).To(Equal(constants.PoliciesDocumentKind))
		})

		It("should create an index for generic resources", func() {
			_, actualIndexName, actualAliasName, documentKind := indexManager.CreateIndexArgsForCall(1)

			Expect(actualIndexName).To(Equal(expectedGenericResourceIndex))
			Expect(actualAliasName).To(Equal(expectedGenericResourceAlias))
			Expect(documentKind).To(Equal(constants.GenericResourcesDocumentKind))
		})

		It("should create an index for environments", func() {
			_, actualIndexName, actualAliasName, documentKind := indexManager.CreateIndexArgsForCall(2)

			Expect(actualIndexName).To(Equal(expectedEnvironmentsIndex))
			Expect(actualAliasName).To(Equal(expectedEnvironmentsAlias))
			Expect(documentKind).To(Equal(constants.PolicyGroupsDocumentKind))
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
				Expect(indexManager.CreateIndexCallCount()).To(Equal(0))
			})
		})

		When("the rode project does not exist", func() {
			BeforeEach(func() {
				expectedGetProjectError = status.Error(codes.NotFound, "not found")
			})

			It("should create the rode project", func() {
				Expect(grafeasProjectsClient.CreateProjectCallCount()).To(Equal(1))

				_, createProjectRequest, _ := grafeasProjectsClient.CreateProjectArgsForCall(0)
				Expect(createProjectRequest.Project.Name).To(Equal(constants.RodeProjectSlug))
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
					Expect(indexManager.CreateIndexCallCount()).To(Equal(0))
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
			expectedBuildNoteName string
			expectedNoteName      string

			listNotesResponse *grafeas_proto.ListNotesResponse
			listNotesError    error

			occurrences                           []*grafeas_proto.Occurrence
			nextPageToken                         string
			listVersionedResourceOccurrencesError error

			expectedPageToken   string
			expectedResourceUri string
			expectedPageSize    int32
			request             *pb.ListVersionedResourceOccurrencesRequest
			actualResponse      *pb.ListVersionedResourceOccurrencesResponse
			actualError         error
		)

		BeforeEach(func() {
			expectedResourceUri = gofakeit.URL()
			nextPageToken = gofakeit.Word()
			expectedPageToken = gofakeit.Word()
			expectedPageSize = gofakeit.Int32()
			expectedBuildNoteName = gofakeit.LetterN(10)
			expectedNoteName = gofakeit.LetterN(10)

			request = &pb.ListVersionedResourceOccurrencesRequest{
				ResourceUri: expectedResourceUri,
				PageToken:   expectedPageToken,
				PageSize:    expectedPageSize,
			}

			occurrences = []*grafeas_proto.Occurrence{
				createRandomOccurrence(grafeas_common_proto.NoteKind_VULNERABILITY),
				createRandomOccurrence(grafeas_common_proto.NoteKind_BUILD),
			}
			occurrences[0].NoteName = expectedNoteName
			occurrences[1].NoteName = expectedBuildNoteName

			listVersionedResourceOccurrencesError = nil

			listNotesResponse = &grafeas_proto.ListNotesResponse{
				Notes: []*grafeas_proto.Note{
					{
						Name: expectedBuildNoteName,
					},
					{
						Name: expectedNoteName,
					},
				},
			}
			listNotesError = nil
		})

		JustBeforeEach(func() {
			grafeasExtensions.ListVersionedResourceOccurrencesReturns(occurrences, nextPageToken, listVersionedResourceOccurrencesError)
			grafeasClient.ListNotesReturns(listNotesResponse, listNotesError)

			actualResponse, actualError = server.ListVersionedResourceOccurrences(ctx, request)
		})

		It("should delegate the occurrence search to the grafeas extensions", func() {
			Expect(grafeasExtensions.ListVersionedResourceOccurrencesCallCount()).To(Equal(1))

			_, resourceUri, pageToken, pageSize := grafeasExtensions.ListVersionedResourceOccurrencesArgsForCall(0)

			Expect(resourceUri).To(Equal(expectedResourceUri))
			Expect(pageToken).To(Equal(expectedPageToken))
			Expect(pageSize).To(Equal(expectedPageSize))
		})

		It("should return the ocurrences", func() {
			Expect(actualResponse.Occurrences).To(Equal(occurrences))
			Expect(actualResponse.NextPageToken).To(Equal(nextPageToken))
			Expect(actualResponse.RelatedNotes).To(BeNil())
		})

		When("an error occurs while listing occurrences", func() {
			BeforeEach(func() {
				listVersionedResourceOccurrencesError = errors.New("error listing occurrences")
			})

			It("should return an error", func() {
				Expect(actualResponse).To(BeNil())
				Expect(actualError).To(HaveOccurred())
				Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
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

		When("the client wants to fetch related notes", func() {
			BeforeEach(func() {
				request.FetchRelatedNotes = true
			})

			It("should list notes with a filter expression matching all note names", func() {
				_, listNotesRequest, _ := grafeasClient.ListNotesArgsForCall(0)

				filters := strings.Split(listNotesRequest.Filter, " || ")
				Expect(filters).To(HaveLen(2))
				Expect(filters).To(ConsistOf(fmt.Sprintf(`"name" == "%s"`, expectedBuildNoteName), fmt.Sprintf(`"name" == "%s"`, expectedNoteName)))

				Expect(listNotesRequest.Parent).To(Equal(constants.RodeProjectSlug))
				Expect(listNotesRequest.PageSize).To(BeEquivalentTo(constants.MaxPageSize))
			})

			It("should respond with the related notes", func() {
				Expect(actualResponse.RelatedNotes).To(HaveLen(2))
				Expect(actualResponse.RelatedNotes[expectedBuildNoteName].Name).To(Equal(expectedBuildNoteName))
				Expect(actualResponse.RelatedNotes[expectedNoteName].Name).To(Equal(expectedNoteName))
			})

			When("listing notes fails", func() {
				BeforeEach(func() {
					listNotesError = errors.New("error listing notes")
				})

				It("should return an error", func() {
					Expect(actualResponse).To(BeNil())
					Expect(actualError).To(HaveOccurred())
					Expect(getGRPCStatusFromError(actualError).Code()).To(Equal(codes.Internal))
				})
			})

			When("no occurrences are returned", func() {
				BeforeEach(func() {
					occurrences = []*grafeas_proto.Occurrence{}
				})

				It("should not fetch related notes", func() {
					Expect(grafeasClient.ListNotesCallCount()).To(Equal(0))
					Expect(actualResponse.RelatedNotes).To(HaveLen(0))
				})
			})

			When("several occurrences have the same note references", func() {
				BeforeEach(func() {
					otherBuildOccurrence := createRandomOccurrence(grafeas_common_proto.NoteKind_BUILD)
					otherBuildOccurrence.NoteName = expectedBuildNoteName

					otherOccurrence := createRandomOccurrence(grafeas_common_proto.NoteKind_DISCOVERY)
					otherOccurrence.NoteName = expectedNoteName

					occurrences = append(occurrences, otherBuildOccurrence, otherOccurrence)
				})

				It("should only list unique notes", func() {
					_, listNotesRequest, _ := grafeasClient.ListNotesArgsForCall(0)

					filters := strings.Split(listNotesRequest.Filter, " || ")
					Expect(filters).To(HaveLen(2))
					Expect(filters).To(ConsistOf(fmt.Sprintf(`"name" == "%s"`, expectedBuildNoteName), fmt.Sprintf(`"name" == "%s"`, expectedNoteName)))
				})

				It("should only return unique notes", func() {
					Expect(actualResponse.RelatedNotes).To(HaveLen(2))
					Expect(actualResponse.RelatedNotes[expectedBuildNoteName].Name).To(Equal(expectedBuildNoteName))
					Expect(actualResponse.RelatedNotes[expectedNoteName].Name).To(Equal(expectedNoteName))
				})
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

			Expect(listOccurrencesRequest.Parent).To(Equal(constants.RodeProjectSlug))
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

			expectedCreatedDiscoveryNote.Name = fmt.Sprintf("%s/notes/%s", constants.RodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedCreatedDiscoveryNote))
			expectedCreatedAttestationNote.Name = fmt.Sprintf("%s/notes/%s", constants.RodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedCreatedAttestationNote))

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
			Expect(listNotesRequest.Parent).To(Equal(constants.RodeProjectSlug))
			Expect(listNotesRequest.Filter).To(Equal(fmt.Sprintf(`name.startsWith("%s/notes/%s-")`, constants.RodeProjectSlug, expectedCollectorId)))
		})

		It("should create the missing notes", func() {
			Expect(grafeasClient.BatchCreateNotesCallCount()).To(Equal(1))

			_, batchCreateNotesRequest, _ := grafeasClient.BatchCreateNotesArgsForCall(0)
			Expect(batchCreateNotesRequest.Parent).To(Equal(constants.RodeProjectSlug))
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
				expectedNoteThatAlreadyExists.Name = fmt.Sprintf("%s/notes/%s", constants.RodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, expectedNoteThatAlreadyExists))

				expectedListNotesResponse.Notes = []*grafeas_proto.Note{
					expectedNoteThatAlreadyExists,
				}
			})

			It("should not attempt to create that note", func() {
				Expect(grafeasClient.BatchCreateNotesCallCount()).To(Equal(1))

				_, batchCreateNotesRequest, _ := grafeasClient.BatchCreateNotesArgsForCall(0)
				Expect(batchCreateNotesRequest.Parent).To(Equal(constants.RodeProjectSlug))
				Expect(batchCreateNotesRequest.Notes).To(HaveLen(1))
				Expect(batchCreateNotesRequest.Notes).To(ContainElement(expectedNotes[1]))
			})
		})

		When("both notes already exist", func() {
			BeforeEach(func() {
				var notesThatAlreadyExist []*grafeas_proto.Note
				for _, note := range expectedNotes {
					noteThatAlreadyExists := deepCopyNote(note)
					noteThatAlreadyExists.Name = fmt.Sprintf("%s/notes/%s", constants.RodeProjectSlug, buildNoteIdFromCollectorId(expectedCollectorId, noteThatAlreadyExists))

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
			Expect(createNoteRequest.Parent).To(Equal(constants.RodeProjectSlug))

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
