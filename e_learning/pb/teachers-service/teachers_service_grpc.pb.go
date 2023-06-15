// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: teachers_service.proto

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

const (
	TeachersService_GetTeacher_FullMethodName      = "/auth.TeachersService/GetTeacher"
	TeachersService_GetTeachers_FullMethodName     = "/auth.TeachersService/GetTeachers"
	TeachersService_CreateTeacher_FullMethodName   = "/auth.TeachersService/CreateTeacher"
	TeachersService_DeleteTeacher_FullMethodName = "/auth.TeachersService/DeleteTeacher"
	TeachersService_UpdateTeacher_FullMethodName = "/auth.TeachersService/UpdateTeacher"
)

// TeachersServiceClient is the client API for TeachersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeachersServiceClient interface {
	GetTeacher(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*Teacher, error)
	GetTeachers(ctx context.Context, in *GetTeachersRequest, opts ...grpc.CallOption) (*TeacherList, error)
	CreateTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error)
	DeleteTeacher(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*Teacher, error)
	UpdateTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error)
}

type teachersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeachersServiceClient(cc grpc.ClientConnInterface) TeachersServiceClient {
	return &teachersServiceClient{cc}
}

func (c *teachersServiceClient) GetTeacher(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, TeachersService_GetTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersServiceClient) GetTeachers(ctx context.Context, in *GetTeachersRequest, opts ...grpc.CallOption) (*TeacherList, error) {
	out := new(TeacherList)
	err := c.cc.Invoke(ctx, TeachersService_GetTeachers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersServiceClient) CreateTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, TeachersService_CreateTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersServiceClient) DeleteTeacher(ctx context.Context, in *TeacherId, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, TeachersService_DeleteTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersServiceClient) UpdateTeacher(ctx context.Context, in *Teacher, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, TeachersService_UpdateTeacher_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachersServiceServer is the server API for TeachersService service.
// All implementations must embed UnimplementedTeachersServiceServer
// for forward compatibility
type TeachersServiceServer interface {
	GetTeacher(context.Context, *TeacherId) (*Teacher, error)
	GetTeachers(context.Context, *GetTeachersRequest) (*TeacherList, error)
	CreateTeacher(context.Context, *Teacher) (*Teacher, error)
	DeleteTeacher(context.Context, *TeacherId) (*Teacher, error)
	UpdateTeacher(context.Context, *Teacher) (*Teacher, error)
	mustEmbedUnimplementedTeachersServiceServer()
}

// UnimplementedTeachersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeachersServiceServer struct {
}

func (UnimplementedTeachersServiceServer) GetTeacher(context.Context, *TeacherId) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacher not implemented")
}
func (UnimplementedTeachersServiceServer) GetTeachers(context.Context, *GetTeachersRequest) (*TeacherList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeachers not implemented")
}
func (UnimplementedTeachersServiceServer) CreateTeacher(context.Context, *Teacher) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeacher not implemented")
}
func (UnimplementedTeachersServiceServer) DeleteTeacher(context.Context, *TeacherId) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeacher not implemented")
}
func (UnimplementedTeachersServiceServer) UpdateTeacher(context.Context, *Teacher) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeacher not implemented")
}
func (UnimplementedTeachersServiceServer) mustEmbedUnimplementedTeachersServiceServer() {}

// UnsafeTeachersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeachersServiceServer will
// result in compilation errors.
type UnsafeTeachersServiceServer interface {
	mustEmbedUnimplementedTeachersServiceServer()
}

func RegisterTeachersServiceServer(s grpc.ServiceRegistrar, srv TeachersServiceServer) {
	s.RegisterService(&TeachersService_ServiceDesc, srv)
}

func _TeachersService_GetTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersServiceServer).GetTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersService_GetTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersServiceServer).GetTeacher(ctx, req.(*TeacherId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersService_GetTeachers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeachersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersServiceServer).GetTeachers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersService_GetTeachers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersServiceServer).GetTeachers(ctx, req.(*GetTeachersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersService_CreateTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersServiceServer).CreateTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersService_CreateTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersServiceServer).CreateTeacher(ctx, req.(*Teacher))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersService_DeleteTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersServiceServer).DeleteTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersService_DeleteTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersServiceServer).DeleteTeacher(ctx, req.(*TeacherId))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersService_UpdateTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Teacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersServiceServer).UpdateTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersService_UpdateTeacher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersServiceServer).UpdateTeacher(ctx, req.(*Teacher))
	}
	return interceptor(ctx, in, info, handler)
}

// TeachersService_ServiceDesc is the grpc.ServiceDesc for TeachersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeachersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.TeachersService",
	HandlerType: (*TeachersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeacher",
			Handler:    _TeachersService_GetTeacher_Handler,
		},
		{
			MethodName: "GetTeachers",
			Handler:    _TeachersService_GetTeachers_Handler,
		},
		{
			MethodName: "CreateTeacher",
			Handler:    _TeachersService_CreateTeacher_Handler,
		},
		{
			MethodName: "DeleteTeacher",
			Handler:    _TeachersService_DeleteTeacher_Handler,
		},
		{
			MethodName: "UpdateTeacher",
			Handler:    _TeachersService_UpdateTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teachers_service.proto",
}
