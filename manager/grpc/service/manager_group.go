package service

import (
	"context"
	"crmservice/config"
	ct "crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerServicegroup struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func (f *ManagerServicegroup) CreateSchedules(ctx context.Context, req *ct.CreateSchedulesRequest) (resp *ct.SchedulesID, err error) {

	f.log.Info("---CreateSchedules--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateSchedules(ctx, req)
	if err != nil {
		f.log.Error("---CreateSchedules--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) GetSchedulesByID(ctx context.Context, req *ct.SchedulesID) (resp *ct.Schedules, err error) {
	s.log.Info("---GetSchedulesByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetSchedulesByID(ctx, req)

	if err != nil {
		s.log.Error("---GetSchedulesByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

// DeleteLesson implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) DeleteLesson(context.Context, *ct.LessonID) (*ct.LessonID, error) {
	panic("unimplemented")
}

// GetAllLesson implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) GetAllLesson(context.Context, *ct.GetAllLessonRequest) (*ct.GetAllLessonResponse, error) {
	panic("unimplemented")
}

// UpdateLessonl implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) UpdateLessonl(context.Context, *ct.UpdateLessonRequest) (*ct.Lesson, error) {
	panic("unimplemented")
}

func (f *ManagerServicegroup) CreateGroup(ctx context.Context, req *ct.CreateGroupRequest) (resp *ct.GroupID, err error) {

	f.log.Info("---CreateGroup--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateGroup(ctx, req)
	if err != nil {
		f.log.Error("---CreateGroup--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) GetGroupByID(ctx context.Context, req *ct.GroupID) (resp *ct.Group, err error) {
	s.log.Info("---GetGroupByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetGroupByID(ctx, req)

	if err != nil {
		s.log.Error("---GetGroupByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

func (f *ManagerServicegroup) CreateTask(ctx context.Context, req *ct.CreateTaskRequest) (resp *ct.TaskID, err error) {

	f.log.Info("---CreateTask--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateTask(ctx, req)
	if err != nil {
		f.log.Error("---CreateTask--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}
func (f *ManagerServicegroup) CreateLesson(ctx context.Context, req *ct.CreateLessonRequest) (resp *ct.LessonID, err error) {

	f.log.Info("---CreateLesson--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateLesson(ctx, req)
	if err != nil {
		f.log.Error("---CreateLesson--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) GetLessonByID(ctx context.Context, req *ct.LessonID) (resp *ct.Lesson, err error) {
	s.log.Info("---GetLessonByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetLessonByID(ctx, req)

	if err != nil {
		s.log.Error("---GetLessonByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

// CreateType implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) CreateType(ctx context.Context, req *ct.CreateTypeRequest) (resp *ct.TypeID, err error) {

	f.log.Info("---CreateType--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateType(ctx, req)
	if err != nil {
		f.log.Error("---CreateType--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

// DeleteGroup implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) DeleteGroup(context.Context, *ct.GroupID) (*ct.GroupID, error) {
	panic("unimplemented")
}

// DeleteJournal implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) DeleteJournal(context.Context, *ct.JournalID) (*ct.JournalID, error) {
	panic("unimplemented")
}

// DeleteType implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) DeleteType(context.Context, *ct.TypeID) (*ct.TypeID, error) {
	panic("unimplemented")
}

// GetAllGroups implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) GetAllGroups(context.Context, *ct.GetAllGroupsRequest) (*ct.GetAllGroupsResponse, error) {
	panic("unimplemented")
}

// GetAllJournals implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) GetAllJournals(context.Context, *ct.GetAllJournalsRequest) (*ct.GetAllJournalsResponse, error) {
	panic("unimplemented")
}

// GetAllTypes implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) GetAllTypes(context.Context, *ct.GetAllTypesRequest) (*ct.GetAllTypesResponse, error) {
	panic("unimplemented")
}

func (s *ManagerServicegroup) GetTaskByID(ctx context.Context, req *ct.TaskID) (resp *ct.Task, err error) {
	s.log.Info("---GetTaskByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetTaskByID(ctx, req)

	if err != nil {
		s.log.Error("---GetTaskByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

func (s *ManagerServicegroup) GetTypeByID(ctx context.Context, req *ct.TypeID) (resp *ct.Type, err error) {
	s.log.Info("---GetTypeByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetTypeByID(ctx, req)

	if err != nil {
		s.log.Error("---GetTypeByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

// UpdateGroup implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) UpdateGroup(context.Context, *ct.UpdateGroupRequest) (*ct.Group, error) {
	panic("unimplemented")
}

// UpdateJournal implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) UpdateJournal(context.Context, *ct.UpdateJournalRequest) (*ct.Journal, error) {
	panic("unimplemented")
}

// UpdateType implements manager_service.ManagerGroupServiceServer.
func (f *ManagerServicegroup) UpdateType(context.Context, *ct.UpdateTypeRequest) (*ct.Type, error) {
	panic("unimplemented")
}

func NewManagerServicegroup(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *ManagerServicegroup {

	return &ManagerServicegroup{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *ManagerServicegroup) CreateJournal(ctx context.Context, req *ct.CreateJournalRequest) (resp *ct.JournalID, err error) {

	f.log.Info("---CreateJournal--->>>", logger.Any("req", req))

	resp, err = f.strg.Managergroup().CreateJournal(ctx, req)
	if err != nil {
		f.log.Error("---CreateJournal--->>>", logger.Error(err))
		return resp, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) GetJournalByID(ctx context.Context, req *ct.JournalID) (resp *ct.Journal, err error) {
	s.log.Info("---GetJournalByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetJournalByID(ctx, req)

	if err != nil {
		s.log.Error("---GetJournalByID--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

func (s *ManagerServicegroup) GetAllTasks(ctx context.Context, req *ct.GetAllTasksRequest) (resp *ct.GetAllTasksResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetAllTasks(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetAllTasksResponse{}, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) UpdateTask(ctx context.Context, req *ct.UpdateTaskRequest) (resp *ct.Task, err error) {
	s.log.Info("---UpdateTask--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().UpdateTask(ctx, req)
	if err != nil {
		s.log.Error("---UpdateTask--->>>", logger.Error(err))
		return &ct.Task{}, err
	}

	return resp, nil
}

func (s *ManagerServicegroup) GetTeacherInfo(ctx context.Context, req *ct.TeacherGetTeacherInfoRequest) (resp *ct.TeacherGetTeacherInfoResponse, err error) {
	s.log.Info("---GetTeacherInfo--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetTeacherInfo(ctx, req)

	if err != nil {
		s.log.Error("---GetTeacherInfo--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}

func (s *ManagerServicegroup) GetSupporTeacherInfo(ctx context.Context, req *ct.TeacherGetTeacherInfoRequest) (resp *ct.TeacherGetSupportInfoResponse, err error) {
	s.log.Info("---GetTeacherInfo--->>>", logger.Any("req", req))

	resp, err = s.strg.Managergroup().GetSupporTeacherInfo(ctx, req)

	if err != nil {
		s.log.Error("---GetTeacherInfo--->>>", logger.Error(err))
		return resp, err

	}

	return resp, nil

}
