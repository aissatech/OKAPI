// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: protos/pages.proto

package protos

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

// PagesClient is the client API for Pages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PagesClient interface {
	Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
	Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyResponse, error)
}

type pagesClient struct {
	cc grpc.ClientConnInterface
}

func NewPagesClient(cc grpc.ClientConnInterface) PagesClient {
	return &pagesClient{cc}
}

func (c *pagesClient) Index(ctx context.Context, in *IndexRequest, opts ...grpc.CallOption) (*IndexResponse, error) {
	out := new(IndexResponse)
	err := c.cc.Invoke(ctx, "/pages.Pages/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagesClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/pages.Pages/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagesClient) Export(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, "/pages.Pages/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pagesClient) Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyResponse, error) {
	out := new(CopyResponse)
	err := c.cc.Invoke(ctx, "/pages.Pages/Copy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PagesServer is the server API for Pages service.
// All implementations must embed UnimplementedPagesServer
// for forward compatibility
type PagesServer interface {
	Index(context.Context, *IndexRequest) (*IndexResponse, error)
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	Export(context.Context, *ExportRequest) (*ExportResponse, error)
	Copy(context.Context, *CopyRequest) (*CopyResponse, error)
	mustEmbedUnimplementedPagesServer()
}

// UnimplementedPagesServer must be embedded to have forward compatible implementations.
type UnimplementedPagesServer struct {
}

func (UnimplementedPagesServer) Index(context.Context, *IndexRequest) (*IndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (UnimplementedPagesServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedPagesServer) Export(context.Context, *ExportRequest) (*ExportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}
func (UnimplementedPagesServer) Copy(context.Context, *CopyRequest) (*CopyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedPagesServer) mustEmbedUnimplementedPagesServer() {}

// UnsafePagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PagesServer will
// result in compilation errors.
type UnsafePagesServer interface {
	mustEmbedUnimplementedPagesServer()
}

func RegisterPagesServer(s grpc.ServiceRegistrar, srv PagesServer) {
	s.RegisterService(&Pages_ServiceDesc, srv)
}

func _Pages_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pages.Pages/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).Index(ctx, req.(*IndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pages_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pages.Pages/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pages_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pages.Pages/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).Export(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pages_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PagesServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pages.Pages/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PagesServer).Copy(ctx, req.(*CopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pages_ServiceDesc is the grpc.ServiceDesc for Pages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pages.Pages",
	HandlerType: (*PagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _Pages_Index_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Pages_Fetch_Handler,
		},
		{
			MethodName: "Export",
			Handler:    _Pages_Export_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _Pages_Copy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pages.proto",
}
