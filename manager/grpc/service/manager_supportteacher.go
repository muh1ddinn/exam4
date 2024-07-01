package service

import (
	"context"
	"crmservice/config"
	ct "crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerServicesupportteacher struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func NewManagersupportteacherService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *ManagerServicesupportteacher {

	return &ManagerServicesupportteacher{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *ManagerServicesupportteacher) Create(ctx context.Context, req *ct.CreateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Managersupportteacher().Createsupportteacher(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err
	}

	return resp, nil
}

func (s *ManagerServicesupportteacher) GetByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacher, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Managersupportteacher().GetsupportteacherByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err

	}

	return resp, nil

}

func (s *ManagerServicesupportteacher) GetAll(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (resp *ct.GetListManagerSupportTeacherResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managersupportteacher().GetAllsupportteacher(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *ManagerServicesupportteacher) Update(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error) {
	s.log.Info("---UpdatesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managersupportteacher().Updatesupportteacher(ctx, req)
	if err != nil {
		s.log.Error("---UpdatesupportteacherManagerService--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err
	}

	return resp, nil
}

func (s *ManagerServicesupportteacher) Delete(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacherPrimaryKey, err error) {

	s.log.Info("---DeletesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Managersupportteacher().Deletesupportteacher(ctx, req)
	if err != nil {
		s.log.Error("---DeletesupportteacherManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}
