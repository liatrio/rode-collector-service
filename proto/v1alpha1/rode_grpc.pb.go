// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	grafeas_go_proto "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RodeClient is the client API for Rode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RodeClient interface {
	BatchCreateOccurrences(ctx context.Context, in *grafeas_go_proto.BatchCreateOccurrencesRequest, opts ...grpc.CallOption) (*grafeas_go_proto.BatchCreateOccurrencesResponse, error)
}

type rodeClient struct {
	cc grpc.ClientConnInterface
}

func NewRodeClient(cc grpc.ClientConnInterface) RodeClient {
	return &rodeClient{cc}
}

func (c *rodeClient) BatchCreateOccurrences(ctx context.Context, in *grafeas_go_proto.BatchCreateOccurrencesRequest, opts ...grpc.CallOption) (*grafeas_go_proto.BatchCreateOccurrencesResponse, error) {
	out := new(grafeas_go_proto.BatchCreateOccurrencesResponse)
	err := c.cc.Invoke(ctx, "/rode.v1alpha1.Rode/BatchCreateOccurrences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RodeServer is the server API for Rode service.
// All implementations must embed UnimplementedRodeServer
// for forward compatibility
type RodeServer interface {
	BatchCreateOccurrences(context.Context, *grafeas_go_proto.BatchCreateOccurrencesRequest) (*grafeas_go_proto.BatchCreateOccurrencesResponse, error)
	mustEmbedUnimplementedRodeServer()
}

// UnimplementedRodeServer must be embedded to have forward compatible implementations.
type UnimplementedRodeServer struct {
}

func (UnimplementedRodeServer) BatchCreateOccurrences(context.Context, *grafeas_go_proto.BatchCreateOccurrencesRequest) (*grafeas_go_proto.BatchCreateOccurrencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateOccurrences not implemented")
}
func (UnimplementedRodeServer) mustEmbedUnimplementedRodeServer() {}

// UnsafeRodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RodeServer will
// result in compilation errors.
type UnsafeRodeServer interface {
	mustEmbedUnimplementedRodeServer()
}

func RegisterRodeServer(s grpc.ServiceRegistrar, srv RodeServer) {
	s.RegisterService(&_Rode_serviceDesc, srv)
}

func _Rode_BatchCreateOccurrences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grafeas_go_proto.BatchCreateOccurrencesRequest)
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
		return srv.(RodeServer).BatchCreateOccurrences(ctx, req.(*grafeas_go_proto.BatchCreateOccurrencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rode.v1alpha1.Rode",
	HandlerType: (*RodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchCreateOccurrences",
			Handler:    _Rode_BatchCreateOccurrences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1alpha1/rode.proto",
}
