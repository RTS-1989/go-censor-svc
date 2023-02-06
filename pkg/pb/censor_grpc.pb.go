// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/pb/censor.proto

package pb

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

// CensorServiceClient is the client API for CensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CensorServiceClient interface {
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
}

type censorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCensorServiceClient(cc grpc.ClientConnInterface) CensorServiceClient {
	return &censorServiceClient{cc}
}

func (c *censorServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/censor.CensorService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CensorServiceServer is the server API for CensorService service.
// All implementations should embed UnimplementedCensorServiceServer
// for forward compatibility
type CensorServiceServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
}

// UnimplementedCensorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCensorServiceServer struct {
}

func (UnimplementedCensorServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}

// UnsafeCensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CensorServiceServer will
// result in compilation errors.
type UnsafeCensorServiceServer interface {
	mustEmbedUnimplementedCensorServiceServer()
}

func RegisterCensorServiceServer(s grpc.ServiceRegistrar, srv CensorServiceServer) {
	s.RegisterService(&CensorService_ServiceDesc, srv)
}

func _CensorService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CensorServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/censor.CensorService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CensorServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CensorService_ServiceDesc is the grpc.ServiceDesc for CensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "censor.CensorService",
	HandlerType: (*CensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CensorService_CreateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/censor.proto",
}
