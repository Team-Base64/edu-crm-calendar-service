// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: delivery/calendar/proto/calendar.proto

// export PATH="$PATH:$(go env GOPATH)/bin"
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative delivery/calendar/proto/calendar.proto

package calendar

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
	CalendarController_GetEventsCalendar_FullMethodName = "/calendar.CalendarController/GetEventsCalendar"
	CalendarController_CreateEvent_FullMethodName       = "/calendar.CalendarController/CreateEvent"
	CalendarController_UpdateEvent_FullMethodName       = "/calendar.CalendarController/UpdateEvent"
	CalendarController_DeleteEvent_FullMethodName       = "/calendar.CalendarController/DeleteEvent"
	CalendarController_CreateCalendar_FullMethodName    = "/calendar.CalendarController/CreateCalendar"
)

// CalendarControllerClient is the client API for CalendarController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarControllerClient interface {
	GetEventsCalendar(ctx context.Context, in *GetEventsRequestCalendar, opts ...grpc.CallOption) (*GetEventsResponse, error)
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Nothing, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*Nothing, error)
	CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type calendarControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarControllerClient(cc grpc.ClientConnInterface) CalendarControllerClient {
	return &calendarControllerClient{cc}
}

func (c *calendarControllerClient) GetEventsCalendar(ctx context.Context, in *GetEventsRequestCalendar, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, CalendarController_GetEventsCalendar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarControllerClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, CalendarController_CreateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarControllerClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Nothing, error) {
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

func (c *calendarControllerClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, CalendarController_CreateCalendar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarControllerServer is the server API for CalendarController service.
// All implementations must embed UnimplementedCalendarControllerServer
// for forward compatibility
type CalendarControllerServer interface {
	GetEventsCalendar(context.Context, *GetEventsRequestCalendar) (*GetEventsResponse, error)
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*Nothing, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*Nothing, error)
	CreateCalendar(context.Context, *CreateCalendarRequest) (*Nothing, error)
	mustEmbedUnimplementedCalendarControllerServer()
}

// UnimplementedCalendarControllerServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarControllerServer struct {
}

func (UnimplementedCalendarControllerServer) GetEventsCalendar(context.Context, *GetEventsRequestCalendar) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsCalendar not implemented")
}
func (UnimplementedCalendarControllerServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedCalendarControllerServer) UpdateEvent(context.Context, *UpdateEventRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedCalendarControllerServer) DeleteEvent(context.Context, *DeleteEventRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (UnimplementedCalendarControllerServer) CreateCalendar(context.Context, *CreateCalendarRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCalendar not implemented")
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

func _CalendarController_GetEventsCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequestCalendar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).GetEventsCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_GetEventsCalendar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).GetEventsCalendar(ctx, req.(*GetEventsRequestCalendar))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarController_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
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
		return srv.(CalendarControllerServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarController_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
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
		return srv.(CalendarControllerServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
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

func _CalendarController_CreateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarControllerServer).CreateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalendarController_CreateCalendar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarControllerServer).CreateCalendar(ctx, req.(*CreateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarController_ServiceDesc is the grpc.ServiceDesc for CalendarController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.CalendarController",
	HandlerType: (*CalendarControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventsCalendar",
			Handler:    _CalendarController_GetEventsCalendar_Handler,
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
		{
			MethodName: "CreateCalendar",
			Handler:    _CalendarController_CreateCalendar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "delivery/calendar/proto/calendar.proto",
}
