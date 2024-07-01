package service

import (
	"context"
	"crmservice/config"
	ct "crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type ManagerServiceadministration struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func NewManageradministrationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *ManagerServiceadministration {

	return &ManagerServiceadministration{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *ManagerServiceadministration) Create(ctx context.Context, req *ct.CreateManagerAdministration) (resp *ct.ManagerAdministration, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Manageradministratio().Createadministration(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerAdministration{}, err
	}

	return resp, nil
}

func (s *ManagerServiceadministration) GetByID(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (resp *ct.ManagerAdministration, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Manageradministratio().GetadministrationrByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerAdministration{}, err

	}

	return resp, nil

}

func (s *ManagerServiceadministration) GetAll(ctx context.Context, req *ct.GetListManagerAdministrationRequest) (resp *ct.GetListManagerAdministrationResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manageradministratio().GetAlladministration(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerAdministrationResponse{}, err
	}

	return resp, nil
}

func (s *ManagerServiceadministration) Update(ctx context.Context, req *ct.UpdateManagerAdministration) (resp *ct.ManagerAdministration, err error) {
	s.log.Info("---UpdateteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manageradministratio().Updateadministration(ctx, req)
	if err != nil {
		s.log.Error("---UpdateteacherManagerService--->>>", logger.Error(err))
		return &ct.ManagerAdministration{}, err
	}

	return resp, nil
}

func (s *ManagerServiceadministration) Delete(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (resp *ct.ManagerAdministrationPrimaryKey, err error) {

	s.log.Info("---DeleteteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Manageradministratio().Deleteadministration(ctx, req)
	if err != nil {
		s.log.Error("---DeleteteacherManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}
