package service

import (
	"context"
	"crmservice/config"
	ct "crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerService struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func NewManagerService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *ManagerService {

	return &ManagerService{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *ManagerService) CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (resp *ct.ManagerStudent, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Manager_student().CreateStudent(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err
	}

	return resp, nil
}

func (s *ManagerService) GetByIDStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Manager_student().GetStudentByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err

	}

	return resp, nil

}

func (s *ManagerService) GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---GetbyidManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manager_student().GetStudentByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidManagerService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err

	}

	return resp, nil

}

func (s *ManagerService) GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (resp *ct.GetListManagerStudentResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manager_student().GetAllStudents(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerStudentResponse{}, err
	}

	return resp, nil
}

func (s *ManagerService) UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---UpdateStudentManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manager_student().UpdateStudent(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentManagerService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err
	}

	return resp, nil
}

func (s *ManagerService) DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudentPrimaryKey, err error) {

	s.log.Info("---DeleteStudentManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manager_student().DeleteStudent(ctx, req)
	if err != nil {
		s.log.Error("---DeleteStudentManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}
