// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: bc/analytic.proto

package bc

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
	AnalyticService_StreamRealtimeData_FullMethodName = "/analytic.AnalyticService/StreamRealtimeData"
)

// AnalyticServiceClient is the client API for AnalyticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalyticServiceClient interface {
	StreamRealtimeData(ctx context.Context, in *AnalyticRequest, opts ...grpc.CallOption) (AnalyticService_StreamRealtimeDataClient, error)
}

type analyticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalyticServiceClient(cc grpc.ClientConnInterface) AnalyticServiceClient {
	return &analyticServiceClient{cc}
}

func (c *analyticServiceClient) StreamRealtimeData(ctx context.Context, in *AnalyticRequest, opts ...grpc.CallOption) (AnalyticService_StreamRealtimeDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &AnalyticService_ServiceDesc.Streams[0], AnalyticService_StreamRealtimeData_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &analyticServiceStreamRealtimeDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AnalyticService_StreamRealtimeDataClient interface {
	Recv() (*AnalyticData, error)
	grpc.ClientStream
}

type analyticServiceStreamRealtimeDataClient struct {
	grpc.ClientStream
}

func (x *analyticServiceStreamRealtimeDataClient) Recv() (*AnalyticData, error) {
	m := new(AnalyticData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnalyticServiceServer is the server API for AnalyticService service.
// All implementations must embed UnimplementedAnalyticServiceServer
// for forward compatibility
type AnalyticServiceServer interface {
	StreamRealtimeData(*AnalyticRequest, AnalyticService_StreamRealtimeDataServer) error
	mustEmbedUnimplementedAnalyticServiceServer()
}

// UnimplementedAnalyticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalyticServiceServer struct {
}

func (UnimplementedAnalyticServiceServer) StreamRealtimeData(*AnalyticRequest, AnalyticService_StreamRealtimeDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRealtimeData not implemented")
}
func (UnimplementedAnalyticServiceServer) mustEmbedUnimplementedAnalyticServiceServer() {}

// UnsafeAnalyticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalyticServiceServer will
// result in compilation errors.
type UnsafeAnalyticServiceServer interface {
	mustEmbedUnimplementedAnalyticServiceServer()
}

func RegisterAnalyticServiceServer(s grpc.ServiceRegistrar, srv AnalyticServiceServer) {
	s.RegisterService(&AnalyticService_ServiceDesc, srv)
}

func _AnalyticService_StreamRealtimeData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AnalyticRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnalyticServiceServer).StreamRealtimeData(m, &analyticServiceStreamRealtimeDataServer{stream})
}

type AnalyticService_StreamRealtimeDataServer interface {
	Send(*AnalyticData) error
	grpc.ServerStream
}

type analyticServiceStreamRealtimeDataServer struct {
	grpc.ServerStream
}

func (x *analyticServiceStreamRealtimeDataServer) Send(m *AnalyticData) error {
	return x.ServerStream.SendMsg(m)
}

// AnalyticService_ServiceDesc is the grpc.ServiceDesc for AnalyticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnalyticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytic.AnalyticService",
	HandlerType: (*AnalyticServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRealtimeData",
			Handler:       _AnalyticService_StreamRealtimeData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bc/analytic.proto",
}
