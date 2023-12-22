// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: controller/ctrl.proto

// export PATH="$PATH:$(go env GOPATH)/bin"
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative controller/ctrl.proto

package ctrl

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
	CalendarController_GetEvents_FullMethodName   = "/ctrl.CalendarController/GetEvents"
	CalendarController_CreateEvent_FullMethodName = "/ctrl.CalendarController/CreateEvent"
	CalendarController_UpdateEvent_FullMethodName = "/ctrl.CalendarController/UpdateEvent"
	CalendarController_DeleteEvent_FullMethodName = "/ctrl.CalendarController/DeleteEvent"
)

// CalendarControllerClient is the client API for CalendarController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarControllerClient interface {
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	CreateEvent(ctx context.Context, in *EventData, opts ...grpc.CallOption) (*Nothing, error)
	UpdateEvent(ctx context.Context, in *EventData, opts ...grpc.CallOption) (*Nothing, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type calendarControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarControllerClient(cc grpc.ClientConnInterface) CalendarControllerClient {
	return &calendarControllerClient{cc}
}

func (c *calendarControllerClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, CalendarController_GetEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarControllerClient) CreateEvent(ctx context.Context, in *EventData, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, CalendarController_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarControllerClient) UpdateEvent(ctx context.Context, in *EventData, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, CalendarController_UpdateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarControllerClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, CalendarController_DeleteEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarControllerServer is the server API for CalendarController service.
// All implementations must embed UnimplementedCalendarControllerServer
// for forward compatibility
type CalendarControllerServer interface {
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	CreateEvent(context.Context, *EventData) (*Nothing, error)
	UpdateEvent(context.Context, *EventData) (*Nothing, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*Nothing, error)
	mustEmbedUnimplementedCalendarControllerServer()
}

// UnimplementedCalendarControllerServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarControllerServer struct {
}

func (UnimplementedCalendarControllerServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedCalendarControllerServer) CreateEvent(context.Context, *EventData) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarControllerServer) UpdateEvent(context.Context, *EventData) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarControllerServer) DeleteEvent(context.Context, *DeleteEventRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarControllerServer) mustEmbedUnimplementedCalendarControllerServer() {}

// UnsafeCalendarControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarControllerServer will
// result in compilation errors.
type UnsafeCalendarControllerServer interface {
	mustEmbedUnimplementedCalendarControllerServer()
}

func RegisterCalendarControllerServer(s grpc.ServiceRegistrar, srv CalendarControllerServer) {
	s.RegisterService(&CalendarController_ServiceDesc, srv)
}

func _CalendarController_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_GetEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarController_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_CreateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).CreateEvent(ctx, req.(*EventData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarController_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_UpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).UpdateEvent(ctx, req.(*EventData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarController_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_DeleteEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarController_ServiceDesc is the grpc.ServiceDesc for CalendarController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ctrl.CalendarController",
	HandlerType: (*CalendarControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _CalendarController_GetEvents_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarController_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _CalendarController_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _CalendarController_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/ctrl.proto",
}
