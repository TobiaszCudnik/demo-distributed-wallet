// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: api/wallet_pb/wallet.proto

package wallet_pb

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

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	AlterBalance(ctx context.Context, in *AlterBalanceReq, opts ...grpc.CallOption) (*AlterBalanceRes, error)
	GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceRes, error)
	CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*CreateWalletRes, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) AlterBalance(ctx context.Context, in *AlterBalanceReq, opts ...grpc.CallOption) (*AlterBalanceRes, error) {
	out := new(AlterBalanceRes)
	err := c.cc.Invoke(ctx, "/wallet_pb.WalletService/AlterBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetBalance(ctx context.Context, in *GetBalanceReq, opts ...grpc.CallOption) (*GetBalanceRes, error) {
	out := new(GetBalanceRes)
	err := c.cc.Invoke(ctx, "/wallet_pb.WalletService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) CreateWallet(ctx context.Context, in *CreateWalletReq, opts ...grpc.CallOption) (*CreateWalletRes, error) {
	out := new(CreateWalletRes)
	err := c.cc.Invoke(ctx, "/wallet_pb.WalletService/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	AlterBalance(context.Context, *AlterBalanceReq) (*AlterBalanceRes, error)
	GetBalance(context.Context, *GetBalanceReq) (*GetBalanceRes, error)
	CreateWallet(context.Context, *CreateWalletReq) (*CreateWalletRes, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) AlterBalance(context.Context, *AlterBalanceReq) (*AlterBalanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterBalance not implemented")
}
func (UnimplementedWalletServiceServer) GetBalance(context.Context, *GetBalanceReq) (*GetBalanceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedWalletServiceServer) CreateWallet(context.Context, *CreateWalletReq) (*CreateWalletRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_AlterBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).AlterBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_pb.WalletService/AlterBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).AlterBalance(ctx, req.(*AlterBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_pb.WalletService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetBalance(ctx, req.(*GetBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet_pb.WalletService/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).CreateWallet(ctx, req.(*CreateWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wallet_pb.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AlterBalance",
			Handler:    _WalletService_AlterBalance_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _WalletService_GetBalance_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _WalletService_CreateWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/wallet_pb/wallet.proto",
}