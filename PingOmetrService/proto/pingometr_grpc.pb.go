// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/pingometr.proto

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

// PingOmetrClient is the client API for PingOmetr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingOmetrClient interface {
	GetFastest(ctx context.Context, in *GetFastestRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetSlowest(ctx context.Context, in *GetSlowestRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetSpecific(ctx context.Context, in *GetSpecificRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAdminData(ctx context.Context, in *GetAdminDataRequest, opts ...grpc.CallOption) (*GetAdminDataResponse, error)
}

type pingOmetrClient struct {
	cc grpc.ClientConnInterface
}

func NewPingOmetrClient(cc grpc.ClientConnInterface) PingOmetrClient {
	return &pingOmetrClient{cc}
}

func (c *pingOmetrClient) GetFastest(ctx context.Context, in *GetFastestRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/v1.PingOmetr/GetFastest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingOmetrClient) GetSlowest(ctx context.Context, in *GetSlowestRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/v1.PingOmetr/GetSlowest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingOmetrClient) GetSpecific(ctx context.Context, in *GetSpecificRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/v1.PingOmetr/GetSpecific", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingOmetrClient) GetAdminData(ctx context.Context, in *GetAdminDataRequest, opts ...grpc.CallOption) (*GetAdminDataResponse, error) {
	out := new(GetAdminDataResponse)
	err := c.cc.Invoke(ctx, "/v1.PingOmetr/GetAdminData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingOmetrServer is the server API for PingOmetr service.
// All implementations must embed UnimplementedPingOmetrServer
// for forward compatibility
type PingOmetrServer interface {
	GetFastest(context.Context, *GetFastestRequest) (*GetResponse, error)
	GetSlowest(context.Context, *GetSlowestRequest) (*GetResponse, error)
	GetSpecific(context.Context, *GetSpecificRequest) (*GetResponse, error)
	GetAdminData(context.Context, *GetAdminDataRequest) (*GetAdminDataResponse, error)
	mustEmbedUnimplementedPingOmetrServer()
}

// UnimplementedPingOmetrServer must be embedded to have forward compatible implementations.
type UnimplementedPingOmetrServer struct {
}

func (UnimplementedPingOmetrServer) GetFastest(context.Context, *GetFastestRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFastest not implemented")
}
func (UnimplementedPingOmetrServer) GetSlowest(context.Context, *GetSlowestRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlowest not implemented")
}
func (UnimplementedPingOmetrServer) GetSpecific(context.Context, *GetSpecificRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecific not implemented")
}
func (UnimplementedPingOmetrServer) GetAdminData(context.Context, *GetAdminDataRequest) (*GetAdminDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdminData not implemented")
}
func (UnimplementedPingOmetrServer) mustEmbedUnimplementedPingOmetrServer() {}

// UnsafePingOmetrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingOmetrServer will
// result in compilation errors.
type UnsafePingOmetrServer interface {
	mustEmbedUnimplementedPingOmetrServer()
}

func RegisterPingOmetrServer(s grpc.ServiceRegistrar, srv PingOmetrServer) {
	s.RegisterService(&PingOmetr_ServiceDesc, srv)
}

func _PingOmetr_GetFastest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFastestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingOmetrServer).GetFastest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PingOmetr/GetFastest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingOmetrServer).GetFastest(ctx, req.(*GetFastestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingOmetr_GetSlowest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSlowestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingOmetrServer).GetSlowest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PingOmetr/GetSlowest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingOmetrServer).GetSlowest(ctx, req.(*GetSlowestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingOmetr_GetSpecific_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingOmetrServer).GetSpecific(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PingOmetr/GetSpecific",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingOmetrServer).GetSpecific(ctx, req.(*GetSpecificRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingOmetr_GetAdminData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdminDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingOmetrServer).GetAdminData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PingOmetr/GetAdminData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingOmetrServer).GetAdminData(ctx, req.(*GetAdminDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PingOmetr_ServiceDesc is the grpc.ServiceDesc for PingOmetr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingOmetr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PingOmetr",
	HandlerType: (*PingOmetrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFastest",
			Handler:    _PingOmetr_GetFastest_Handler,
		},
		{
			MethodName: "GetSlowest",
			Handler:    _PingOmetr_GetSlowest_Handler,
		},
		{
			MethodName: "GetSpecific",
			Handler:    _PingOmetr_GetSpecific_Handler,
		},
		{
			MethodName: "GetAdminData",
			Handler:    _PingOmetr_GetAdminData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pingometr.proto",
}
