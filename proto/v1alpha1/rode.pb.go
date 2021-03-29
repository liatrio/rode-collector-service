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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: proto/v1alpha1/rode.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grafeas_go_proto "github.com/rode/rode/protodeps/grafeas/proto/v1beta1/grafeas_go_proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request to create occurrences in batch.
type BatchCreateOccurrencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The occurrences to create.
	Occurrences []*grafeas_go_proto.Occurrence `protobuf:"bytes,1,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
}

func (x *BatchCreateOccurrencesRequest) Reset() {
	*x = BatchCreateOccurrencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateOccurrencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateOccurrencesRequest) ProtoMessage() {}

func (x *BatchCreateOccurrencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateOccurrencesRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateOccurrencesRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{0}
}

func (x *BatchCreateOccurrencesRequest) GetOccurrences() []*grafeas_go_proto.Occurrence {
	if x != nil {
		return x.Occurrences
	}
	return nil
}

// Response for creating occurrences in batch.
type BatchCreateOccurrencesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The occurrences that were created.
	Occurrences []*grafeas_go_proto.Occurrence `protobuf:"bytes,1,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
}

func (x *BatchCreateOccurrencesResponse) Reset() {
	*x = BatchCreateOccurrencesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateOccurrencesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateOccurrencesResponse) ProtoMessage() {}

func (x *BatchCreateOccurrencesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateOccurrencesResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateOccurrencesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{1}
}

func (x *BatchCreateOccurrencesResponse) GetOccurrences() []*grafeas_go_proto.Occurrence {
	if x != nil {
		return x.Occurrences
	}
	return nil
}

// modeled after Grafeas' ListOccurrence request/response
// https://github.com/grafeas/grafeas/blob/5b072a9930eace404066502b49a72e5b420d3576/proto/v1beta1/grafeas.proto#L345-L374
type ListResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListResourcesRequest) Reset() {
	*x = ListResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesRequest) ProtoMessage() {}

func (x *ListResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesRequest.ProtoReflect.Descriptor instead.
func (*ListResourcesRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{2}
}

func (x *ListResourcesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListResourcesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListResourcesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources     []*grafeas_go_proto.Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	NextPageToken string                       `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListResourcesResponse) Reset() {
	*x = ListResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcesResponse) ProtoMessage() {}

func (x *ListResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcesResponse.ProtoReflect.Descriptor instead.
func (*ListResourcesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{3}
}

func (x *ListResourcesResponse) GetResources() []*grafeas_go_proto.Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ListResourcesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListOccurrencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListOccurrencesRequest) Reset() {
	*x = ListOccurrencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOccurrencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOccurrencesRequest) ProtoMessage() {}

func (x *ListOccurrencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOccurrencesRequest.ProtoReflect.Descriptor instead.
func (*ListOccurrencesRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{4}
}

func (x *ListOccurrencesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListOccurrencesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListOccurrencesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response for listing occurrences.
type ListOccurrencesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The occurrences requested.
	Occurrences   []*grafeas_go_proto.Occurrence `protobuf:"bytes,1,rep,name=occurrences,proto3" json:"occurrences,omitempty"`
	NextPageToken string                         `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListOccurrencesResponse) Reset() {
	*x = ListOccurrencesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOccurrencesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOccurrencesResponse) ProtoMessage() {}

func (x *ListOccurrencesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOccurrencesResponse.ProtoReflect.Descriptor instead.
func (*ListOccurrencesResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{5}
}

func (x *ListOccurrencesResponse) GetOccurrences() []*grafeas_go_proto.Occurrence {
	if x != nil {
		return x.Occurrences
	}
	return nil
}

func (x *ListOccurrencesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type UpdateOccurrenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Occurrence *grafeas_go_proto.Occurrence `protobuf:"bytes,2,opt,name=occurrence,proto3" json:"occurrence,omitempty"`
	UpdateMask *field_mask.FieldMask        `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateOccurrenceRequest) Reset() {
	*x = UpdateOccurrenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1alpha1_rode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOccurrenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOccurrenceRequest) ProtoMessage() {}

func (x *UpdateOccurrenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1alpha1_rode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOccurrenceRequest.ProtoReflect.Descriptor instead.
func (*UpdateOccurrenceRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1alpha1_rode_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOccurrenceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOccurrenceRequest) GetOccurrence() *grafeas_go_proto.Occurrence {
	if x != nil {
		return x.Occurrence
	}
	return nil
}

func (x *UpdateOccurrenceRequest) GetUpdateMask() *field_mask.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_proto_v1alpha1_rode_proto protoreflect.FileDescriptor

var file_proto_v1alpha1_rode_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5e, 0x0a, 0x1d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x5f, 0x0a, 0x1e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x22, 0x6a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x78, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65,
	0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa3, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x32, 0x9f, 0x0b,
	0x0a, 0x04, 0x52, 0x6f, 0x64, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x2c, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x63, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x3a, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a,
	0x0e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x24, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x7d, 0x3a, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x72,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x72, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x32, 0x1a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x3a, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0xda,
	0x41, 0x1b, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x61, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b, 0x2e,
	0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1f, 0x2e,
	0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0xda, 0x41, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x83, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x72, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x72, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x32, 0x17, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xda, 0x41,
	0x0e, 0x69, 0x64, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x6f,
	0x64, 0x65, 0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1alpha1_rode_proto_rawDescOnce sync.Once
	file_proto_v1alpha1_rode_proto_rawDescData = file_proto_v1alpha1_rode_proto_rawDesc
)

func file_proto_v1alpha1_rode_proto_rawDescGZIP() []byte {
	file_proto_v1alpha1_rode_proto_rawDescOnce.Do(func() {
		file_proto_v1alpha1_rode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1alpha1_rode_proto_rawDescData)
	})
	return file_proto_v1alpha1_rode_proto_rawDescData
}

var file_proto_v1alpha1_rode_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_v1alpha1_rode_proto_goTypes = []interface{}{
	(*BatchCreateOccurrencesRequest)(nil),  // 0: rode.v1alpha1.BatchCreateOccurrencesRequest
	(*BatchCreateOccurrencesResponse)(nil), // 1: rode.v1alpha1.BatchCreateOccurrencesResponse
	(*ListResourcesRequest)(nil),           // 2: rode.v1alpha1.ListResourcesRequest
	(*ListResourcesResponse)(nil),          // 3: rode.v1alpha1.ListResourcesResponse
	(*ListOccurrencesRequest)(nil),         // 4: rode.v1alpha1.ListOccurrencesRequest
	(*ListOccurrencesResponse)(nil),        // 5: rode.v1alpha1.ListOccurrencesResponse
	(*UpdateOccurrenceRequest)(nil),        // 6: rode.v1alpha1.UpdateOccurrenceRequest
	(*grafeas_go_proto.Occurrence)(nil),    // 7: grafeas.v1beta1.Occurrence
	(*grafeas_go_proto.Resource)(nil),      // 8: grafeas.v1beta1.Resource
	(*field_mask.FieldMask)(nil),           // 9: google.protobuf.FieldMask
	(*EvaluatePolicyRequest)(nil),          // 10: rode.v1alpha1.EvaluatePolicyRequest
	(*PolicyEntity)(nil),                   // 11: rode.v1alpha1.PolicyEntity
	(*GetPolicyRequest)(nil),               // 12: rode.v1alpha1.GetPolicyRequest
	(*DeletePolicyRequest)(nil),            // 13: rode.v1alpha1.DeletePolicyRequest
	(*ListPoliciesRequest)(nil),            // 14: rode.v1alpha1.ListPoliciesRequest
	(*ValidatePolicyRequest)(nil),          // 15: rode.v1alpha1.ValidatePolicyRequest
	(*UpdatePolicyRequest)(nil),            // 16: rode.v1alpha1.UpdatePolicyRequest
	(*EvaluatePolicyResponse)(nil),         // 17: rode.v1alpha1.EvaluatePolicyResponse
	(*Policy)(nil),                         // 18: rode.v1alpha1.Policy
	(*empty.Empty)(nil),                    // 19: google.protobuf.Empty
	(*ListPoliciesResponse)(nil),           // 20: rode.v1alpha1.ListPoliciesResponse
	(*ValidatePolicyResponse)(nil),         // 21: rode.v1alpha1.ValidatePolicyResponse
}
var file_proto_v1alpha1_rode_proto_depIdxs = []int32{
	7,  // 0: rode.v1alpha1.BatchCreateOccurrencesRequest.occurrences:type_name -> grafeas.v1beta1.Occurrence
	7,  // 1: rode.v1alpha1.BatchCreateOccurrencesResponse.occurrences:type_name -> grafeas.v1beta1.Occurrence
	8,  // 2: rode.v1alpha1.ListResourcesResponse.resources:type_name -> grafeas.v1beta1.Resource
	7,  // 3: rode.v1alpha1.ListOccurrencesResponse.occurrences:type_name -> grafeas.v1beta1.Occurrence
	7,  // 4: rode.v1alpha1.UpdateOccurrenceRequest.occurrence:type_name -> grafeas.v1beta1.Occurrence
	9,  // 5: rode.v1alpha1.UpdateOccurrenceRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 6: rode.v1alpha1.Rode.BatchCreateOccurrences:input_type -> rode.v1alpha1.BatchCreateOccurrencesRequest
	10, // 7: rode.v1alpha1.Rode.EvaluatePolicy:input_type -> rode.v1alpha1.EvaluatePolicyRequest
	2,  // 8: rode.v1alpha1.Rode.ListResources:input_type -> rode.v1alpha1.ListResourcesRequest
	4,  // 9: rode.v1alpha1.Rode.ListOccurrences:input_type -> rode.v1alpha1.ListOccurrencesRequest
	6,  // 10: rode.v1alpha1.Rode.UpdateOccurrence:input_type -> rode.v1alpha1.UpdateOccurrenceRequest
	11, // 11: rode.v1alpha1.Rode.CreatePolicy:input_type -> rode.v1alpha1.PolicyEntity
	12, // 12: rode.v1alpha1.Rode.GetPolicy:input_type -> rode.v1alpha1.GetPolicyRequest
	13, // 13: rode.v1alpha1.Rode.DeletePolicy:input_type -> rode.v1alpha1.DeletePolicyRequest
	14, // 14: rode.v1alpha1.Rode.ListPolicies:input_type -> rode.v1alpha1.ListPoliciesRequest
	15, // 15: rode.v1alpha1.Rode.ValidatePolicy:input_type -> rode.v1alpha1.ValidatePolicyRequest
	16, // 16: rode.v1alpha1.Rode.UpdatePolicy:input_type -> rode.v1alpha1.UpdatePolicyRequest
	1,  // 17: rode.v1alpha1.Rode.BatchCreateOccurrences:output_type -> rode.v1alpha1.BatchCreateOccurrencesResponse
	17, // 18: rode.v1alpha1.Rode.EvaluatePolicy:output_type -> rode.v1alpha1.EvaluatePolicyResponse
	3,  // 19: rode.v1alpha1.Rode.ListResources:output_type -> rode.v1alpha1.ListResourcesResponse
	5,  // 20: rode.v1alpha1.Rode.ListOccurrences:output_type -> rode.v1alpha1.ListOccurrencesResponse
	7,  // 21: rode.v1alpha1.Rode.UpdateOccurrence:output_type -> grafeas.v1beta1.Occurrence
	18, // 22: rode.v1alpha1.Rode.CreatePolicy:output_type -> rode.v1alpha1.Policy
	18, // 23: rode.v1alpha1.Rode.GetPolicy:output_type -> rode.v1alpha1.Policy
	19, // 24: rode.v1alpha1.Rode.DeletePolicy:output_type -> google.protobuf.Empty
	20, // 25: rode.v1alpha1.Rode.ListPolicies:output_type -> rode.v1alpha1.ListPoliciesResponse
	21, // 26: rode.v1alpha1.Rode.ValidatePolicy:output_type -> rode.v1alpha1.ValidatePolicyResponse
	18, // 27: rode.v1alpha1.Rode.UpdatePolicy:output_type -> rode.v1alpha1.Policy
	17, // [17:28] is the sub-list for method output_type
	6,  // [6:17] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_v1alpha1_rode_proto_init() }
func file_proto_v1alpha1_rode_proto_init() {
	if File_proto_v1alpha1_rode_proto != nil {
		return
	}
	file_proto_v1alpha1_rode_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_v1alpha1_rode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateOccurrencesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateOccurrencesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOccurrencesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOccurrencesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_v1alpha1_rode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOccurrenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1alpha1_rode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1alpha1_rode_proto_goTypes,
		DependencyIndexes: file_proto_v1alpha1_rode_proto_depIdxs,
		MessageInfos:      file_proto_v1alpha1_rode_proto_msgTypes,
	}.Build()
	File_proto_v1alpha1_rode_proto = out.File
	file_proto_v1alpha1_rode_proto_rawDesc = nil
	file_proto_v1alpha1_rode_proto_goTypes = nil
	file_proto_v1alpha1_rode_proto_depIdxs = nil
}
