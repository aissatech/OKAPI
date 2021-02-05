// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NamespacesClient is the client API for Namespaces service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespacesClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
}

type namespacesClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespacesClient(cc grpc.ClientConnInterface) NamespacesClient {
	return &namespacesClient{cc}
}

func (c *namespacesClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/namespaces.Namespaces/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespacesServer is the server API for Namespaces service.
// All implementations must embed UnimplementedNamespacesServer
// for forward compatibility
type NamespacesServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	mustEmbedUnimplementedNamespacesServer()
}

// UnimplementedNamespacesServer must be embedded to have forward compatible implementations.
type UnimplementedNamespacesServer struct {
}

func (UnimplementedNamespacesServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedNamespacesServer) mustEmbedUnimplementedNamespacesServer() {}

// UnsafeNamespacesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespacesServer will
// result in compilation errors.
type UnsafeNamespacesServer interface {
	mustEmbedUnimplementedNamespacesServer()
}

func RegisterNamespacesServer(s grpc.ServiceRegistrar, srv NamespacesServer) {
	s.RegisterService(&_Namespaces_serviceDesc, srv)
}

func _Namespaces_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/namespaces.Namespaces/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Namespaces_serviceDesc = grpc.ServiceDesc{
	ServiceName: "namespaces.Namespaces",
	HandlerType: (*NamespacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _Namespaces_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/namespaces.proto",
}