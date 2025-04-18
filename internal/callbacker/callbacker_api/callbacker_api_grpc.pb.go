// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: internal/callbacker/callbacker_api/callbacker_api.proto

package callbacker_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CallbackerAPI_Health_FullMethodName          = "/callbacker_api.CallbackerAPI/Health"
	CallbackerAPI_SendCallback_FullMethodName    = "/callbacker_api.CallbackerAPI/SendCallback"
	CallbackerAPI_UpdateInstances_FullMethodName = "/callbacker_api.CallbackerAPI/UpdateInstances"
)

// CallbackerAPIClient is the client API for CallbackerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallbackerAPIClient interface {
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	SendCallback(ctx context.Context, in *SendCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateInstances(ctx context.Context, in *UpdateInstancesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type callbackerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCallbackerAPIClient(cc grpc.ClientConnInterface) CallbackerAPIClient {
	return &callbackerAPIClient{cc}
}

func (c *callbackerAPIClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, CallbackerAPI_Health_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackerAPIClient) SendCallback(ctx context.Context, in *SendCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CallbackerAPI_SendCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackerAPIClient) UpdateInstances(ctx context.Context, in *UpdateInstancesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CallbackerAPI_UpdateInstances_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallbackerAPIServer is the server API for CallbackerAPI service.
// All implementations must embed UnimplementedCallbackerAPIServer
// for forward compatibility.
type CallbackerAPIServer interface {
	Health(context.Context, *emptypb.Empty) (*HealthResponse, error)
	SendCallback(context.Context, *SendCallbackRequest) (*emptypb.Empty, error)
	UpdateInstances(context.Context, *UpdateInstancesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCallbackerAPIServer()
}

// UnimplementedCallbackerAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCallbackerAPIServer struct{}

func (UnimplementedCallbackerAPIServer) Health(context.Context, *emptypb.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCallbackerAPIServer) SendCallback(context.Context, *SendCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCallback not implemented")
}
func (UnimplementedCallbackerAPIServer) UpdateInstances(context.Context, *UpdateInstancesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstances not implemented")
}
func (UnimplementedCallbackerAPIServer) mustEmbedUnimplementedCallbackerAPIServer() {}
func (UnimplementedCallbackerAPIServer) testEmbeddedByValue()                       {}

// UnsafeCallbackerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallbackerAPIServer will
// result in compilation errors.
type UnsafeCallbackerAPIServer interface {
	mustEmbedUnimplementedCallbackerAPIServer()
}

func RegisterCallbackerAPIServer(s grpc.ServiceRegistrar, srv CallbackerAPIServer) {
	// If the following call panics, it indicates UnimplementedCallbackerAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CallbackerAPI_ServiceDesc, srv)
}

func _CallbackerAPI_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackerAPIServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CallbackerAPI_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackerAPIServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackerAPI_SendCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackerAPIServer).SendCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CallbackerAPI_SendCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackerAPIServer).SendCallback(ctx, req.(*SendCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackerAPI_UpdateInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackerAPIServer).UpdateInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CallbackerAPI_UpdateInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackerAPIServer).UpdateInstances(ctx, req.(*UpdateInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CallbackerAPI_ServiceDesc is the grpc.ServiceDesc for CallbackerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallbackerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "callbacker_api.CallbackerAPI",
	HandlerType: (*CallbackerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _CallbackerAPI_Health_Handler,
		},
		{
			MethodName: "SendCallback",
			Handler:    _CallbackerAPI_SendCallback_Handler,
		},
		{
			MethodName: "UpdateInstances",
			Handler:    _CallbackerAPI_UpdateInstances_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/callbacker/callbacker_api/callbacker_api.proto",
}
