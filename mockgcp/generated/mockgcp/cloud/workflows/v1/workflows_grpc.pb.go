// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/workflows/v1/workflows.proto

package workflowspb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WorkflowsClient is the client API for Workflows service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowsClient interface {
	// Lists workflows in a given project and location.
	// The default order is not specified.
	ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error)
	// Gets details of a single workflow.
	GetWorkflow(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error)
	// Creates a new workflow. If a workflow with the specified name already
	// exists in the specified project and location, the long running operation
	// returns a [ALREADY_EXISTS][google.rpc.Code.ALREADY_EXISTS] error.
	CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a workflow with the specified name.
	// This method also cancels and deletes all running executions of the
	// workflow.
	DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Updates an existing workflow.
	// Running this method has no impact on already running executions of the
	// workflow. A new revision of the workflow might be created as a result of a
	// successful update operation. In that case, the new revision is used
	// in new workflow executions.
	UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists revisions for a given workflow.
	ListWorkflowRevisions(ctx context.Context, in *ListWorkflowRevisionsRequest, opts ...grpc.CallOption) (*ListWorkflowRevisionsResponse, error)
}

type workflowsClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowsClient(cc grpc.ClientConnInterface) WorkflowsClient {
	return &workflowsClient{cc}
}

func (c *workflowsClient) ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error) {
	out := new(ListWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowsClient) GetWorkflow(ctx context.Context, in *GetWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/GetWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowsClient) CreateWorkflow(ctx context.Context, in *CreateWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/CreateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowsClient) DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/DeleteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowsClient) UpdateWorkflow(ctx context.Context, in *UpdateWorkflowRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/UpdateWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowsClient) ListWorkflowRevisions(ctx context.Context, in *ListWorkflowRevisionsRequest, opts ...grpc.CallOption) (*ListWorkflowRevisionsResponse, error) {
	out := new(ListWorkflowRevisionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.workflows.v1.Workflows/ListWorkflowRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowsServer is the server API for Workflows service.
// All implementations must embed UnimplementedWorkflowsServer
// for forward compatibility
type WorkflowsServer interface {
	// Lists workflows in a given project and location.
	// The default order is not specified.
	ListWorkflows(context.Context, *ListWorkflowsRequest) (*ListWorkflowsResponse, error)
	// Gets details of a single workflow.
	GetWorkflow(context.Context, *GetWorkflowRequest) (*Workflow, error)
	// Creates a new workflow. If a workflow with the specified name already
	// exists in the specified project and location, the long running operation
	// returns a [ALREADY_EXISTS][google.rpc.Code.ALREADY_EXISTS] error.
	CreateWorkflow(context.Context, *CreateWorkflowRequest) (*longrunningpb.Operation, error)
	// Deletes a workflow with the specified name.
	// This method also cancels and deletes all running executions of the
	// workflow.
	DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*longrunningpb.Operation, error)
	// Updates an existing workflow.
	// Running this method has no impact on already running executions of the
	// workflow. A new revision of the workflow might be created as a result of a
	// successful update operation. In that case, the new revision is used
	// in new workflow executions.
	UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*longrunningpb.Operation, error)
	// Lists revisions for a given workflow.
	ListWorkflowRevisions(context.Context, *ListWorkflowRevisionsRequest) (*ListWorkflowRevisionsResponse, error)
	mustEmbedUnimplementedWorkflowsServer()
}

// UnimplementedWorkflowsServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowsServer struct {
}

func (UnimplementedWorkflowsServer) ListWorkflows(context.Context, *ListWorkflowsRequest) (*ListWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflows not implemented")
}
func (UnimplementedWorkflowsServer) GetWorkflow(context.Context, *GetWorkflowRequest) (*Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflow not implemented")
}
func (UnimplementedWorkflowsServer) CreateWorkflow(context.Context, *CreateWorkflowRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkflow not implemented")
}
func (UnimplementedWorkflowsServer) DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedWorkflowsServer) UpdateWorkflow(context.Context, *UpdateWorkflowRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkflow not implemented")
}
func (UnimplementedWorkflowsServer) ListWorkflowRevisions(context.Context, *ListWorkflowRevisionsRequest) (*ListWorkflowRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowRevisions not implemented")
}
func (UnimplementedWorkflowsServer) mustEmbedUnimplementedWorkflowsServer() {}

// UnsafeWorkflowsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowsServer will
// result in compilation errors.
type UnsafeWorkflowsServer interface {
	mustEmbedUnimplementedWorkflowsServer()
}

func RegisterWorkflowsServer(s grpc.ServiceRegistrar, srv WorkflowsServer) {
	s.RegisterService(&Workflows_ServiceDesc, srv)
}

func _Workflows_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).ListWorkflows(ctx, req.(*ListWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflows_GetWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).GetWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/GetWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).GetWorkflow(ctx, req.(*GetWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflows_CreateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).CreateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/CreateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).CreateWorkflow(ctx, req.(*CreateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflows_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/DeleteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).DeleteWorkflow(ctx, req.(*DeleteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflows_UpdateWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).UpdateWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/UpdateWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).UpdateWorkflow(ctx, req.(*UpdateWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Workflows_ListWorkflowRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowsServer).ListWorkflowRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.workflows.v1.Workflows/ListWorkflowRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowsServer).ListWorkflowRevisions(ctx, req.(*ListWorkflowRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Workflows_ServiceDesc is the grpc.ServiceDesc for Workflows service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Workflows_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.workflows.v1.Workflows",
	HandlerType: (*WorkflowsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkflows",
			Handler:    _Workflows_ListWorkflows_Handler,
		},
		{
			MethodName: "GetWorkflow",
			Handler:    _Workflows_GetWorkflow_Handler,
		},
		{
			MethodName: "CreateWorkflow",
			Handler:    _Workflows_CreateWorkflow_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _Workflows_DeleteWorkflow_Handler,
		},
		{
			MethodName: "UpdateWorkflow",
			Handler:    _Workflows_UpdateWorkflow_Handler,
		},
		{
			MethodName: "ListWorkflowRevisions",
			Handler:    _Workflows_ListWorkflowRevisions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/workflows/v1/workflows.proto",
}
