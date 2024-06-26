// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: dd/configuration.proto

package dd

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectConfigService_GetProjectConfig_FullMethodName = "/ProjectConfigService/GetProjectConfig"
)

// ProjectConfigServiceClient is the client API for ProjectConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectConfigServiceClient interface {
	GetProjectConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ProjectConfigResponse, error)
}

type projectConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectConfigServiceClient(cc grpc.ClientConnInterface) ProjectConfigServiceClient {
	return &projectConfigServiceClient{cc}
}

func (c *projectConfigServiceClient) GetProjectConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ProjectConfigResponse, error) {
	out := new(ProjectConfigResponse)
	err := c.cc.Invoke(ctx, ProjectConfigService_GetProjectConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectConfigServiceServer is the server API for ProjectConfigService service.
// All implementations must embed UnimplementedProjectConfigServiceServer
// for forward compatibility
type ProjectConfigServiceServer interface {
	GetProjectConfig(context.Context, *ConfigRequest) (*ProjectConfigResponse, error)
	mustEmbedUnimplementedProjectConfigServiceServer()
}

// UnimplementedProjectConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectConfigServiceServer struct {
}

func (UnimplementedProjectConfigServiceServer) GetProjectConfig(context.Context, *ConfigRequest) (*ProjectConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectConfig not implemented")
}
func (UnimplementedProjectConfigServiceServer) mustEmbedUnimplementedProjectConfigServiceServer() {}

// UnsafeProjectConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectConfigServiceServer will
// result in compilation errors.
type UnsafeProjectConfigServiceServer interface {
	mustEmbedUnimplementedProjectConfigServiceServer()
}

func RegisterProjectConfigServiceServer(s grpc.ServiceRegistrar, srv ProjectConfigServiceServer) {
	s.RegisterService(&ProjectConfigService_ServiceDesc, srv)
}

func _ProjectConfigService_GetProjectConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectConfigServiceServer).GetProjectConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectConfigService_GetProjectConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectConfigServiceServer).GetProjectConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectConfigService_ServiceDesc is the grpc.ServiceDesc for ProjectConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProjectConfigService",
	HandlerType: (*ProjectConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProjectConfig",
			Handler:    _ProjectConfigService_GetProjectConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dd/configuration.proto",
}
