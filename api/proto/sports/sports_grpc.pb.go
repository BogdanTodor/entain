// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: sports/sports.proto

package sports

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

// SportsClient is the client API for Sports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SportsClient interface {
	// ListSportsEvents will return a collection of all sports events.
	ListSportsEvents(ctx context.Context, in *ListSportsEventsRequest, opts ...grpc.CallOption) (*ListSportsEventsResponse, error)
	// GetSportsEventById will return a single sports event based on the provided id
	GetSportsEventById(ctx context.Context, in *GetSportsEventRequest, opts ...grpc.CallOption) (*GetSportsEventResponse, error)
}

type sportsClient struct {
	cc grpc.ClientConnInterface
}

func NewSportsClient(cc grpc.ClientConnInterface) SportsClient {
	return &sportsClient{cc}
}

func (c *sportsClient) ListSportsEvents(ctx context.Context, in *ListSportsEventsRequest, opts ...grpc.CallOption) (*ListSportsEventsResponse, error) {
	out := new(ListSportsEventsResponse)
	err := c.cc.Invoke(ctx, "/sports.Sports/ListSportsEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sportsClient) GetSportsEventById(ctx context.Context, in *GetSportsEventRequest, opts ...grpc.CallOption) (*GetSportsEventResponse, error) {
	out := new(GetSportsEventResponse)
	err := c.cc.Invoke(ctx, "/sports.Sports/GetSportsEventById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SportsServer is the server API for Sports service.
// All implementations must embed UnimplementedSportsServer
// for forward compatibility
type SportsServer interface {
	// ListSportsEvents will return a collection of all sports events.
	ListSportsEvents(context.Context, *ListSportsEventsRequest) (*ListSportsEventsResponse, error)
	// GetSportsEventById will return a single sports event based on the provided id
	GetSportsEventById(context.Context, *GetSportsEventRequest) (*GetSportsEventResponse, error)
	mustEmbedUnimplementedSportsServer()
}

// UnimplementedSportsServer must be embedded to have forward compatible implementations.
type UnimplementedSportsServer struct {
}

func (UnimplementedSportsServer) ListSportsEvents(context.Context, *ListSportsEventsRequest) (*ListSportsEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSportsEvents not implemented")
}
func (UnimplementedSportsServer) GetSportsEventById(context.Context, *GetSportsEventRequest) (*GetSportsEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSportsEventById not implemented")
}
func (UnimplementedSportsServer) mustEmbedUnimplementedSportsServer() {}

// UnsafeSportsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SportsServer will
// result in compilation errors.
type UnsafeSportsServer interface {
	mustEmbedUnimplementedSportsServer()
}

func RegisterSportsServer(s grpc.ServiceRegistrar, srv SportsServer) {
	s.RegisterService(&Sports_ServiceDesc, srv)
}

func _Sports_ListSportsEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSportsEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportsServer).ListSportsEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sports.Sports/ListSportsEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportsServer).ListSportsEvents(ctx, req.(*ListSportsEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sports_GetSportsEventById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSportsEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportsServer).GetSportsEventById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sports.Sports/GetSportsEventById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportsServer).GetSportsEventById(ctx, req.(*GetSportsEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sports_ServiceDesc is the grpc.ServiceDesc for Sports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sports.Sports",
	HandlerType: (*SportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSportsEvents",
			Handler:    _Sports_ListSportsEvents_Handler,
		},
		{
			MethodName: "GetSportsEventById",
			Handler:    _Sports_GetSportsEventById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sports/sports.proto",
}
