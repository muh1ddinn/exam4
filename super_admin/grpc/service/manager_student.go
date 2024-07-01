package service

import (
	"context"
	"superadmin/config"
	ct "superadmin/genproto/superadmin_service"
	"superadmin/grpc/client"
	"superadmin/storage"

	"github.com/saidamir98/udevs_pkg/logger"
)

type SuperadminrService struct {
	cfg     config.Config
	log     logger.LoggerI
	strg    storage.StorageI
	service client.ServciceManagerI
}

func NewSuperadminrService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServciceManagerI) *SuperadminrService {

	return &SuperadminrService{
		cfg:     cfg,
		log:     log,
		strg:    strg,
		service: srvs,
	}

}

func (f *SuperadminrService) CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (resp *ct.ManagerStudent, err error) {

	f.log.Info("---CreateSuperadminrService--->>>", logger.Any("req", req))

	resp, err = f.strg.Superadmin_student().CreateStudent(ctx, req)
	if err != nil {
		f.log.Error("---CreateSuperadminrService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetByIDStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetStudentByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err

	}

	return resp, nil

}

func (s *SuperadminrService) GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---GetbyidSuperadminrService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetStudentByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidSuperadminrService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err

	}

	return resp, nil

}

func (s *SuperadminrService) GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (resp *ct.GetListManagerStudentResponse, err error) {
	s.log.Info("---GetAllSuperadminrService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetAllStudents(ctx, req)
	if err != nil {
		s.log.Error("---GetAllSuperadminrService--->>>", logger.Error(err))
		return &ct.GetListManagerStudentResponse{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (resp *ct.ManagerStudent, err error) {
	s.log.Info("---UpdateStudentSuperadminrService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().UpdateStudent(ctx, req)
	if err != nil {
		s.log.Error("---UpdateStudentSuperadminrService--->>>", logger.Error(err))
		return &ct.ManagerStudent{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudentPrimaryKey, err error) {

	s.log.Info("---DeleteStudentSuperadminrService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().DeleteStudent(ctx, req)
	if err != nil {
		s.log.Error("---DeleteStudentSuperadminrService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}

func (f *SuperadminrService) CreateSupportTeacher(ctx context.Context, req *ct.CreateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Superadmin_student().CreateSupportTeacher(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetSupportTeacherByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacher, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetSupportTeacherByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err

	}

	return resp, nil

}

func (s *SuperadminrService) GetAllSupportTeachers(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (resp *ct.GetListManagerSupportTeacherResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetAllSupportTeachers(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerSupportTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) UpdateSupportTeacher(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error) {
	s.log.Info("---UpdatesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().UpdateSupportTeacher(ctx, req)
	if err != nil {
		s.log.Error("---UpdatesupportteacherManagerService--->>>", logger.Error(err))
		return &ct.ManagerSupportTeacher{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) DeleteSupportTeacher(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacherPrimaryKey, err error) {

	s.log.Info("---DeletesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().DeleteSupportTeacher(ctx, req)
	if err != nil {
		s.log.Error("---DeletesupportteacherManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}

func (f *SuperadminrService) CreateTeacher(ctx context.Context, req *ct.CreateManagerTeacher) (resp *ct.ManagerTeacher, err error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err = f.strg.Superadmin_student().CreateTeacher(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetTeacherByID(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacher, err error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetTeacherByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err

	}

	return resp, nil

}

func (s *SuperadminrService) GetAllTeachers(ctx context.Context, req *ct.GetListManagerTeacherRequest) (resp *ct.GetListManagerTeacherResponse, err error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetAllTeachers(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerTeacherResponse{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) UpdateTeacher(ctx context.Context, req *ct.UpdateManagerTeacher) (resp *ct.ManagerTeacher, err error) {
	s.log.Info("---UpdatesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().UpdateTeacher(ctx, req)
	if err != nil {
		s.log.Error("---UpdatesupportteacherManagerService--->>>", logger.Error(err))
		return &ct.ManagerTeacher{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) DeleteTeacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacherPrimaryKey, err error) {

	s.log.Info("---DeletesupportteacherManagerService--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().DeleteTeacher(ctx, req)
	if err != nil {
		s.log.Error("---DeletesupportteacherManagerService--->>>", logger.Error(err))
		return resp, nil
	}

	return resp, nil
}

func (s *SuperadminrService) CreateManager(ctx context.Context, req *ct.CreateManagerRequest) (resp *ct.Manager, err error) {
	s.log.Info("---CreateManager--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().CreateManager(ctx, req)
	if err != nil {
		s.log.Error("---CreateManager--->>>", logger.Error(err))
		return &ct.Manager{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetManagerByID(ctx context.Context, req *ct.ManagerPrimaryKey) (resp *ct.Manager, err error) {
	s.log.Info("---GetManagerByID--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetManagerByID(ctx, req)
	if err != nil {
		s.log.Error("---GetManagerByID--->>>", logger.Error(err))
		return &ct.Manager{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetAllManagers(ctx context.Context, req *ct.GetAllManagersRequest) (resp *ct.GetAllManagersResponse, err error) {
	s.log.Info("---GetAllManagers--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().GetAllManagers(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagers--->>>", logger.Error(err))
		return &ct.GetAllManagersResponse{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) UpdateManager(ctx context.Context, req *ct.UpdateManagerRequest) (resp *ct.Manager, err error) {
	s.log.Info("---UpdateManager--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().UpdateManager(ctx, req)
	if err != nil {
		s.log.Error("---UpdateManager--->>>", logger.Error(err))
		return &ct.Manager{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) DeleteManager(ctx context.Context, req *ct.ManagerPrimaryKey) (resp *ct.ManagerPrimaryKey, err error) {
	s.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	resp, err = s.strg.Superadmin_student().DeleteManager(ctx, req)
	if err != nil {
		s.log.Error("---DeleteManager--->>>", logger.Error(err))
		return &ct.ManagerPrimaryKey{}, nil
	}

	return resp, nil
}

func (f *SuperadminrService) CreateAdministrator(ctx context.Context, req *ct.CreateManagerAdministrator) (*ct.ManagerAdministrator, error) {

	f.log.Info("---CreateManagerService--->>>", logger.Any("req", req))

	resp, err := f.strg.Superadmin_student().CreateAdministrator(ctx, req)
	if err != nil {
		f.log.Error("---CreateManagerService--->>>", logger.Error(err))
		return &ct.ManagerAdministrator{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) GetAdministratorByID(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (*ct.ManagerAdministrator, error) {
	s.log.Info("---GetbyidProductOrder--->>>", logger.Any("req", req))

	resp, err := s.strg.Superadmin_student().GetAdministratorByID(ctx, req)

	if err != nil {
		s.log.Error("---GetbyidProductOrder--->>>", logger.Error(err))
		return &ct.ManagerAdministrator{}, err

	}

	return resp, nil

}

func (s *SuperadminrService) GetAllAdministrators(ctx context.Context, req *ct.GetListManagerAdministratorRequest) (*ct.GetListManagerAdministratorResponse, error) {
	s.log.Info("---GetAllManagerService--->>>", logger.Any("req", req))

	resp, err := s.strg.Superadmin_student().GetAllAdministrators(ctx, req)
	if err != nil {
		s.log.Error("---GetAllManagerService--->>>", logger.Error(err))
		return &ct.GetListManagerAdministratorResponse{}, err
	}

	return resp, nil
}

func (s *SuperadminrService) DeleteAdministrator(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (*ct.ManagerAdministratorPrimaryKey, error) {
	s.log.Info("---DeleteManager--->>>", logger.Any("req", req))

	resp, err := s.strg.Superadmin_student().DeleteAdministrator(ctx, req)
	if err != nil {
		s.log.Error("---DeleteManager--->>>", logger.Error(err))
		return &ct.ManagerAdministratorPrimaryKey{}, nil
	}

	return resp, nil
}
func (s SuperadminrService) ExportToCSV(ctx context.Context, req *ct.TableName) (*ct.Empty, error) {
	resp, err := s.strg.Superadmin_student().ExportToCSV(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
