// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: deployment.proto

package deployment_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Types of platforms.
type Deployment_Platform int32

const (
	// Unknown.
	Deployment_PLATFORM_UNSPECIFIED Deployment_Platform = 0
	// Google Container Engine.
	Deployment_GKE Deployment_Platform = 1
	// Google App Engine: Flexible Environment.
	Deployment_FLEX Deployment_Platform = 2
	// Custom user-defined platform.
	Deployment_CUSTOM Deployment_Platform = 3
)

// Enum value maps for Deployment_Platform.
var (
	Deployment_Platform_name = map[int32]string{
		0: "PLATFORM_UNSPECIFIED",
		1: "GKE",
		2: "FLEX",
		3: "CUSTOM",
	}
	Deployment_Platform_value = map[string]int32{
		"PLATFORM_UNSPECIFIED": 0,
		"GKE":                  1,
		"FLEX":                 2,
		"CUSTOM":               3,
	}
)

func (x Deployment_Platform) Enum() *Deployment_Platform {
	p := new(Deployment_Platform)
	*p = x
	return p
}

func (x Deployment_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Deployment_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_deployment_proto_enumTypes[0].Descriptor()
}

func (Deployment_Platform) Type() protoreflect.EnumType {
	return &file_deployment_proto_enumTypes[0]
}

func (x Deployment_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Deployment_Platform.Descriptor instead.
func (Deployment_Platform) EnumDescriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{2, 0}
}

// An artifact that can be deployed in some runtime.
type Deployable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource URI for the artifact being deployed.
	ResourceUri []string `protobuf:"bytes,1,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
}

func (x *Deployable) Reset() {
	*x = Deployable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployable) ProtoMessage() {}

func (x *Deployable) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployable.ProtoReflect.Descriptor instead.
func (*Deployable) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployable) GetResourceUri() []string {
	if x != nil {
		return x.ResourceUri
	}
	return nil
}

// Details of a deployment occurrence.
type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Deployment history for the resource.
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *Details) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

// The period during which some deployable was active in a runtime.
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity of the user that triggered this deployment.
	UserEmail string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// Required. Beginning of the lifetime of this deployment.
	DeployTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=deploy_time,json=deployTime,proto3" json:"deploy_time,omitempty"`
	// End of the lifetime of this deployment.
	UndeployTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=undeploy_time,json=undeployTime,proto3" json:"undeploy_time,omitempty"`
	// Configuration used to create this deployment.
	Config string `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Address of the runtime element hosting this deployment.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Output only. Resource URI for the artifact being deployed taken from
	// the deployable field with the same name.
	ResourceUri []string `protobuf:"bytes,6,rep,name=resource_uri,json=resourceUri,proto3" json:"resource_uri,omitempty"`
	// Platform hosting this deployment.
	Platform Deployment_Platform `protobuf:"varint,7,opt,name=platform,proto3,enum=grafeas.v1beta1.deployment.Deployment_Platform" json:"platform,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_deployment_proto_rawDescGZIP(), []int{2}
}

func (x *Deployment) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Deployment) GetDeployTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeployTime
	}
	return nil
}

func (x *Deployment) GetUndeployTime() *timestamp.Timestamp {
	if x != nil {
		return x.UndeployTime
	}
	return nil
}

func (x *Deployment) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Deployment) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Deployment) GetResourceUri() []string {
	if x != nil {
		return x.ResourceUri
	}
	return nil
}

func (x *Deployment) GetPlatform() Deployment_Platform {
	if x != nil {
		return x.Platform
	}
	return Deployment_PLATFORM_UNSPECIFIED
}

var File_deployment_proto protoreflect.FileDescriptor

var file_deployment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x69,
	0x22, 0x51, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x46, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x90, 0x03, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x75, 0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x75, 0x6e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x72, 0x69, 0x12, 0x4b, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x22, 0x43, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x4b, 0x45, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x4c, 0x45, 0x58, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x03, 0x42, 0x78, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61, 0x74, 0x72, 0x69, 0x6f, 0x2f, 0x72, 0x6f,
	0x64, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x64, 0x65, 0x70, 0x73,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deployment_proto_rawDescOnce sync.Once
	file_deployment_proto_rawDescData = file_deployment_proto_rawDesc
)

func file_deployment_proto_rawDescGZIP() []byte {
	file_deployment_proto_rawDescOnce.Do(func() {
		file_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_deployment_proto_rawDescData)
	})
	return file_deployment_proto_rawDescData
}

var file_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deployment_proto_goTypes = []interface{}{
	(Deployment_Platform)(0),    // 0: grafeas.v1beta1.deployment.Deployment.Platform
	(*Deployable)(nil),          // 1: grafeas.v1beta1.deployment.Deployable
	(*Details)(nil),             // 2: grafeas.v1beta1.deployment.Details
	(*Deployment)(nil),          // 3: grafeas.v1beta1.deployment.Deployment
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_deployment_proto_depIdxs = []int32{
	3, // 0: grafeas.v1beta1.deployment.Details.deployment:type_name -> grafeas.v1beta1.deployment.Deployment
	4, // 1: grafeas.v1beta1.deployment.Deployment.deploy_time:type_name -> google.protobuf.Timestamp
	4, // 2: grafeas.v1beta1.deployment.Deployment.undeploy_time:type_name -> google.protobuf.Timestamp
	0, // 3: grafeas.v1beta1.deployment.Deployment.platform:type_name -> grafeas.v1beta1.deployment.Deployment.Platform
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_deployment_proto_init() }
func file_deployment_proto_init() {
	if File_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployable); i {
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
		file_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
			RawDescriptor: file_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deployment_proto_goTypes,
		DependencyIndexes: file_deployment_proto_depIdxs,
		EnumInfos:         file_deployment_proto_enumTypes,
		MessageInfos:      file_deployment_proto_msgTypes,
	}.Build()
	File_deployment_proto = out.File
	file_deployment_proto_rawDesc = nil
	file_deployment_proto_goTypes = nil
	file_deployment_proto_depIdxs = nil
}
