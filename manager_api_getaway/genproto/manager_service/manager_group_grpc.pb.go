// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: manager_group.proto

package manager_service

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

// ManagerGroupServiceClient is the client API for ManagerGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerGroupServiceClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*GroupID, error)
	GetGroupByID(ctx context.Context, in *GroupID, opts ...grpc.CallOption) (*Group, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskID, error)
	GetTaskByID(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*Task, error)
	GetAllTasks(ctx context.Context, in *GetAllTasksRequest, opts ...grpc.CallOption) (*GetAllTasksResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error)
	CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*TypeID, error)
	GetTypeByID(ctx context.Context, in *TypeID, opts ...grpc.CallOption) (*Type, error)
	CreateSchedules(ctx context.Context, in *CreateSchedulesRequest, opts ...grpc.CallOption) (*SchedulesID, error)
	GetSchedulesByID(ctx context.Context, in *SchedulesID, opts ...grpc.CallOption) (*Schedules, error)
	CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*JournalID, error)
	GetJournalByID(ctx context.Context, in *JournalID, opts ...grpc.CallOption) (*Journal, error)
	CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*LessonID, error)
	GetLessonByID(ctx context.Context, in *LessonID, opts ...grpc.CallOption) (*Lesson, error)
	GetTeacherInfo(ctx context.Context, in *TeacherGetTeacherInfoRequest, opts ...grpc.CallOption) (*TeacherGetTeacherInfoResponse, error)
	GetSupporTeacherInfo(ctx context.Context, in *TeacherGetTeacherInfoRequest, opts ...grpc.CallOption) (*TeacherGetSupportInfoResponse, error)
}

type managerGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerGroupServiceClient(cc grpc.ClientConnInterface) ManagerGroupServiceClient {
	return &managerGroupServiceClient{cc}
}

func (c *managerGroupServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*GroupID, error) {
	out := new(GroupID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetGroupByID(ctx context.Context, in *GroupID, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetGroupByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*TaskID, error) {
	out := new(TaskID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetTaskByID(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetTaskByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetAllTasks(ctx context.Context, in *GetAllTasksRequest, opts ...grpc.CallOption) (*GetAllTasksResponse, error) {
	out := new(GetAllTasksResponse)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetAllTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) CreateType(ctx context.Context, in *CreateTypeRequest, opts ...grpc.CallOption) (*TypeID, error) {
	out := new(TypeID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetTypeByID(ctx context.Context, in *TypeID, opts ...grpc.CallOption) (*Type, error) {
	out := new(Type)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetTypeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) CreateSchedules(ctx context.Context, in *CreateSchedulesRequest, opts ...grpc.CallOption) (*SchedulesID, error) {
	out := new(SchedulesID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateSchedules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetSchedulesByID(ctx context.Context, in *SchedulesID, opts ...grpc.CallOption) (*Schedules, error) {
	out := new(Schedules)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetSchedulesByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) CreateJournal(ctx context.Context, in *CreateJournalRequest, opts ...grpc.CallOption) (*JournalID, error) {
	out := new(JournalID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateJournal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetJournalByID(ctx context.Context, in *JournalID, opts ...grpc.CallOption) (*Journal, error) {
	out := new(Journal)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetJournalByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) CreateLesson(ctx context.Context, in *CreateLessonRequest, opts ...grpc.CallOption) (*LessonID, error) {
	out := new(LessonID)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/CreateLesson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetLessonByID(ctx context.Context, in *LessonID, opts ...grpc.CallOption) (*Lesson, error) {
	out := new(Lesson)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetLessonByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetTeacherInfo(ctx context.Context, in *TeacherGetTeacherInfoRequest, opts ...grpc.CallOption) (*TeacherGetTeacherInfoResponse, error) {
	out := new(TeacherGetTeacherInfoResponse)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetTeacherInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerGroupServiceClient) GetSupporTeacherInfo(ctx context.Context, in *TeacherGetTeacherInfoRequest, opts ...grpc.CallOption) (*TeacherGetSupportInfoResponse, error) {
	out := new(TeacherGetSupportInfoResponse)
	err := c.cc.Invoke(ctx, "/manager_service.ManagerGroupService/GetSupporTeacherInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerGroupServiceServer is the server API for ManagerGroupService service.
// All implementations should embed UnimplementedManagerGroupServiceServer
// for forward compatibility
type ManagerGroupServiceServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*GroupID, error)
	GetGroupByID(context.Context, *GroupID) (*Group, error)
	CreateTask(context.Context, *CreateTaskRequest) (*TaskID, error)
	GetTaskByID(context.Context, *TaskID) (*Task, error)
	GetAllTasks(context.Context, *GetAllTasksRequest) (*GetAllTasksResponse, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error)
	CreateType(context.Context, *CreateTypeRequest) (*TypeID, error)
	GetTypeByID(context.Context, *TypeID) (*Type, error)
	CreateSchedules(context.Context, *CreateSchedulesRequest) (*SchedulesID, error)
	GetSchedulesByID(context.Context, *SchedulesID) (*Schedules, error)
	CreateJournal(context.Context, *CreateJournalRequest) (*JournalID, error)
	GetJournalByID(context.Context, *JournalID) (*Journal, error)
	CreateLesson(context.Context, *CreateLessonRequest) (*LessonID, error)
	GetLessonByID(context.Context, *LessonID) (*Lesson, error)
	GetTeacherInfo(context.Context, *TeacherGetTeacherInfoRequest) (*TeacherGetTeacherInfoResponse, error)
	GetSupporTeacherInfo(context.Context, *TeacherGetTeacherInfoRequest) (*TeacherGetSupportInfoResponse, error)
}

// UnimplementedManagerGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedManagerGroupServiceServer struct {
}

func (UnimplementedManagerGroupServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*GroupID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetGroupByID(context.Context, *GroupID) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) CreateTask(context.Context, *CreateTaskRequest) (*TaskID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetTaskByID(context.Context, *TaskID) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetAllTasks(context.Context, *GetAllTasksRequest) (*GetAllTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTasks not implemented")
}
func (UnimplementedManagerGroupServiceServer) UpdateTask(context.Context, *UpdateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedManagerGroupServiceServer) CreateType(context.Context, *CreateTypeRequest) (*TypeID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateType not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetTypeByID(context.Context, *TypeID) (*Type, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypeByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) CreateSchedules(context.Context, *CreateSchedulesRequest) (*SchedulesID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedules not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetSchedulesByID(context.Context, *SchedulesID) (*Schedules, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedulesByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) CreateJournal(context.Context, *CreateJournalRequest) (*JournalID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJournal not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetJournalByID(context.Context, *JournalID) (*Journal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJournalByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) CreateLesson(context.Context, *CreateLessonRequest) (*LessonID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLesson not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetLessonByID(context.Context, *LessonID) (*Lesson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonByID not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetTeacherInfo(context.Context, *TeacherGetTeacherInfoRequest) (*TeacherGetTeacherInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacherInfo not implemented")
}
func (UnimplementedManagerGroupServiceServer) GetSupporTeacherInfo(context.Context, *TeacherGetTeacherInfoRequest) (*TeacherGetSupportInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupporTeacherInfo not implemented")
}

// UnsafeManagerGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerGroupServiceServer will
// result in compilation errors.
type UnsafeManagerGroupServiceServer interface {
	mustEmbedUnimplementedManagerGroupServiceServer()
}

func RegisterManagerGroupServiceServer(s grpc.ServiceRegistrar, srv ManagerGroupServiceServer) {
	s.RegisterService(&ManagerGroupService_ServiceDesc, srv)
}

func _ManagerGroupService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetGroupByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetGroupByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetGroupByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetGroupByID(ctx, req.(*GroupID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetTaskByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetTaskByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetTaskByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetTaskByID(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetAllTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetAllTasks(ctx, req.(*GetAllTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_CreateType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateType(ctx, req.(*CreateTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetTypeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetTypeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetTypeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetTypeByID(ctx, req.(*TypeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_CreateSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateSchedules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateSchedules(ctx, req.(*CreateSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetSchedulesByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulesID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetSchedulesByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetSchedulesByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetSchedulesByID(ctx, req.(*SchedulesID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_CreateJournal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJournalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateJournal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateJournal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateJournal(ctx, req.(*CreateJournalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetJournalByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JournalID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetJournalByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetJournalByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetJournalByID(ctx, req.(*JournalID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_CreateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLessonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).CreateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/CreateLesson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).CreateLesson(ctx, req.(*CreateLessonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetLessonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetLessonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetLessonByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetLessonByID(ctx, req.(*LessonID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetTeacherInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherGetTeacherInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetTeacherInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetTeacherInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetTeacherInfo(ctx, req.(*TeacherGetTeacherInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerGroupService_GetSupporTeacherInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeacherGetTeacherInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerGroupServiceServer).GetSupporTeacherInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager_service.ManagerGroupService/GetSupporTeacherInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerGroupServiceServer).GetSupporTeacherInfo(ctx, req.(*TeacherGetTeacherInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerGroupService_ServiceDesc is the grpc.ServiceDesc for ManagerGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager_service.ManagerGroupService",
	HandlerType: (*ManagerGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _ManagerGroupService_CreateGroup_Handler,
		},
		{
			MethodName: "GetGroupByID",
			Handler:    _ManagerGroupService_GetGroupByID_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _ManagerGroupService_CreateTask_Handler,
		},
		{
			MethodName: "GetTaskByID",
			Handler:    _ManagerGroupService_GetTaskByID_Handler,
		},
		{
			MethodName: "GetAllTasks",
			Handler:    _ManagerGroupService_GetAllTasks_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _ManagerGroupService_UpdateTask_Handler,
		},
		{
			MethodName: "CreateType",
			Handler:    _ManagerGroupService_CreateType_Handler,
		},
		{
			MethodName: "GetTypeByID",
			Handler:    _ManagerGroupService_GetTypeByID_Handler,
		},
		{
			MethodName: "CreateSchedules",
			Handler:    _ManagerGroupService_CreateSchedules_Handler,
		},
		{
			MethodName: "GetSchedulesByID",
			Handler:    _ManagerGroupService_GetSchedulesByID_Handler,
		},
		{
			MethodName: "CreateJournal",
			Handler:    _ManagerGroupService_CreateJournal_Handler,
		},
		{
			MethodName: "GetJournalByID",
			Handler:    _ManagerGroupService_GetJournalByID_Handler,
		},
		{
			MethodName: "CreateLesson",
			Handler:    _ManagerGroupService_CreateLesson_Handler,
		},
		{
			MethodName: "GetLessonByID",
			Handler:    _ManagerGroupService_GetLessonByID_Handler,
		},
		{
			MethodName: "GetTeacherInfo",
			Handler:    _ManagerGroupService_GetTeacherInfo_Handler,
		},
		{
			MethodName: "GetSupporTeacherInfo",
			Handler:    _ManagerGroupService_GetSupporTeacherInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager_group.proto",
}
