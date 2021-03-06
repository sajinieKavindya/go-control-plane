// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package config

import (
	context "context"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConfigDiscoveryServiceClient is the client API for ConfigDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigDiscoveryServiceClient interface {
	StreamConfigs(ctx context.Context, opts ...grpc.CallOption) (ConfigDiscoveryService_StreamConfigsClient, error)
	FetchConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type configDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigDiscoveryServiceClient(cc grpc.ClientConnInterface) ConfigDiscoveryServiceClient {
	return &configDiscoveryServiceClient{cc}
}

func (c *configDiscoveryServiceClient) StreamConfigs(ctx context.Context, opts ...grpc.CallOption) (ConfigDiscoveryService_StreamConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ConfigDiscoveryService_serviceDesc.Streams[0], "/discovery.service.config.ConfigDiscoveryService/StreamConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDiscoveryServiceStreamConfigsClient{stream}
	return x, nil
}

type ConfigDiscoveryService_StreamConfigsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type configDiscoveryServiceStreamConfigsClient struct {
	grpc.ClientStream
}

func (x *configDiscoveryServiceStreamConfigsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *configDiscoveryServiceStreamConfigsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *configDiscoveryServiceClient) FetchConfigs(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/discovery.service.config.ConfigDiscoveryService/FetchConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigDiscoveryServiceServer is the server API for ConfigDiscoveryService service.
// All implementations must embed UnimplementedConfigDiscoveryServiceServer
// for forward compatibility
type ConfigDiscoveryServiceServer interface {
	StreamConfigs(ConfigDiscoveryService_StreamConfigsServer) error
	FetchConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
	mustEmbedUnimplementedConfigDiscoveryServiceServer()
}

// UnimplementedConfigDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigDiscoveryServiceServer struct {
}

func (UnimplementedConfigDiscoveryServiceServer) StreamConfigs(ConfigDiscoveryService_StreamConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamConfigs not implemented")
}
func (UnimplementedConfigDiscoveryServiceServer) FetchConfigs(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchConfigs not implemented")
}
func (UnimplementedConfigDiscoveryServiceServer) mustEmbedUnimplementedConfigDiscoveryServiceServer() {
}

// UnsafeConfigDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigDiscoveryServiceServer will
// result in compilation errors.
type UnsafeConfigDiscoveryServiceServer interface {
	mustEmbedUnimplementedConfigDiscoveryServiceServer()
}

func RegisterConfigDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ConfigDiscoveryServiceServer) {
	s.RegisterService(&_ConfigDiscoveryService_serviceDesc, srv)
}

func _ConfigDiscoveryService_StreamConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConfigDiscoveryServiceServer).StreamConfigs(&configDiscoveryServiceStreamConfigsServer{stream})
}

type ConfigDiscoveryService_StreamConfigsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type configDiscoveryServiceStreamConfigsServer struct {
	grpc.ServerStream
}

func (x *configDiscoveryServiceStreamConfigsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *configDiscoveryServiceStreamConfigsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConfigDiscoveryService_FetchConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigDiscoveryServiceServer).FetchConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.service.config.ConfigDiscoveryService/FetchConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigDiscoveryServiceServer).FetchConfigs(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.service.config.ConfigDiscoveryService",
	HandlerType: (*ConfigDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchConfigs",
			Handler:    _ConfigDiscoveryService_FetchConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamConfigs",
			Handler:       _ConfigDiscoveryService_StreamConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "wso2/discovery/service/config/cds.proto",
}
