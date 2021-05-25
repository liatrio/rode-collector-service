// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grafeas_go_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RodeClient is the client API for Rode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RodeClient interface {
	// Create occurrences
	BatchCreateOccurrences(ctx context.Context, in *BatchCreateOccurrencesRequest, opts ...grpc.CallOption) (*BatchCreateOccurrencesResponse, error)
	// Verify that an artifact satisfies a policy
	EvaluatePolicy(ctx context.Context, in *EvaluatePolicyRequest, opts ...grpc.CallOption) (*EvaluatePolicyResponse, error)
	// List resource URI
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
	ListGenericResources(ctx context.Context, in *ListGenericResourcesRequest, opts ...grpc.CallOption) (*ListGenericResourcesResponse, error)
	// ListGenericResourceVersions can be used to list all known versions of a generic resource. Versions will always include
	// the unique identifier (in the case of Docker images, the sha256) and will optionally include any related names (in the
	// case of Docker images, any associated tags for the image).
	ListGenericResourceVersions(ctx context.Context, in *ListGenericResourceVersionsRequest, opts ...grpc.CallOption) (*ListGenericResourceVersionsResponse, error)
	ListVersionedResourceOccurrences(ctx context.Context, in *ListVersionedResourceOccurrencesRequest, opts ...grpc.CallOption) (*ListVersionedResourceOccurrencesResponse, error)
	ListOccurrences(ctx context.Context, in *ListOccurrencesRequest, opts ...grpc.CallOption) (*ListOccurrencesResponse, error)
	UpdateOccurrence(ctx context.Context, in *UpdateOccurrenceRequest, opts ...grpc.CallOption) (*grafeas_go_proto.Occurrence, error)
	CreatePolicy(ctx context.Context, in *PolicyEntity, opts ...grpc.CallOption) (*Policy, error)
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
	ValidatePolicy(ctx context.Context, in *ValidatePolicyRequest, opts ...grpc.CallOption) (*ValidatePolicyResponse, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// RegisterCollector accepts a collector ID and a list of notes that this collector will reference when creating
	// occurrences. The response will contain the notes with the fully qualified note name. This operation is idempotent,
	// so any notes that already exist will not be re-created. Collectors are expected to invoke this RPC each time they
	// start.
	RegisterCollector(ctx context.Context, in *RegisterCollectorRequest, opts ...grpc.CallOption) (*RegisterCollectorResponse, error)
	// CreateNote acts as a simple proxy to the grafeas CreateNote rpc
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*grafeas_go_proto.Note, error)
}

type rodeClient struct {
	cc grpc.ClientConnInterface
}

func NewRodeClient(cc grpc.ClientConnInterface) RodeClient {
	return &rodeClient{cc}
}

func (c *rodeClient) BatchCreateOccurrences(ctx context.Context, in *BatchCreateOccurrencesRequest, opts ...grpc.CallOption) (*BatchCreateOccurrencesResponse, error) {
	out := new(BatchCreateOccurrencesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/BatchCreateOccurrences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) EvaluatePolicy(ctx context.Context, in *EvaluatePolicyRequest, opts ...grpc.CallOption) (*EvaluatePolicyResponse, error) {
	out := new(EvaluatePolicyResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/EvaluatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListGenericResources(ctx context.Context, in *ListGenericResourcesRequest, opts ...grpc.CallOption) (*ListGenericResourcesResponse, error) {
	out := new(ListGenericResourcesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListGenericResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListGenericResourceVersions(ctx context.Context, in *ListGenericResourceVersionsRequest, opts ...grpc.CallOption) (*ListGenericResourceVersionsResponse, error) {
	out := new(ListGenericResourceVersionsResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListGenericResourceVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListVersionedResourceOccurrences(ctx context.Context, in *ListVersionedResourceOccurrencesRequest, opts ...grpc.CallOption) (*ListVersionedResourceOccurrencesResponse, error) {
	out := new(ListVersionedResourceOccurrencesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListVersionedResourceOccurrences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListOccurrences(ctx context.Context, in *ListOccurrencesRequest, opts ...grpc.CallOption) (*ListOccurrencesResponse, error) {
	out := new(ListOccurrencesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListOccurrences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) UpdateOccurrence(ctx context.Context, in *UpdateOccurrenceRequest, opts ...grpc.CallOption) (*grafeas_go_proto.Occurrence, error) {
	out := new(grafeas_go_proto.Occurrence)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/UpdateOccurrence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) CreatePolicy(ctx context.Context, in *PolicyEntity, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/CreatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) ValidatePolicy(ctx context.Context, in *ValidatePolicyRequest, opts ...grpc.CallOption) (*ValidatePolicyResponse, error) {
	out := new(ValidatePolicyResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/ValidatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/UpdatePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) RegisterCollector(ctx context.Context, in *RegisterCollectorRequest, opts ...grpc.CallOption) (*RegisterCollectorResponse, error) {
	out := new(RegisterCollectorResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/RegisterCollector", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rodeClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*grafeas_go_proto.Note, error) {
	out := new(grafeas_go_proto.Note)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/CreateNote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RodeServer is the server API for Rode service.
// All implementations should embed UnimplementedRodeServer
// for forward compatibility
type RodeServer interface {
	// Create occurrences
	BatchCreateOccurrences(context.Context, *BatchCreateOccurrencesRequest) (*BatchCreateOccurrencesResponse, error)
	// Verify that an artifact satisfies a policy
	EvaluatePolicy(context.Context, *EvaluatePolicyRequest) (*EvaluatePolicyResponse, error)
	// List resource URI
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
	ListGenericResources(context.Context, *ListGenericResourcesRequest) (*ListGenericResourcesResponse, error)
	// ListGenericResourceVersions can be used to list all known versions of a generic resource. Versions will always include
	// the unique identifier (in the case of Docker images, the sha256) and will optionally include any related names (in the
	// case of Docker images, any associated tags for the image).
	ListGenericResourceVersions(context.Context, *ListGenericResourceVersionsRequest) (*ListGenericResourceVersionsResponse, error)
	ListVersionedResourceOccurrences(context.Context, *ListVersionedResourceOccurrencesRequest) (*ListVersionedResourceOccurrencesResponse, error)
	ListOccurrences(context.Context, *ListOccurrencesRequest) (*ListOccurrencesResponse, error)
	UpdateOccurrence(context.Context, *UpdateOccurrenceRequest) (*grafeas_go_proto.Occurrence, error)
	CreatePolicy(context.Context, *PolicyEntity) (*Policy, error)
	GetPolicy(context.Context, *GetPolicyRequest) (*Policy, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*empty.Empty, error)
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	ValidatePolicy(context.Context, *ValidatePolicyRequest) (*ValidatePolicyResponse, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error)
	// RegisterCollector accepts a collector ID and a list of notes that this collector will reference when creating
	// occurrences. The response will contain the notes with the fully qualified note name. This operation is idempotent,
	// so any notes that already exist will not be re-created. Collectors are expected to invoke this RPC each time they
	// start.
	RegisterCollector(context.Context, *RegisterCollectorRequest) (*RegisterCollectorResponse, error)
	// CreateNote acts as a simple proxy to the grafeas CreateNote rpc
	CreateNote(context.Context, *CreateNoteRequest) (*grafeas_go_proto.Note, error)
}

// UnimplementedRodeServer should be embedded to have forward compatible implementations.
type UnimplementedRodeServer struct {
}

func (UnimplementedRodeServer) BatchCreateOccurrences(context.Context, *BatchCreateOccurrencesRequest) (*BatchCreateOccurrencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateOccurrences not implemented")
}
func (UnimplementedRodeServer) EvaluatePolicy(context.Context, *EvaluatePolicyRequest) (*EvaluatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvaluatePolicy not implemented")
}
func (UnimplementedRodeServer) ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResources not implemented")
}
func (UnimplementedRodeServer) ListGenericResources(context.Context, *ListGenericResourcesRequest) (*ListGenericResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGenericResources not implemented")
}
func (UnimplementedRodeServer) ListGenericResourceVersions(context.Context, *ListGenericResourceVersionsRequest) (*ListGenericResourceVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGenericResourceVersions not implemented")
}
func (UnimplementedRodeServer) ListVersionedResourceOccurrences(context.Context, *ListVersionedResourceOccurrencesRequest) (*ListVersionedResourceOccurrencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersionedResourceOccurrences not implemented")
}
func (UnimplementedRodeServer) ListOccurrences(context.Context, *ListOccurrencesRequest) (*ListOccurrencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOccurrences not implemented")
}
func (UnimplementedRodeServer) UpdateOccurrence(context.Context, *UpdateOccurrenceRequest) (*grafeas_go_proto.Occurrence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOccurrence not implemented")
}
func (UnimplementedRodeServer) CreatePolicy(context.Context, *PolicyEntity) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedRodeServer) GetPolicy(context.Context, *GetPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedRodeServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedRodeServer) ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedRodeServer) ValidatePolicy(context.Context, *ValidatePolicyRequest) (*ValidatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePolicy not implemented")
}
func (UnimplementedRodeServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedRodeServer) RegisterCollector(context.Context, *RegisterCollectorRequest) (*RegisterCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCollector not implemented")
}
func (UnimplementedRodeServer) CreateNote(context.Context, *CreateNoteRequest) (*grafeas_go_proto.Note, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}

// UnsafeRodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RodeServer will
// result in compilation errors.
type UnsafeRodeServer interface {
	mustEmbedUnimplementedRodeServer()
}

func RegisterRodeServer(s grpc.ServiceRegistrar, srv RodeServer) {
	s.RegisterService(&Rode_ServiceDesc, srv)
}

func _Rode_BatchCreateOccurrences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateOccurrencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).BatchCreateOccurrences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/BatchCreateOccurrences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).BatchCreateOccurrences(ctx, req.(*BatchCreateOccurrencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_EvaluatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).EvaluatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/EvaluatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).EvaluatePolicy(ctx, req.(*EvaluatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListGenericResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGenericResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListGenericResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListGenericResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListGenericResources(ctx, req.(*ListGenericResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListGenericResourceVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGenericResourceVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListGenericResourceVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListGenericResourceVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListGenericResourceVersions(ctx, req.(*ListGenericResourceVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListVersionedResourceOccurrences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionedResourceOccurrencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListVersionedResourceOccurrences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListVersionedResourceOccurrences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListVersionedResourceOccurrences(ctx, req.(*ListVersionedResourceOccurrencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListOccurrences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOccurrencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListOccurrences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListOccurrences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListOccurrences(ctx, req.(*ListOccurrencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_UpdateOccurrence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOccurrenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).UpdateOccurrence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/UpdateOccurrence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).UpdateOccurrence(ctx, req.(*UpdateOccurrenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/CreatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).CreatePolicy(ctx, req.(*PolicyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_ValidatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).ValidatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/ValidatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).ValidatePolicy(ctx, req.(*ValidatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/UpdatePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_RegisterCollector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCollectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).RegisterCollector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/RegisterCollector",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).RegisterCollector(ctx, req.(*RegisterCollectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rode_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RodeServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rode.v1alpha1.Rode/CreateNote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RodeServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rode_ServiceDesc is the grpc.ServiceDesc for Rode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rode.v1alpha1.Rode",
	HandlerType: (*RodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchCreateOccurrences",
			Handler:    _Rode_BatchCreateOccurrences_Handler,
		},
		{
			MethodName: "EvaluatePolicy",
			Handler:    _Rode_EvaluatePolicy_Handler,
		},
		{
			MethodName: "ListResources",
			Handler:    _Rode_ListResources_Handler,
		},
		{
			MethodName: "ListGenericResources",
			Handler:    _Rode_ListGenericResources_Handler,
		},
		{
			MethodName: "ListGenericResourceVersions",
			Handler:    _Rode_ListGenericResourceVersions_Handler,
		},
		{
			MethodName: "ListVersionedResourceOccurrences",
			Handler:    _Rode_ListVersionedResourceOccurrences_Handler,
		},
		{
			MethodName: "ListOccurrences",
			Handler:    _Rode_ListOccurrences_Handler,
		},
		{
			MethodName: "UpdateOccurrence",
			Handler:    _Rode_UpdateOccurrence_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Rode_CreatePolicy_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _Rode_GetPolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _Rode_DeletePolicy_Handler,
		},
		{
			MethodName: "ListPolicies",
			Handler:    _Rode_ListPolicies_Handler,
		},
		{
			MethodName: "ValidatePolicy",
			Handler:    _Rode_ValidatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _Rode_UpdatePolicy_Handler,
		},
		{
			MethodName: "RegisterCollector",
			Handler:    _Rode_RegisterCollector_Handler,
		},
		{
			MethodName: "CreateNote",
			Handler:    _Rode_CreateNote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/rode.proto",
}
