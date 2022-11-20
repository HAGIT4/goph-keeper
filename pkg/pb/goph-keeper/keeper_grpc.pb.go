// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: keeper.proto

package goph_keeper

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/keeper.Auth/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/keeper.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.Auth/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}

// LoginPassKeeperClient is the client API for LoginPassKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginPassKeeperClient interface {
	SaveLoginPass(ctx context.Context, in *SaveLoginPassRequest, opts ...grpc.CallOption) (*SaveLoginPassResponse, error)
	GetLoginPass(ctx context.Context, in *GetLoginPassRequest, opts ...grpc.CallOption) (*GetLoginPassResponse, error)
	ListLoginPassKeywords(ctx context.Context, in *ListLoginPassKeywordsRequest, opts ...grpc.CallOption) (*ListLoginPassKeywordsResponse, error)
}

type loginPassKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginPassKeeperClient(cc grpc.ClientConnInterface) LoginPassKeeperClient {
	return &loginPassKeeperClient{cc}
}

func (c *loginPassKeeperClient) SaveLoginPass(ctx context.Context, in *SaveLoginPassRequest, opts ...grpc.CallOption) (*SaveLoginPassResponse, error) {
	out := new(SaveLoginPassResponse)
	err := c.cc.Invoke(ctx, "/keeper.LoginPassKeeper/SaveLoginPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginPassKeeperClient) GetLoginPass(ctx context.Context, in *GetLoginPassRequest, opts ...grpc.CallOption) (*GetLoginPassResponse, error) {
	out := new(GetLoginPassResponse)
	err := c.cc.Invoke(ctx, "/keeper.LoginPassKeeper/GetLoginPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginPassKeeperClient) ListLoginPassKeywords(ctx context.Context, in *ListLoginPassKeywordsRequest, opts ...grpc.CallOption) (*ListLoginPassKeywordsResponse, error) {
	out := new(ListLoginPassKeywordsResponse)
	err := c.cc.Invoke(ctx, "/keeper.LoginPassKeeper/ListLoginPassKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginPassKeeperServer is the server API for LoginPassKeeper service.
// All implementations must embed UnimplementedLoginPassKeeperServer
// for forward compatibility
type LoginPassKeeperServer interface {
	SaveLoginPass(context.Context, *SaveLoginPassRequest) (*SaveLoginPassResponse, error)
	GetLoginPass(context.Context, *GetLoginPassRequest) (*GetLoginPassResponse, error)
	ListLoginPassKeywords(context.Context, *ListLoginPassKeywordsRequest) (*ListLoginPassKeywordsResponse, error)
	mustEmbedUnimplementedLoginPassKeeperServer()
}

// UnimplementedLoginPassKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedLoginPassKeeperServer struct {
}

func (UnimplementedLoginPassKeeperServer) SaveLoginPass(context.Context, *SaveLoginPassRequest) (*SaveLoginPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLoginPass not implemented")
}
func (UnimplementedLoginPassKeeperServer) GetLoginPass(context.Context, *GetLoginPassRequest) (*GetLoginPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoginPass not implemented")
}
func (UnimplementedLoginPassKeeperServer) ListLoginPassKeywords(context.Context, *ListLoginPassKeywordsRequest) (*ListLoginPassKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLoginPassKeywords not implemented")
}
func (UnimplementedLoginPassKeeperServer) mustEmbedUnimplementedLoginPassKeeperServer() {}

// UnsafeLoginPassKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginPassKeeperServer will
// result in compilation errors.
type UnsafeLoginPassKeeperServer interface {
	mustEmbedUnimplementedLoginPassKeeperServer()
}

func RegisterLoginPassKeeperServer(s grpc.ServiceRegistrar, srv LoginPassKeeperServer) {
	s.RegisterService(&LoginPassKeeper_ServiceDesc, srv)
}

func _LoginPassKeeper_SaveLoginPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLoginPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginPassKeeperServer).SaveLoginPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.LoginPassKeeper/SaveLoginPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginPassKeeperServer).SaveLoginPass(ctx, req.(*SaveLoginPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginPassKeeper_GetLoginPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLoginPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginPassKeeperServer).GetLoginPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.LoginPassKeeper/GetLoginPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginPassKeeperServer).GetLoginPass(ctx, req.(*GetLoginPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginPassKeeper_ListLoginPassKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLoginPassKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginPassKeeperServer).ListLoginPassKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.LoginPassKeeper/ListLoginPassKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginPassKeeperServer).ListLoginPassKeywords(ctx, req.(*ListLoginPassKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginPassKeeper_ServiceDesc is the grpc.ServiceDesc for LoginPassKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginPassKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.LoginPassKeeper",
	HandlerType: (*LoginPassKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveLoginPass",
			Handler:    _LoginPassKeeper_SaveLoginPass_Handler,
		},
		{
			MethodName: "GetLoginPass",
			Handler:    _LoginPassKeeper_GetLoginPass_Handler,
		},
		{
			MethodName: "ListLoginPassKeywords",
			Handler:    _LoginPassKeeper_ListLoginPassKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}

// TextDataKeeperClient is the client API for TextDataKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextDataKeeperClient interface {
	SaveTextData(ctx context.Context, in *SaveTextDataRequest, opts ...grpc.CallOption) (*SaveTextDataResponse, error)
	GetTextData(ctx context.Context, in *GetTextDataRequest, opts ...grpc.CallOption) (*GetTextDataResponse, error)
}

type textDataKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewTextDataKeeperClient(cc grpc.ClientConnInterface) TextDataKeeperClient {
	return &textDataKeeperClient{cc}
}

func (c *textDataKeeperClient) SaveTextData(ctx context.Context, in *SaveTextDataRequest, opts ...grpc.CallOption) (*SaveTextDataResponse, error) {
	out := new(SaveTextDataResponse)
	err := c.cc.Invoke(ctx, "/keeper.TextDataKeeper/SaveTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textDataKeeperClient) GetTextData(ctx context.Context, in *GetTextDataRequest, opts ...grpc.CallOption) (*GetTextDataResponse, error) {
	out := new(GetTextDataResponse)
	err := c.cc.Invoke(ctx, "/keeper.TextDataKeeper/GetTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextDataKeeperServer is the server API for TextDataKeeper service.
// All implementations must embed UnimplementedTextDataKeeperServer
// for forward compatibility
type TextDataKeeperServer interface {
	SaveTextData(context.Context, *SaveTextDataRequest) (*SaveTextDataResponse, error)
	GetTextData(context.Context, *GetTextDataRequest) (*GetTextDataResponse, error)
	mustEmbedUnimplementedTextDataKeeperServer()
}

// UnimplementedTextDataKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedTextDataKeeperServer struct {
}

func (UnimplementedTextDataKeeperServer) SaveTextData(context.Context, *SaveTextDataRequest) (*SaveTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTextData not implemented")
}
func (UnimplementedTextDataKeeperServer) GetTextData(context.Context, *GetTextDataRequest) (*GetTextDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextData not implemented")
}
func (UnimplementedTextDataKeeperServer) mustEmbedUnimplementedTextDataKeeperServer() {}

// UnsafeTextDataKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextDataKeeperServer will
// result in compilation errors.
type UnsafeTextDataKeeperServer interface {
	mustEmbedUnimplementedTextDataKeeperServer()
}

func RegisterTextDataKeeperServer(s grpc.ServiceRegistrar, srv TextDataKeeperServer) {
	s.RegisterService(&TextDataKeeper_ServiceDesc, srv)
}

func _TextDataKeeper_SaveTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextDataKeeperServer).SaveTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.TextDataKeeper/SaveTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextDataKeeperServer).SaveTextData(ctx, req.(*SaveTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextDataKeeper_GetTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextDataKeeperServer).GetTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.TextDataKeeper/GetTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextDataKeeperServer).GetTextData(ctx, req.(*GetTextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextDataKeeper_ServiceDesc is the grpc.ServiceDesc for TextDataKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextDataKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.TextDataKeeper",
	HandlerType: (*TextDataKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveTextData",
			Handler:    _TextDataKeeper_SaveTextData_Handler,
		},
		{
			MethodName: "GetTextData",
			Handler:    _TextDataKeeper_GetTextData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}

// CardDataKeeperClient is the client API for CardDataKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardDataKeeperClient interface {
	SaveCardData(ctx context.Context, in *SaveCardDataRequest, opts ...grpc.CallOption) (*SaveCardDataResponse, error)
	GetCardData(ctx context.Context, in *GetCardDataRequest, opts ...grpc.CallOption) (*GetCardDataResponse, error)
	ListCardDataKeywords(ctx context.Context, in *ListCardDataKeywordsRequest, opts ...grpc.CallOption) (*ListCardDataKeywordsResponse, error)
}

type cardDataKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewCardDataKeeperClient(cc grpc.ClientConnInterface) CardDataKeeperClient {
	return &cardDataKeeperClient{cc}
}

func (c *cardDataKeeperClient) SaveCardData(ctx context.Context, in *SaveCardDataRequest, opts ...grpc.CallOption) (*SaveCardDataResponse, error) {
	out := new(SaveCardDataResponse)
	err := c.cc.Invoke(ctx, "/keeper.CardDataKeeper/SaveCardData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDataKeeperClient) GetCardData(ctx context.Context, in *GetCardDataRequest, opts ...grpc.CallOption) (*GetCardDataResponse, error) {
	out := new(GetCardDataResponse)
	err := c.cc.Invoke(ctx, "/keeper.CardDataKeeper/GetCardData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardDataKeeperClient) ListCardDataKeywords(ctx context.Context, in *ListCardDataKeywordsRequest, opts ...grpc.CallOption) (*ListCardDataKeywordsResponse, error) {
	out := new(ListCardDataKeywordsResponse)
	err := c.cc.Invoke(ctx, "/keeper.CardDataKeeper/ListCardDataKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardDataKeeperServer is the server API for CardDataKeeper service.
// All implementations must embed UnimplementedCardDataKeeperServer
// for forward compatibility
type CardDataKeeperServer interface {
	SaveCardData(context.Context, *SaveCardDataRequest) (*SaveCardDataResponse, error)
	GetCardData(context.Context, *GetCardDataRequest) (*GetCardDataResponse, error)
	ListCardDataKeywords(context.Context, *ListCardDataKeywordsRequest) (*ListCardDataKeywordsResponse, error)
	mustEmbedUnimplementedCardDataKeeperServer()
}

// UnimplementedCardDataKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedCardDataKeeperServer struct {
}

func (UnimplementedCardDataKeeperServer) SaveCardData(context.Context, *SaveCardDataRequest) (*SaveCardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCardData not implemented")
}
func (UnimplementedCardDataKeeperServer) GetCardData(context.Context, *GetCardDataRequest) (*GetCardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardData not implemented")
}
func (UnimplementedCardDataKeeperServer) ListCardDataKeywords(context.Context, *ListCardDataKeywordsRequest) (*ListCardDataKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCardDataKeywords not implemented")
}
func (UnimplementedCardDataKeeperServer) mustEmbedUnimplementedCardDataKeeperServer() {}

// UnsafeCardDataKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardDataKeeperServer will
// result in compilation errors.
type UnsafeCardDataKeeperServer interface {
	mustEmbedUnimplementedCardDataKeeperServer()
}

func RegisterCardDataKeeperServer(s grpc.ServiceRegistrar, srv CardDataKeeperServer) {
	s.RegisterService(&CardDataKeeper_ServiceDesc, srv)
}

func _CardDataKeeper_SaveCardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardDataKeeperServer).SaveCardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.CardDataKeeper/SaveCardData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardDataKeeperServer).SaveCardData(ctx, req.(*SaveCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardDataKeeper_GetCardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardDataKeeperServer).GetCardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.CardDataKeeper/GetCardData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardDataKeeperServer).GetCardData(ctx, req.(*GetCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardDataKeeper_ListCardDataKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCardDataKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardDataKeeperServer).ListCardDataKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.CardDataKeeper/ListCardDataKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardDataKeeperServer).ListCardDataKeywords(ctx, req.(*ListCardDataKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardDataKeeper_ServiceDesc is the grpc.ServiceDesc for CardDataKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardDataKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.CardDataKeeper",
	HandlerType: (*CardDataKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveCardData",
			Handler:    _CardDataKeeper_SaveCardData_Handler,
		},
		{
			MethodName: "GetCardData",
			Handler:    _CardDataKeeper_GetCardData_Handler,
		},
		{
			MethodName: "ListCardDataKeywords",
			Handler:    _CardDataKeeper_ListCardDataKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}

// BinaryKeeperClient is the client API for BinaryKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BinaryKeeperClient interface {
	SaveBinary(ctx context.Context, in *SaveBinaryRequest, opts ...grpc.CallOption) (*SaveBinaryResponse, error)
	GetBinary(ctx context.Context, in *GetBinaryRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error)
	ListBinaryKeywords(ctx context.Context, in *ListBinaryKeywordsRequest, opts ...grpc.CallOption) (*ListBinaryKeywordsResponse, error)
}

type binaryKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewBinaryKeeperClient(cc grpc.ClientConnInterface) BinaryKeeperClient {
	return &binaryKeeperClient{cc}
}

func (c *binaryKeeperClient) SaveBinary(ctx context.Context, in *SaveBinaryRequest, opts ...grpc.CallOption) (*SaveBinaryResponse, error) {
	out := new(SaveBinaryResponse)
	err := c.cc.Invoke(ctx, "/keeper.BinaryKeeper/SaveBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryKeeperClient) GetBinary(ctx context.Context, in *GetBinaryRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error) {
	out := new(GetBinaryResponse)
	err := c.cc.Invoke(ctx, "/keeper.BinaryKeeper/GetBinary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryKeeperClient) ListBinaryKeywords(ctx context.Context, in *ListBinaryKeywordsRequest, opts ...grpc.CallOption) (*ListBinaryKeywordsResponse, error) {
	out := new(ListBinaryKeywordsResponse)
	err := c.cc.Invoke(ctx, "/keeper.BinaryKeeper/ListBinaryKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinaryKeeperServer is the server API for BinaryKeeper service.
// All implementations must embed UnimplementedBinaryKeeperServer
// for forward compatibility
type BinaryKeeperServer interface {
	SaveBinary(context.Context, *SaveBinaryRequest) (*SaveBinaryResponse, error)
	GetBinary(context.Context, *GetBinaryRequest) (*GetBinaryResponse, error)
	ListBinaryKeywords(context.Context, *ListBinaryKeywordsRequest) (*ListBinaryKeywordsResponse, error)
	mustEmbedUnimplementedBinaryKeeperServer()
}

// UnimplementedBinaryKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedBinaryKeeperServer struct {
}

func (UnimplementedBinaryKeeperServer) SaveBinary(context.Context, *SaveBinaryRequest) (*SaveBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinary not implemented")
}
func (UnimplementedBinaryKeeperServer) GetBinary(context.Context, *GetBinaryRequest) (*GetBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinary not implemented")
}
func (UnimplementedBinaryKeeperServer) ListBinaryKeywords(context.Context, *ListBinaryKeywordsRequest) (*ListBinaryKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBinaryKeywords not implemented")
}
func (UnimplementedBinaryKeeperServer) mustEmbedUnimplementedBinaryKeeperServer() {}

// UnsafeBinaryKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BinaryKeeperServer will
// result in compilation errors.
type UnsafeBinaryKeeperServer interface {
	mustEmbedUnimplementedBinaryKeeperServer()
}

func RegisterBinaryKeeperServer(s grpc.ServiceRegistrar, srv BinaryKeeperServer) {
	s.RegisterService(&BinaryKeeper_ServiceDesc, srv)
}

func _BinaryKeeper_SaveBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryKeeperServer).SaveBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.BinaryKeeper/SaveBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryKeeperServer).SaveBinary(ctx, req.(*SaveBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryKeeper_GetBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBinaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryKeeperServer).GetBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.BinaryKeeper/GetBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryKeeperServer).GetBinary(ctx, req.(*GetBinaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryKeeper_ListBinaryKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBinaryKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryKeeperServer).ListBinaryKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.BinaryKeeper/ListBinaryKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryKeeperServer).ListBinaryKeywords(ctx, req.(*ListBinaryKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BinaryKeeper_ServiceDesc is the grpc.ServiceDesc for BinaryKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BinaryKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.BinaryKeeper",
	HandlerType: (*BinaryKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveBinary",
			Handler:    _BinaryKeeper_SaveBinary_Handler,
		},
		{
			MethodName: "GetBinary",
			Handler:    _BinaryKeeper_GetBinary_Handler,
		},
		{
			MethodName: "ListBinaryKeywords",
			Handler:    _BinaryKeeper_ListBinaryKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "keeper.proto",
}
