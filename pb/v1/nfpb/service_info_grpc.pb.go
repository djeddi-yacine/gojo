// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/nfpb/service_info.proto

package nfpbv1

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
	InfoService_CreateGenres_FullMethodName    = "/v1.nfpbv1.InfoService/CreateGenres"
	InfoService_CreateStudios_FullMethodName   = "/v1.nfpbv1.InfoService/CreateStudios"
	InfoService_CreateLanguages_FullMethodName = "/v1.nfpbv1.InfoService/CreateLanguages"
	InfoService_GetAllGenres_FullMethodName    = "/v1.nfpbv1.InfoService/GetAllGenres"
	InfoService_GetAllStudios_FullMethodName   = "/v1.nfpbv1.InfoService/GetAllStudios"
	InfoService_GetAllLanguages_FullMethodName = "/v1.nfpbv1.InfoService/GetAllLanguages"
)

// InfoServiceClient is the client API for InfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfoServiceClient interface {
	CreateGenres(ctx context.Context, in *CreateGenresRequest, opts ...grpc.CallOption) (*CreateGenresResponse, error)
	CreateStudios(ctx context.Context, in *CreateStudiosRequest, opts ...grpc.CallOption) (*CreateStudiosResponse, error)
	CreateLanguages(ctx context.Context, in *CreateLanguagesRequest, opts ...grpc.CallOption) (*CreateLanguagesResponse, error)
	GetAllGenres(ctx context.Context, in *GetAllGenresRequest, opts ...grpc.CallOption) (*GetAllGenresResponse, error)
	GetAllStudios(ctx context.Context, in *GetAllStudiosRequest, opts ...grpc.CallOption) (*GetAllStudiosResponse, error)
	GetAllLanguages(ctx context.Context, in *GetAllLanguagesRequest, opts ...grpc.CallOption) (*GetAllLanguagesResponse, error)
}

type infoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfoServiceClient(cc grpc.ClientConnInterface) InfoServiceClient {
	return &infoServiceClient{cc}
}

func (c *infoServiceClient) CreateGenres(ctx context.Context, in *CreateGenresRequest, opts ...grpc.CallOption) (*CreateGenresResponse, error) {
	out := new(CreateGenresResponse)
	err := c.cc.Invoke(ctx, InfoService_CreateGenres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) CreateStudios(ctx context.Context, in *CreateStudiosRequest, opts ...grpc.CallOption) (*CreateStudiosResponse, error) {
	out := new(CreateStudiosResponse)
	err := c.cc.Invoke(ctx, InfoService_CreateStudios_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) CreateLanguages(ctx context.Context, in *CreateLanguagesRequest, opts ...grpc.CallOption) (*CreateLanguagesResponse, error) {
	out := new(CreateLanguagesResponse)
	err := c.cc.Invoke(ctx, InfoService_CreateLanguages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAllGenres(ctx context.Context, in *GetAllGenresRequest, opts ...grpc.CallOption) (*GetAllGenresResponse, error) {
	out := new(GetAllGenresResponse)
	err := c.cc.Invoke(ctx, InfoService_GetAllGenres_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAllStudios(ctx context.Context, in *GetAllStudiosRequest, opts ...grpc.CallOption) (*GetAllStudiosResponse, error) {
	out := new(GetAllStudiosResponse)
	err := c.cc.Invoke(ctx, InfoService_GetAllStudios_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infoServiceClient) GetAllLanguages(ctx context.Context, in *GetAllLanguagesRequest, opts ...grpc.CallOption) (*GetAllLanguagesResponse, error) {
	out := new(GetAllLanguagesResponse)
	err := c.cc.Invoke(ctx, InfoService_GetAllLanguages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfoServiceServer is the server API for InfoService service.
// All implementations must embed UnimplementedInfoServiceServer
// for forward compatibility
type InfoServiceServer interface {
	CreateGenres(context.Context, *CreateGenresRequest) (*CreateGenresResponse, error)
	CreateStudios(context.Context, *CreateStudiosRequest) (*CreateStudiosResponse, error)
	CreateLanguages(context.Context, *CreateLanguagesRequest) (*CreateLanguagesResponse, error)
	GetAllGenres(context.Context, *GetAllGenresRequest) (*GetAllGenresResponse, error)
	GetAllStudios(context.Context, *GetAllStudiosRequest) (*GetAllStudiosResponse, error)
	GetAllLanguages(context.Context, *GetAllLanguagesRequest) (*GetAllLanguagesResponse, error)
	mustEmbedUnimplementedInfoServiceServer()
}

// UnimplementedInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInfoServiceServer struct {
}

func (UnimplementedInfoServiceServer) CreateGenres(context.Context, *CreateGenresRequest) (*CreateGenresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGenres not implemented")
}
func (UnimplementedInfoServiceServer) CreateStudios(context.Context, *CreateStudiosRequest) (*CreateStudiosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudios not implemented")
}
func (UnimplementedInfoServiceServer) CreateLanguages(context.Context, *CreateLanguagesRequest) (*CreateLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLanguages not implemented")
}
func (UnimplementedInfoServiceServer) GetAllGenres(context.Context, *GetAllGenresRequest) (*GetAllGenresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGenres not implemented")
}
func (UnimplementedInfoServiceServer) GetAllStudios(context.Context, *GetAllStudiosRequest) (*GetAllStudiosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStudios not implemented")
}
func (UnimplementedInfoServiceServer) GetAllLanguages(context.Context, *GetAllLanguagesRequest) (*GetAllLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLanguages not implemented")
}
func (UnimplementedInfoServiceServer) mustEmbedUnimplementedInfoServiceServer() {}

// UnsafeInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfoServiceServer will
// result in compilation errors.
type UnsafeInfoServiceServer interface {
	mustEmbedUnimplementedInfoServiceServer()
}

func RegisterInfoServiceServer(s grpc.ServiceRegistrar, srv InfoServiceServer) {
	s.RegisterService(&InfoService_ServiceDesc, srv)
}

func _InfoService_CreateGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGenresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).CreateGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_CreateGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).CreateGenres(ctx, req.(*CreateGenresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_CreateStudios_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudiosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).CreateStudios(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_CreateStudios_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).CreateStudios(ctx, req.(*CreateStudiosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_CreateLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLanguagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).CreateLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_CreateLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).CreateLanguages(ctx, req.(*CreateLanguagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAllGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGenresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAllGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_GetAllGenres_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAllGenres(ctx, req.(*GetAllGenresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAllStudios_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllStudiosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAllStudios(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_GetAllStudios_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAllStudios(ctx, req.(*GetAllStudiosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfoService_GetAllLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllLanguagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServiceServer).GetAllLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InfoService_GetAllLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServiceServer).GetAllLanguages(ctx, req.(*GetAllLanguagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfoService_ServiceDesc is the grpc.ServiceDesc for InfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.nfpbv1.InfoService",
	HandlerType: (*InfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGenres",
			Handler:    _InfoService_CreateGenres_Handler,
		},
		{
			MethodName: "CreateStudios",
			Handler:    _InfoService_CreateStudios_Handler,
		},
		{
			MethodName: "CreateLanguages",
			Handler:    _InfoService_CreateLanguages_Handler,
		},
		{
			MethodName: "GetAllGenres",
			Handler:    _InfoService_GetAllGenres_Handler,
		},
		{
			MethodName: "GetAllStudios",
			Handler:    _InfoService_GetAllStudios_Handler,
		},
		{
			MethodName: "GetAllLanguages",
			Handler:    _InfoService_GetAllLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/nfpb/service_info.proto",
}