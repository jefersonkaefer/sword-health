// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.12.2
// source: infra/grpc/task.proto

package grpc_task

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	CreateTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	ListTasksRequest(ctx context.Context, in *TasksListRequest, opts ...grpc.CallOption) (*TaskList, error)
	FindOneTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	UpdateTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) CreateTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/grpc_task.TaskService/CreateTaskRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTasksRequest(ctx context.Context, in *TasksListRequest, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := c.cc.Invoke(ctx, "/grpc_task.TaskService/ListTasksRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) FindOneTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/grpc_task.TaskService/FindOneTaskRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) UpdateTaskRequest(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/grpc_task.TaskService/UpdateTaskRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	CreateTaskRequest(context.Context, *TaskRequest) (*Task, error)
	ListTasksRequest(context.Context, *TasksListRequest) (*TaskList, error)
	FindOneTaskRequest(context.Context, *TaskRequest) (*Task, error)
	UpdateTaskRequest(context.Context, *TaskRequest) (*Task, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) CreateTaskRequest(context.Context, *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskRequest not implemented")
}
func (UnimplementedTaskServiceServer) ListTasksRequest(context.Context, *TasksListRequest) (*TaskList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTasksRequest not implemented")
}
func (UnimplementedTaskServiceServer) FindOneTaskRequest(context.Context, *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneTaskRequest not implemented")
}
func (UnimplementedTaskServiceServer) UpdateTaskRequest(context.Context, *TaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskRequest not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_CreateTaskRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).CreateTaskRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_task.TaskService/CreateTaskRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).CreateTaskRequest(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTasksRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTasksRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_task.TaskService/ListTasksRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTasksRequest(ctx, req.(*TasksListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_FindOneTaskRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).FindOneTaskRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_task.TaskService/FindOneTaskRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).FindOneTaskRequest(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_UpdateTaskRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).UpdateTaskRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_task.TaskService/UpdateTaskRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).UpdateTaskRequest(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_task.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskRequest",
			Handler:    _TaskService_CreateTaskRequest_Handler,
		},
		{
			MethodName: "ListTasksRequest",
			Handler:    _TaskService_ListTasksRequest_Handler,
		},
		{
			MethodName: "FindOneTaskRequest",
			Handler:    _TaskService_FindOneTaskRequest_Handler,
		},
		{
			MethodName: "UpdateTaskRequest",
			Handler:    _TaskService_UpdateTaskRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/grpc/task.proto",
}
