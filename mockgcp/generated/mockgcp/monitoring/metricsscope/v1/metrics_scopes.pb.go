// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mockgcp/monitoring/metricsscope/v1/metrics_scopes.proto

package metricsscopepb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Batch operation states.
type OperationMetadata_State int32

const (
	// Invalid.
	OperationMetadata_STATE_UNSPECIFIED OperationMetadata_State = 0
	// Request has been received.
	OperationMetadata_CREATED OperationMetadata_State = 1
	// Request is actively being processed.
	OperationMetadata_RUNNING OperationMetadata_State = 2
	// The batch processing is done.
	OperationMetadata_DONE OperationMetadata_State = 3
	// The batch processing was cancelled.
	OperationMetadata_CANCELLED OperationMetadata_State = 4
)

// Enum value maps for OperationMetadata_State.
var (
	OperationMetadata_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATED",
		2: "RUNNING",
		3: "DONE",
		4: "CANCELLED",
	}
	OperationMetadata_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATED":           1,
		"RUNNING":           2,
		"DONE":              3,
		"CANCELLED":         4,
	}
)

func (x OperationMetadata_State) Enum() *OperationMetadata_State {
	p := new(OperationMetadata_State)
	*p = x
	return p
}

func (x OperationMetadata_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationMetadata_State) Descriptor() protoreflect.EnumDescriptor {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_enumTypes[0].Descriptor()
}

func (OperationMetadata_State) Type() protoreflect.EnumType {
	return &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_enumTypes[0]
}

func (x OperationMetadata_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationMetadata_State.Descriptor instead.
func (OperationMetadata_State) EnumDescriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{5, 0}
}

// Request for the `GetMetricsScope` method.
type GetMetricsScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the `Metrics Scope`.
	// Example:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetMetricsScopeRequest) Reset() {
	*x = GetMetricsScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetricsScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsScopeRequest) ProtoMessage() {}

func (x *GetMetricsScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsScopeRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsScopeRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{0}
}

func (x *GetMetricsScopeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for the `ListMetricsScopesByMonitoredProject` method.
type ListMetricsScopesByMonitoredProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the `Monitored Project` being requested.
	// Example:
	// `projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	MonitoredResourceContainer string `protobuf:"bytes,1,opt,name=monitored_resource_container,json=monitoredResourceContainer,proto3" json:"monitored_resource_container,omitempty"`
}

func (x *ListMetricsScopesByMonitoredProjectRequest) Reset() {
	*x = ListMetricsScopesByMonitoredProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMetricsScopesByMonitoredProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricsScopesByMonitoredProjectRequest) ProtoMessage() {}

func (x *ListMetricsScopesByMonitoredProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricsScopesByMonitoredProjectRequest.ProtoReflect.Descriptor instead.
func (*ListMetricsScopesByMonitoredProjectRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{1}
}

func (x *ListMetricsScopesByMonitoredProjectRequest) GetMonitoredResourceContainer() string {
	if x != nil {
		return x.MonitoredResourceContainer
	}
	return ""
}

// Response for the `ListMetricsScopesByMonitoredProject` method.
type ListMetricsScopesByMonitoredProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A set of all metrics scopes that the specified monitored project has been
	// added to.
	MetricsScopes []*MetricsScope `protobuf:"bytes,1,rep,name=metrics_scopes,json=metricsScopes,proto3" json:"metrics_scopes,omitempty"`
}

func (x *ListMetricsScopesByMonitoredProjectResponse) Reset() {
	*x = ListMetricsScopesByMonitoredProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMetricsScopesByMonitoredProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMetricsScopesByMonitoredProjectResponse) ProtoMessage() {}

func (x *ListMetricsScopesByMonitoredProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMetricsScopesByMonitoredProjectResponse.ProtoReflect.Descriptor instead.
func (*ListMetricsScopesByMonitoredProjectResponse) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{2}
}

func (x *ListMetricsScopesByMonitoredProjectResponse) GetMetricsScopes() []*MetricsScope {
	if x != nil {
		return x.MetricsScopes
	}
	return nil
}

// Request for the `CreateMonitoredProject` method.
type CreateMonitoredProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the existing `Metrics Scope` that will monitor this
	// project.
	// Example:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The initial `MonitoredProject` configuration.
	// Specify only the `monitored_project.name` field. All other fields are
	// ignored. The `monitored_project.name` must be in the format:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	MonitoredProject *MonitoredProject `protobuf:"bytes,2,opt,name=monitored_project,json=monitoredProject,proto3" json:"monitored_project,omitempty"`
}

func (x *CreateMonitoredProjectRequest) Reset() {
	*x = CreateMonitoredProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMonitoredProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMonitoredProjectRequest) ProtoMessage() {}

func (x *CreateMonitoredProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMonitoredProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateMonitoredProjectRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMonitoredProjectRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateMonitoredProjectRequest) GetMonitoredProject() *MonitoredProject {
	if x != nil {
		return x.MonitoredProject
	}
	return nil
}

// Request for the `DeleteMonitoredProject` method.
type DeleteMonitoredProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the `MonitoredProject`.
	// Example:
	// `locations/global/metricsScopes/{SCOPING_PROJECT_ID_OR_NUMBER}/projects/{MONITORED_PROJECT_ID_OR_NUMBER}`
	//
	// Authorization requires the following [Google
	// IAM](https://cloud.google.com/iam) permissions on both the `Metrics Scope`
	// and on the `MonitoredProject`: `monitoring.metricsScopes.link`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteMonitoredProjectRequest) Reset() {
	*x = DeleteMonitoredProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMonitoredProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMonitoredProjectRequest) ProtoMessage() {}

func (x *DeleteMonitoredProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMonitoredProjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteMonitoredProjectRequest) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteMonitoredProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Contains metadata for longrunning operation for the edit Metrics Scope
// endpoints.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current state of the batch operation.
	State OperationMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=mockgcp.monitoring.metricsscope.v1.OperationMetadata_State" json:"state,omitempty"`
	// The time when the batch request was received.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The time when the operation result was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP(), []int{5}
}

func (x *OperationMetadata) GetState() OperationMetadata_State {
	if x != nil {
		return x.State
	}
	return OperationMetadata_STATE_UNSPECIFIED
}

func (x *OperationMetadata) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto protoreflect.FileDescriptor

var file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6d, 0x6f, 0x63, 0x6b, 0x67,
	0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x28, 0x0a, 0x26,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x2a,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x1c, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x1a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x22, 0x86, 0x01, 0x0a, 0x2b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x57, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6d, 0x6f, 0x63, 0x6b,
	0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0d, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x1d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x67, 0x0a, 0x1d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x32, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2c, 0x0a, 0x2a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb3, 0x02, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x6d, 0x6f, 0x63,
	0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xc8, 0x09, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0xbb, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x3a, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x96, 0x02, 0x0a, 0x23,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x4e, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x4f, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x12, 0x46, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x3a, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x42, 0x79, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x91, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x41, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x94, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x22, 0x36, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x3a, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0xda, 0x41, 0x18, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0xca, 0x41, 0x25, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0xee, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x41, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x2a, 0x36, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0x41, 0x2a, 0x0a,
	0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xda, 0x01, 0xca, 0x41, 0x19, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0xba, 0x01, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x72, 0x65, 0x61, 0x64, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77,
	0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x91, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x6d,
	0x6f, 0x63, 0x6b, 0x67, 0x63, 0x70, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x70, 0x62, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x70, 0x62, 0xaa, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x2b, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescOnce sync.Once
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescData = file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDesc
)

func file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescGZIP() []byte {
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescOnce.Do(func() {
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescData = protoimpl.X.CompressGZIP(file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescData)
	})
	return file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDescData
}

var file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_goTypes = []interface{}{
	(OperationMetadata_State)(0),                        // 0: mockgcp.monitoring.metricsscope.v1.OperationMetadata.State
	(*GetMetricsScopeRequest)(nil),                      // 1: mockgcp.monitoring.metricsscope.v1.GetMetricsScopeRequest
	(*ListMetricsScopesByMonitoredProjectRequest)(nil),  // 2: mockgcp.monitoring.metricsscope.v1.ListMetricsScopesByMonitoredProjectRequest
	(*ListMetricsScopesByMonitoredProjectResponse)(nil), // 3: mockgcp.monitoring.metricsscope.v1.ListMetricsScopesByMonitoredProjectResponse
	(*CreateMonitoredProjectRequest)(nil),               // 4: mockgcp.monitoring.metricsscope.v1.CreateMonitoredProjectRequest
	(*DeleteMonitoredProjectRequest)(nil),               // 5: mockgcp.monitoring.metricsscope.v1.DeleteMonitoredProjectRequest
	(*OperationMetadata)(nil),                           // 6: mockgcp.monitoring.metricsscope.v1.OperationMetadata
	(*MetricsScope)(nil),                                // 7: mockgcp.monitoring.metricsscope.v1.MetricsScope
	(*MonitoredProject)(nil),                            // 8: mockgcp.monitoring.metricsscope.v1.MonitoredProject
	(*timestamp.Timestamp)(nil),                         // 9: google.protobuf.Timestamp
	(*longrunningpb.Operation)(nil),                     // 10: google.longrunning.Operation
}
var file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_depIdxs = []int32{
	7,  // 0: mockgcp.monitoring.metricsscope.v1.ListMetricsScopesByMonitoredProjectResponse.metrics_scopes:type_name -> mockgcp.monitoring.metricsscope.v1.MetricsScope
	8,  // 1: mockgcp.monitoring.metricsscope.v1.CreateMonitoredProjectRequest.monitored_project:type_name -> mockgcp.monitoring.metricsscope.v1.MonitoredProject
	0,  // 2: mockgcp.monitoring.metricsscope.v1.OperationMetadata.state:type_name -> mockgcp.monitoring.metricsscope.v1.OperationMetadata.State
	9,  // 3: mockgcp.monitoring.metricsscope.v1.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	9,  // 4: mockgcp.monitoring.metricsscope.v1.OperationMetadata.update_time:type_name -> google.protobuf.Timestamp
	1,  // 5: mockgcp.monitoring.metricsscope.v1.MetricsScopes.GetMetricsScope:input_type -> mockgcp.monitoring.metricsscope.v1.GetMetricsScopeRequest
	2,  // 6: mockgcp.monitoring.metricsscope.v1.MetricsScopes.ListMetricsScopesByMonitoredProject:input_type -> mockgcp.monitoring.metricsscope.v1.ListMetricsScopesByMonitoredProjectRequest
	4,  // 7: mockgcp.monitoring.metricsscope.v1.MetricsScopes.CreateMonitoredProject:input_type -> mockgcp.monitoring.metricsscope.v1.CreateMonitoredProjectRequest
	5,  // 8: mockgcp.monitoring.metricsscope.v1.MetricsScopes.DeleteMonitoredProject:input_type -> mockgcp.monitoring.metricsscope.v1.DeleteMonitoredProjectRequest
	7,  // 9: mockgcp.monitoring.metricsscope.v1.MetricsScopes.GetMetricsScope:output_type -> mockgcp.monitoring.metricsscope.v1.MetricsScope
	3,  // 10: mockgcp.monitoring.metricsscope.v1.MetricsScopes.ListMetricsScopesByMonitoredProject:output_type -> mockgcp.monitoring.metricsscope.v1.ListMetricsScopesByMonitoredProjectResponse
	10, // 11: mockgcp.monitoring.metricsscope.v1.MetricsScopes.CreateMonitoredProject:output_type -> google.longrunning.Operation
	10, // 12: mockgcp.monitoring.metricsscope.v1.MetricsScopes.DeleteMonitoredProject:output_type -> google.longrunning.Operation
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_init() }
func file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_init() {
	if File_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto != nil {
		return
	}
	file_mockgcp_monitoring_metricsscope_v1_metrics_scope_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetricsScopeRequest); i {
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
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricsScopesByMonitoredProjectRequest); i {
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
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMetricsScopesByMonitoredProjectResponse); i {
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
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMonitoredProjectRequest); i {
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
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMonitoredProjectRequest); i {
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
		file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_goTypes,
		DependencyIndexes: file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_depIdxs,
		EnumInfos:         file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_enumTypes,
		MessageInfos:      file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_msgTypes,
	}.Build()
	File_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto = out.File
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_rawDesc = nil
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_goTypes = nil
	file_mockgcp_monitoring_metricsscope_v1_metrics_scopes_proto_depIdxs = nil
}
