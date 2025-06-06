// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/aiplatform/v1/migration_service.proto

package aiplatformpb

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

// MigrationServiceClient is the client API for MigrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MigrationServiceClient interface {
	// Searches all of the resources in automl.googleapis.com,
	// datalabeling.googleapis.com and ml.googleapis.com that can be migrated to
	// Vertex AI's given location.
	SearchMigratableResources(ctx context.Context, in *SearchMigratableResourcesRequest, opts ...grpc.CallOption) (*SearchMigratableResourcesResponse, error)
	// Batch migrates resources from ml.googleapis.com, automl.googleapis.com,
	// and datalabeling.googleapis.com to Vertex AI.
	BatchMigrateResources(ctx context.Context, in *BatchMigrateResourcesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type migrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMigrationServiceClient(cc grpc.ClientConnInterface) MigrationServiceClient {
	return &migrationServiceClient{cc}
}

func (c *migrationServiceClient) SearchMigratableResources(ctx context.Context, in *SearchMigratableResourcesRequest, opts ...grpc.CallOption) (*SearchMigratableResourcesResponse, error) {
	out := new(SearchMigratableResourcesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1.MigrationService/SearchMigratableResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationServiceClient) BatchMigrateResources(ctx context.Context, in *BatchMigrateResourcesRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1.MigrationService/BatchMigrateResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServiceServer is the server API for MigrationService service.
// All implementations must embed UnimplementedMigrationServiceServer
// for forward compatibility
type MigrationServiceServer interface {
	// Searches all of the resources in automl.googleapis.com,
	// datalabeling.googleapis.com and ml.googleapis.com that can be migrated to
	// Vertex AI's given location.
	SearchMigratableResources(context.Context, *SearchMigratableResourcesRequest) (*SearchMigratableResourcesResponse, error)
	// Batch migrates resources from ml.googleapis.com, automl.googleapis.com,
	// and datalabeling.googleapis.com to Vertex AI.
	BatchMigrateResources(context.Context, *BatchMigrateResourcesRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedMigrationServiceServer()
}

// UnimplementedMigrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMigrationServiceServer struct {
}

func (UnimplementedMigrationServiceServer) SearchMigratableResources(context.Context, *SearchMigratableResourcesRequest) (*SearchMigratableResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMigratableResources not implemented")
}
func (UnimplementedMigrationServiceServer) BatchMigrateResources(context.Context, *BatchMigrateResourcesRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchMigrateResources not implemented")
}
func (UnimplementedMigrationServiceServer) mustEmbedUnimplementedMigrationServiceServer() {}

// UnsafeMigrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MigrationServiceServer will
// result in compilation errors.
type UnsafeMigrationServiceServer interface {
	mustEmbedUnimplementedMigrationServiceServer()
}

func RegisterMigrationServiceServer(s grpc.ServiceRegistrar, srv MigrationServiceServer) {
	s.RegisterService(&MigrationService_ServiceDesc, srv)
}

func _MigrationService_SearchMigratableResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMigratableResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServiceServer).SearchMigratableResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1.MigrationService/SearchMigratableResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServiceServer).SearchMigratableResources(ctx, req.(*SearchMigratableResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MigrationService_BatchMigrateResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchMigrateResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServiceServer).BatchMigrateResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1.MigrationService/BatchMigrateResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServiceServer).BatchMigrateResources(ctx, req.(*BatchMigrateResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MigrationService_ServiceDesc is the grpc.ServiceDesc for MigrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MigrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.aiplatform.v1.MigrationService",
	HandlerType: (*MigrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMigratableResources",
			Handler:    _MigrationService_SearchMigratableResources_Handler,
		},
		{
			MethodName: "BatchMigrateResources",
			Handler:    _MigrationService_BatchMigrateResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/aiplatform/v1/migration_service.proto",
}
