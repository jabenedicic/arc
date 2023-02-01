// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: metamorph/metamorph_api/metamorph_api.proto

package metamorph_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetaMorphAPIClient is the client API for MetaMorphAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaMorphAPIClient interface {
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error)
	PutTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionStatus, error)
	GetTransaction(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*Transaction, error)
	GetTransactionStatus(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*TransactionStatus, error)
}

type metaMorphAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaMorphAPIClient(cc grpc.ClientConnInterface) MetaMorphAPIClient {
	return &metaMorphAPIClient{cc}
}

func (c *metaMorphAPIClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/metamorph_api.MetaMorphAPI/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaMorphAPIClient) PutTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*TransactionStatus, error) {
	out := new(TransactionStatus)
	err := c.cc.Invoke(ctx, "/metamorph_api.MetaMorphAPI/PutTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaMorphAPIClient) GetTransaction(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/metamorph_api.MetaMorphAPI/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaMorphAPIClient) GetTransactionStatus(ctx context.Context, in *TransactionStatusRequest, opts ...grpc.CallOption) (*TransactionStatus, error) {
	out := new(TransactionStatus)
	err := c.cc.Invoke(ctx, "/metamorph_api.MetaMorphAPI/GetTransactionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaMorphAPIServer is the server API for MetaMorphAPI service.
// All implementations must embed UnimplementedMetaMorphAPIServer
// for forward compatibility
type MetaMorphAPIServer interface {
	Health(context.Context, *emptypb.Empty) (*HealthResponse, error)
	PutTransaction(context.Context, *TransactionRequest) (*TransactionStatus, error)
	GetTransaction(context.Context, *TransactionStatusRequest) (*Transaction, error)
	GetTransactionStatus(context.Context, *TransactionStatusRequest) (*TransactionStatus, error)
	mustEmbedUnimplementedMetaMorphAPIServer()
}

// UnimplementedMetaMorphAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMetaMorphAPIServer struct {
}

func (UnimplementedMetaMorphAPIServer) Health(context.Context, *emptypb.Empty) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedMetaMorphAPIServer) PutTransaction(context.Context, *TransactionRequest) (*TransactionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutTransaction not implemented")
}
func (UnimplementedMetaMorphAPIServer) GetTransaction(context.Context, *TransactionStatusRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedMetaMorphAPIServer) GetTransactionStatus(context.Context, *TransactionStatusRequest) (*TransactionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionStatus not implemented")
}
func (UnimplementedMetaMorphAPIServer) mustEmbedUnimplementedMetaMorphAPIServer() {}

// UnsafeMetaMorphAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaMorphAPIServer will
// result in compilation errors.
type UnsafeMetaMorphAPIServer interface {
	mustEmbedUnimplementedMetaMorphAPIServer()
}

func RegisterMetaMorphAPIServer(s grpc.ServiceRegistrar, srv MetaMorphAPIServer) {
	s.RegisterService(&MetaMorphAPI_ServiceDesc, srv)
}

func _MetaMorphAPI_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaMorphAPIServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metamorph_api.MetaMorphAPI/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaMorphAPIServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaMorphAPI_PutTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaMorphAPIServer).PutTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metamorph_api.MetaMorphAPI/PutTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaMorphAPIServer).PutTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaMorphAPI_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaMorphAPIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metamorph_api.MetaMorphAPI/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaMorphAPIServer).GetTransaction(ctx, req.(*TransactionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaMorphAPI_GetTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaMorphAPIServer).GetTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metamorph_api.MetaMorphAPI/GetTransactionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaMorphAPIServer).GetTransactionStatus(ctx, req.(*TransactionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetaMorphAPI_ServiceDesc is the grpc.ServiceDesc for MetaMorphAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaMorphAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metamorph_api.MetaMorphAPI",
	HandlerType: (*MetaMorphAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _MetaMorphAPI_Health_Handler,
		},
		{
			MethodName: "PutTransaction",
			Handler:    _MetaMorphAPI_PutTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _MetaMorphAPI_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionStatus",
			Handler:    _MetaMorphAPI_GetTransactionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metamorph/metamorph_api/metamorph_api.proto",
}
