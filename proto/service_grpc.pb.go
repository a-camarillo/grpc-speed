// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: proto/service.proto

package proto

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
	SpeedFetcher_FetchSpeed_FullMethodName = "/SpeedFetcher/FetchSpeed"
)

// SpeedFetcherClient is the client API for SpeedFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpeedFetcherClient interface {
	FetchSpeed(ctx context.Context, opts ...grpc.CallOption) (SpeedFetcher_FetchSpeedClient, error)
}

type speedFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewSpeedFetcherClient(cc grpc.ClientConnInterface) SpeedFetcherClient {
	return &speedFetcherClient{cc}
}

func (c *speedFetcherClient) FetchSpeed(ctx context.Context, opts ...grpc.CallOption) (SpeedFetcher_FetchSpeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpeedFetcher_ServiceDesc.Streams[0], SpeedFetcher_FetchSpeed_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &speedFetcherFetchSpeedClient{stream}
	return x, nil
}

type SpeedFetcher_FetchSpeedClient interface {
	Send(*TravelInfo) error
	Recv() (*TravelResponse, error)
	grpc.ClientStream
}

type speedFetcherFetchSpeedClient struct {
	grpc.ClientStream
}

func (x *speedFetcherFetchSpeedClient) Send(m *TravelInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *speedFetcherFetchSpeedClient) Recv() (*TravelResponse, error) {
	m := new(TravelResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeedFetcherServer is the server API for SpeedFetcher service.
// All implementations must embed UnimplementedSpeedFetcherServer
// for forward compatibility
type SpeedFetcherServer interface {
	FetchSpeed(SpeedFetcher_FetchSpeedServer) error
	mustEmbedUnimplementedSpeedFetcherServer()
}

// UnimplementedSpeedFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedSpeedFetcherServer struct {
}

func (UnimplementedSpeedFetcherServer) FetchSpeed(SpeedFetcher_FetchSpeedServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchSpeed not implemented")
}
func (UnimplementedSpeedFetcherServer) mustEmbedUnimplementedSpeedFetcherServer() {}

// UnsafeSpeedFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpeedFetcherServer will
// result in compilation errors.
type UnsafeSpeedFetcherServer interface {
	mustEmbedUnimplementedSpeedFetcherServer()
}

func RegisterSpeedFetcherServer(s grpc.ServiceRegistrar, srv SpeedFetcherServer) {
	s.RegisterService(&SpeedFetcher_ServiceDesc, srv)
}

func _SpeedFetcher_FetchSpeed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SpeedFetcherServer).FetchSpeed(&speedFetcherFetchSpeedServer{stream})
}

type SpeedFetcher_FetchSpeedServer interface {
	Send(*TravelResponse) error
	Recv() (*TravelInfo, error)
	grpc.ServerStream
}

type speedFetcherFetchSpeedServer struct {
	grpc.ServerStream
}

func (x *speedFetcherFetchSpeedServer) Send(m *TravelResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *speedFetcherFetchSpeedServer) Recv() (*TravelInfo, error) {
	m := new(TravelInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpeedFetcher_ServiceDesc is the grpc.ServiceDesc for SpeedFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpeedFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SpeedFetcher",
	HandlerType: (*SpeedFetcherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchSpeed",
			Handler:       _SpeedFetcher_FetchSpeed_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}