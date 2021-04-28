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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: proto/v1beta1/cvss.proto

package cvss_go_proto

import (
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

type CVSSv3_AttackVector int32

const (
	CVSSv3_ATTACK_VECTOR_UNSPECIFIED CVSSv3_AttackVector = 0
	CVSSv3_ATTACK_VECTOR_NETWORK     CVSSv3_AttackVector = 1
	CVSSv3_ATTACK_VECTOR_ADJACENT    CVSSv3_AttackVector = 2
	CVSSv3_ATTACK_VECTOR_LOCAL       CVSSv3_AttackVector = 3
	CVSSv3_ATTACK_VECTOR_PHYSICAL    CVSSv3_AttackVector = 4
)

// Enum value maps for CVSSv3_AttackVector.
var (
	CVSSv3_AttackVector_name = map[int32]string{
		0: "ATTACK_VECTOR_UNSPECIFIED",
		1: "ATTACK_VECTOR_NETWORK",
		2: "ATTACK_VECTOR_ADJACENT",
		3: "ATTACK_VECTOR_LOCAL",
		4: "ATTACK_VECTOR_PHYSICAL",
	}
	CVSSv3_AttackVector_value = map[string]int32{
		"ATTACK_VECTOR_UNSPECIFIED": 0,
		"ATTACK_VECTOR_NETWORK":     1,
		"ATTACK_VECTOR_ADJACENT":    2,
		"ATTACK_VECTOR_LOCAL":       3,
		"ATTACK_VECTOR_PHYSICAL":    4,
	}
)

func (x CVSSv3_AttackVector) Enum() *CVSSv3_AttackVector {
	p := new(CVSSv3_AttackVector)
	*p = x
	return p
}

func (x CVSSv3_AttackVector) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_AttackVector) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[0].Descriptor()
}

func (CVSSv3_AttackVector) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[0]
}

func (x CVSSv3_AttackVector) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_AttackVector.Descriptor instead.
func (CVSSv3_AttackVector) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 0}
}

type CVSSv3_AttackComplexity int32

const (
	CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED CVSSv3_AttackComplexity = 0
	CVSSv3_ATTACK_COMPLEXITY_LOW         CVSSv3_AttackComplexity = 1
	CVSSv3_ATTACK_COMPLEXITY_HIGH        CVSSv3_AttackComplexity = 2
)

// Enum value maps for CVSSv3_AttackComplexity.
var (
	CVSSv3_AttackComplexity_name = map[int32]string{
		0: "ATTACK_COMPLEXITY_UNSPECIFIED",
		1: "ATTACK_COMPLEXITY_LOW",
		2: "ATTACK_COMPLEXITY_HIGH",
	}
	CVSSv3_AttackComplexity_value = map[string]int32{
		"ATTACK_COMPLEXITY_UNSPECIFIED": 0,
		"ATTACK_COMPLEXITY_LOW":         1,
		"ATTACK_COMPLEXITY_HIGH":        2,
	}
)

func (x CVSSv3_AttackComplexity) Enum() *CVSSv3_AttackComplexity {
	p := new(CVSSv3_AttackComplexity)
	*p = x
	return p
}

func (x CVSSv3_AttackComplexity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_AttackComplexity) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[1].Descriptor()
}

func (CVSSv3_AttackComplexity) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[1]
}

func (x CVSSv3_AttackComplexity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_AttackComplexity.Descriptor instead.
func (CVSSv3_AttackComplexity) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 1}
}

type CVSSv3_PrivilegesRequired int32

const (
	CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED CVSSv3_PrivilegesRequired = 0
	CVSSv3_PRIVILEGES_REQUIRED_NONE        CVSSv3_PrivilegesRequired = 1
	CVSSv3_PRIVILEGES_REQUIRED_LOW         CVSSv3_PrivilegesRequired = 2
	CVSSv3_PRIVILEGES_REQUIRED_HIGH        CVSSv3_PrivilegesRequired = 3
)

// Enum value maps for CVSSv3_PrivilegesRequired.
var (
	CVSSv3_PrivilegesRequired_name = map[int32]string{
		0: "PRIVILEGES_REQUIRED_UNSPECIFIED",
		1: "PRIVILEGES_REQUIRED_NONE",
		2: "PRIVILEGES_REQUIRED_LOW",
		3: "PRIVILEGES_REQUIRED_HIGH",
	}
	CVSSv3_PrivilegesRequired_value = map[string]int32{
		"PRIVILEGES_REQUIRED_UNSPECIFIED": 0,
		"PRIVILEGES_REQUIRED_NONE":        1,
		"PRIVILEGES_REQUIRED_LOW":         2,
		"PRIVILEGES_REQUIRED_HIGH":        3,
	}
)

func (x CVSSv3_PrivilegesRequired) Enum() *CVSSv3_PrivilegesRequired {
	p := new(CVSSv3_PrivilegesRequired)
	*p = x
	return p
}

func (x CVSSv3_PrivilegesRequired) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_PrivilegesRequired) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[2].Descriptor()
}

func (CVSSv3_PrivilegesRequired) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[2]
}

func (x CVSSv3_PrivilegesRequired) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_PrivilegesRequired.Descriptor instead.
func (CVSSv3_PrivilegesRequired) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 2}
}

type CVSSv3_UserInteraction int32

const (
	CVSSv3_USER_INTERACTION_UNSPECIFIED CVSSv3_UserInteraction = 0
	CVSSv3_USER_INTERACTION_NONE        CVSSv3_UserInteraction = 1
	CVSSv3_USER_INTERACTION_REQUIRED    CVSSv3_UserInteraction = 2
)

// Enum value maps for CVSSv3_UserInteraction.
var (
	CVSSv3_UserInteraction_name = map[int32]string{
		0: "USER_INTERACTION_UNSPECIFIED",
		1: "USER_INTERACTION_NONE",
		2: "USER_INTERACTION_REQUIRED",
	}
	CVSSv3_UserInteraction_value = map[string]int32{
		"USER_INTERACTION_UNSPECIFIED": 0,
		"USER_INTERACTION_NONE":        1,
		"USER_INTERACTION_REQUIRED":    2,
	}
)

func (x CVSSv3_UserInteraction) Enum() *CVSSv3_UserInteraction {
	p := new(CVSSv3_UserInteraction)
	*p = x
	return p
}

func (x CVSSv3_UserInteraction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_UserInteraction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[3].Descriptor()
}

func (CVSSv3_UserInteraction) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[3]
}

func (x CVSSv3_UserInteraction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_UserInteraction.Descriptor instead.
func (CVSSv3_UserInteraction) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 3}
}

type CVSSv3_Scope int32

const (
	CVSSv3_SCOPE_UNSPECIFIED CVSSv3_Scope = 0
	CVSSv3_SCOPE_UNCHANGED   CVSSv3_Scope = 1
	CVSSv3_SCOPE_CHANGED     CVSSv3_Scope = 2
)

// Enum value maps for CVSSv3_Scope.
var (
	CVSSv3_Scope_name = map[int32]string{
		0: "SCOPE_UNSPECIFIED",
		1: "SCOPE_UNCHANGED",
		2: "SCOPE_CHANGED",
	}
	CVSSv3_Scope_value = map[string]int32{
		"SCOPE_UNSPECIFIED": 0,
		"SCOPE_UNCHANGED":   1,
		"SCOPE_CHANGED":     2,
	}
)

func (x CVSSv3_Scope) Enum() *CVSSv3_Scope {
	p := new(CVSSv3_Scope)
	*p = x
	return p
}

func (x CVSSv3_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[4].Descriptor()
}

func (CVSSv3_Scope) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[4]
}

func (x CVSSv3_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_Scope.Descriptor instead.
func (CVSSv3_Scope) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 4}
}

type CVSSv3_Impact int32

const (
	CVSSv3_IMPACT_UNSPECIFIED CVSSv3_Impact = 0
	CVSSv3_IMPACT_HIGH        CVSSv3_Impact = 1
	CVSSv3_IMPACT_LOW         CVSSv3_Impact = 2
	CVSSv3_IMPACT_NONE        CVSSv3_Impact = 3
)

// Enum value maps for CVSSv3_Impact.
var (
	CVSSv3_Impact_name = map[int32]string{
		0: "IMPACT_UNSPECIFIED",
		1: "IMPACT_HIGH",
		2: "IMPACT_LOW",
		3: "IMPACT_NONE",
	}
	CVSSv3_Impact_value = map[string]int32{
		"IMPACT_UNSPECIFIED": 0,
		"IMPACT_HIGH":        1,
		"IMPACT_LOW":         2,
		"IMPACT_NONE":        3,
	}
)

func (x CVSSv3_Impact) Enum() *CVSSv3_Impact {
	p := new(CVSSv3_Impact)
	*p = x
	return p
}

func (x CVSSv3_Impact) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CVSSv3_Impact) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_v1beta1_cvss_proto_enumTypes[5].Descriptor()
}

func (CVSSv3_Impact) Type() protoreflect.EnumType {
	return &file_proto_v1beta1_cvss_proto_enumTypes[5]
}

func (x CVSSv3_Impact) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CVSSv3_Impact.Descriptor instead.
func (CVSSv3_Impact) EnumDescriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0, 5}
}

// Common Vulnerability Scoring System version 3.
// For details, see https://www.first.org/cvss/specification-document
type CVSSv3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base score is a function of the base metric scores.
	BaseScore           float32 `protobuf:"fixed32,1,opt,name=base_score,json=baseScore,proto3" json:"base_score,omitempty"`
	ExploitabilityScore float32 `protobuf:"fixed32,2,opt,name=exploitability_score,json=exploitabilityScore,proto3" json:"exploitability_score,omitempty"`
	ImpactScore         float32 `protobuf:"fixed32,3,opt,name=impact_score,json=impactScore,proto3" json:"impact_score,omitempty"`
	// Base Metrics
	// Represents the intrinsic characteristics of a vulnerability that are
	// constant over time and across user environments.
	AttackVector          CVSSv3_AttackVector       `protobuf:"varint,5,opt,name=attack_vector,json=attackVector,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_AttackVector" json:"attack_vector,omitempty"`
	AttackComplexity      CVSSv3_AttackComplexity   `protobuf:"varint,6,opt,name=attack_complexity,json=attackComplexity,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_AttackComplexity" json:"attack_complexity,omitempty"`
	PrivilegesRequired    CVSSv3_PrivilegesRequired `protobuf:"varint,7,opt,name=privileges_required,json=privilegesRequired,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_PrivilegesRequired" json:"privileges_required,omitempty"`
	UserInteraction       CVSSv3_UserInteraction    `protobuf:"varint,8,opt,name=user_interaction,json=userInteraction,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_UserInteraction" json:"user_interaction,omitempty"`
	Scope                 CVSSv3_Scope              `protobuf:"varint,9,opt,name=scope,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_Scope" json:"scope,omitempty"`
	ConfidentialityImpact CVSSv3_Impact             `protobuf:"varint,10,opt,name=confidentiality_impact,json=confidentialityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_Impact" json:"confidentiality_impact,omitempty"`
	IntegrityImpact       CVSSv3_Impact             `protobuf:"varint,11,opt,name=integrity_impact,json=integrityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_Impact" json:"integrity_impact,omitempty"`
	AvailabilityImpact    CVSSv3_Impact             `protobuf:"varint,12,opt,name=availability_impact,json=availabilityImpact,proto3,enum=grafeas.v1beta1.vulnerability.CVSSv3_Impact" json:"availability_impact,omitempty"`
}

func (x *CVSSv3) Reset() {
	*x = CVSSv3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1beta1_cvss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVSSv3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVSSv3) ProtoMessage() {}

func (x *CVSSv3) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1beta1_cvss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVSSv3.ProtoReflect.Descriptor instead.
func (*CVSSv3) Descriptor() ([]byte, []int) {
	return file_proto_v1beta1_cvss_proto_rawDescGZIP(), []int{0}
}

func (x *CVSSv3) GetBaseScore() float32 {
	if x != nil {
		return x.BaseScore
	}
	return 0
}

func (x *CVSSv3) GetExploitabilityScore() float32 {
	if x != nil {
		return x.ExploitabilityScore
	}
	return 0
}

func (x *CVSSv3) GetImpactScore() float32 {
	if x != nil {
		return x.ImpactScore
	}
	return 0
}

func (x *CVSSv3) GetAttackVector() CVSSv3_AttackVector {
	if x != nil {
		return x.AttackVector
	}
	return CVSSv3_ATTACK_VECTOR_UNSPECIFIED
}

func (x *CVSSv3) GetAttackComplexity() CVSSv3_AttackComplexity {
	if x != nil {
		return x.AttackComplexity
	}
	return CVSSv3_ATTACK_COMPLEXITY_UNSPECIFIED
}

func (x *CVSSv3) GetPrivilegesRequired() CVSSv3_PrivilegesRequired {
	if x != nil {
		return x.PrivilegesRequired
	}
	return CVSSv3_PRIVILEGES_REQUIRED_UNSPECIFIED
}

func (x *CVSSv3) GetUserInteraction() CVSSv3_UserInteraction {
	if x != nil {
		return x.UserInteraction
	}
	return CVSSv3_USER_INTERACTION_UNSPECIFIED
}

func (x *CVSSv3) GetScope() CVSSv3_Scope {
	if x != nil {
		return x.Scope
	}
	return CVSSv3_SCOPE_UNSPECIFIED
}

func (x *CVSSv3) GetConfidentialityImpact() CVSSv3_Impact {
	if x != nil {
		return x.ConfidentialityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (x *CVSSv3) GetIntegrityImpact() CVSSv3_Impact {
	if x != nil {
		return x.IntegrityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

func (x *CVSSv3) GetAvailabilityImpact() CVSSv3_Impact {
	if x != nil {
		return x.AvailabilityImpact
	}
	return CVSSv3_IMPACT_UNSPECIFIED
}

var File_proto_v1beta1_cvss_proto protoreflect.FileDescriptor

var file_proto_v1beta1_cvss_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x63, 0x76, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x92, 0x0c, 0x0a, 0x06, 0x43, 0x56,
	0x53, 0x53, 0x76, 0x33, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x69, 0x74, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x13, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x69, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x69, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x63, 0x0a, 0x11, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e,
	0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56,
	0x53, 0x53, 0x76, 0x33, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x69, 0x74, 0x79, 0x52, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x69, 0x74, 0x79, 0x12, 0x69, 0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x12,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x60, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53,
	0x53, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x63, 0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x49,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x57, 0x0a, 0x10,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x49, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x49,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x5d, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x56, 0x53, 0x53, 0x76, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f,
	0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56,
	0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52,
	0x5f, 0x41, 0x44, 0x4a, 0x41, 0x43, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x41,
	0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x56,
	0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x50, 0x48, 0x59, 0x53, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04,
	0x22, 0x6c, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x78, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x54, 0x54, 0x41, 0x43,
	0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57,
	0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x58, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x02, 0x22, 0x92,
	0x01, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x49, 0x56, 0x49, 0x4c, 0x45,
	0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52,
	0x49, 0x56, 0x49, 0x4c, 0x45, 0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45,
	0x44, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49, 0x56,
	0x49, 0x4c, 0x45, 0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f,
	0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x49, 0x56, 0x49, 0x4c, 0x45,
	0x47, 0x45, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x48, 0x49, 0x47,
	0x48, 0x10, 0x03, 0x22, 0x6d, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x22, 0x46, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43, 0x4f, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x02, 0x22, 0x52, 0x0a, 0x06, 0x49, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4d, 0x50, 0x41, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42, 0x6e,
	0x0a, 0x20, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x64, 0x65, 0x70, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x76, 0x73, 0x73, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1beta1_cvss_proto_rawDescOnce sync.Once
	file_proto_v1beta1_cvss_proto_rawDescData = file_proto_v1beta1_cvss_proto_rawDesc
)

func file_proto_v1beta1_cvss_proto_rawDescGZIP() []byte {
	file_proto_v1beta1_cvss_proto_rawDescOnce.Do(func() {
		file_proto_v1beta1_cvss_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1beta1_cvss_proto_rawDescData)
	})
	return file_proto_v1beta1_cvss_proto_rawDescData
}

var file_proto_v1beta1_cvss_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_proto_v1beta1_cvss_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_v1beta1_cvss_proto_goTypes = []interface{}{
	(CVSSv3_AttackVector)(0),       // 0: grafeas.v1beta1.vulnerability.CVSSv3.AttackVector
	(CVSSv3_AttackComplexity)(0),   // 1: grafeas.v1beta1.vulnerability.CVSSv3.AttackComplexity
	(CVSSv3_PrivilegesRequired)(0), // 2: grafeas.v1beta1.vulnerability.CVSSv3.PrivilegesRequired
	(CVSSv3_UserInteraction)(0),    // 3: grafeas.v1beta1.vulnerability.CVSSv3.UserInteraction
	(CVSSv3_Scope)(0),              // 4: grafeas.v1beta1.vulnerability.CVSSv3.Scope
	(CVSSv3_Impact)(0),             // 5: grafeas.v1beta1.vulnerability.CVSSv3.Impact
	(*CVSSv3)(nil),                 // 6: grafeas.v1beta1.vulnerability.CVSSv3
}
var file_proto_v1beta1_cvss_proto_depIdxs = []int32{
	0, // 0: grafeas.v1beta1.vulnerability.CVSSv3.attack_vector:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.AttackVector
	1, // 1: grafeas.v1beta1.vulnerability.CVSSv3.attack_complexity:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.AttackComplexity
	2, // 2: grafeas.v1beta1.vulnerability.CVSSv3.privileges_required:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.PrivilegesRequired
	3, // 3: grafeas.v1beta1.vulnerability.CVSSv3.user_interaction:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.UserInteraction
	4, // 4: grafeas.v1beta1.vulnerability.CVSSv3.scope:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.Scope
	5, // 5: grafeas.v1beta1.vulnerability.CVSSv3.confidentiality_impact:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.Impact
	5, // 6: grafeas.v1beta1.vulnerability.CVSSv3.integrity_impact:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.Impact
	5, // 7: grafeas.v1beta1.vulnerability.CVSSv3.availability_impact:type_name -> grafeas.v1beta1.vulnerability.CVSSv3.Impact
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_v1beta1_cvss_proto_init() }
func file_proto_v1beta1_cvss_proto_init() {
	if File_proto_v1beta1_cvss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1beta1_cvss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVSSv3); i {
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
			RawDescriptor: file_proto_v1beta1_cvss_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1beta1_cvss_proto_goTypes,
		DependencyIndexes: file_proto_v1beta1_cvss_proto_depIdxs,
		EnumInfos:         file_proto_v1beta1_cvss_proto_enumTypes,
		MessageInfos:      file_proto_v1beta1_cvss_proto_msgTypes,
	}.Build()
	File_proto_v1beta1_cvss_proto = out.File
	file_proto_v1beta1_cvss_proto_rawDesc = nil
	file_proto_v1beta1_cvss_proto_goTypes = nil
	file_proto_v1beta1_cvss_proto_depIdxs = nil
}
