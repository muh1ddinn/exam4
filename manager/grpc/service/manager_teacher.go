package service

import (
	"context"
	"crmservice/config"
	ct "crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerServiceteacher struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func NewManagerteacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *ManagerServiceteacher {

	return &ManagerServiceteacher{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *ManagerServiceteacher) Createteacher(ctx context.Context, req *ct.CreateManagerTeacher) (resp *ct.ManagerTeacher, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Managerteacher().Createteacher(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err
	}

	return resp, nil
}

func (s *ManagerServiceteacher) GetteacherByI(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacher, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Managerteacher().GetteacherByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err

	}

	return resp, nil

}

func (s *ManagerServiceteacher) GetAllteacher(ctx context.Context, req *ct.GetListManagerTeacherRequest) (resp *ct.GetListManagerTeacherResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managerteacher().GetAllteacher(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerTeacherResponse{}, err
	}

	return resp, nil
}

func (s *ManagerServiceteacher) Updateteacher(ctx context.Context, req *ct.UpdateManagerTeacher) (resp *ct.ManagerTeacher, err error) {
	s.log.Info("---UpdateteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managerteacher().Updateteacher(ctx, req)
	if err != nil {
		s.log.Error("---UpdateteacherManagerService--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err
	}

	return resp, nil
}

func (s *ManagerServiceteacher) Deleteteacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacherPrimaryKey, err error) {

	s.log.Info("---DeleteteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managerteacher().Deleteteacher(ctx, req)
	if err != nil {
		s.log.Error("---DeleteteacherManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}
