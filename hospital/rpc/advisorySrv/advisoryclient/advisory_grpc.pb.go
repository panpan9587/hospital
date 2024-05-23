// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: rpc/doc/advisory.proto

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
	Advisory_RecordsUserConsultationInformation_FullMethodName = "/advisory.Advisory/RecordsUserConsultationInformation"
)

// AdvisoryClient is the client API for Advisory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvisoryClient interface {
	RecordsUserConsultationInformation(ctx context.Context, in *RecordsUserConsultationInformationReq, opts ...grpc.CallOption) (*RecordsUserConsultationInformationRes, error)
}

type advisoryClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvisoryClient(cc grpc.ClientConnInterface) AdvisoryClient {
	return &advisoryClient{cc}
}

func (c *advisoryClient) RecordsUserConsultationInformation(ctx context.Context, in *RecordsUserConsultationInformationReq, opts ...grpc.CallOption) (*RecordsUserConsultationInformationRes, error) {
	out := new(RecordsUserConsultationInformationRes)
	err := c.cc.Invoke(ctx, Advisory_RecordsUserConsultationInformation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvisoryServer is the server API for Advisory service.
// All implementations must embed UnimplementedAdvisoryServer
// for forward compatibility
type AdvisoryServer interface {
	RecordsUserConsultationInformation(context.Context, *RecordsUserConsultationInformationReq) (*RecordsUserConsultationInformationRes, error)
	mustEmbedUnimplementedAdvisoryServer()
}

// UnimplementedAdvisoryServer must be embedded to have forward compatible implementations.
type UnimplementedAdvisoryServer struct {
}

func (UnimplementedAdvisoryServer) RecordsUserConsultationInformation(context.Context, *RecordsUserConsultationInformationReq) (*RecordsUserConsultationInformationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordsUserConsultationInformation not implemented")
}
func (UnimplementedAdvisoryServer) mustEmbedUnimplementedAdvisoryServer() {}

// UnsafeAdvisoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvisoryServer will
// result in compilation errors.
type UnsafeAdvisoryServer interface {
	mustEmbedUnimplementedAdvisoryServer()
}

func RegisterAdvisoryServer(s grpc.ServiceRegistrar, srv AdvisoryServer) {
	s.RegisterService(&Advisory_ServiceDesc, srv)
}

func _Advisory_RecordsUserConsultationInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordsUserConsultationInformationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvisoryServer).RecordsUserConsultationInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Advisory_RecordsUserConsultationInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvisoryServer).RecordsUserConsultationInformation(ctx, req.(*RecordsUserConsultationInformationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Advisory_ServiceDesc is the grpc.ServiceDesc for Advisory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Advisory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "advisory.Advisory",
	HandlerType: (*AdvisoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordsUserConsultationInformation",
			Handler:    _Advisory_RecordsUserConsultationInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/doc/advisory.proto",
}