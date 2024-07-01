package storage

import (
	"context"
	ct "superadmin/genproto/superadmin_service"
)

type StorageI interface {
	CloseDB()
	Superadmin_student() Superadmin_crmRepoRepoI
}

type Superadmin_crmRepoRepoI interface {
	CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (resp *ct.ManagerStudent, err error)
	GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error)
	GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (resp *ct.GetListManagerStudentResponse, err error)
	UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (resp *ct.ManagerStudent, err error)
	DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudentPrimaryKey, err error)

	CreateTeacher(ctx context.Context, req *ct.CreateManagerTeacher) (resp *ct.ManagerTeacher, err error)
	GetTeacherByID(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacher, err error)
	GetAllTeachers(ctx context.Context, req *ct.GetListManagerTeacherRequest) (resp *ct.GetListManagerTeacherResponse, err error)
	UpdateTeacher(ctx context.Context, req *ct.UpdateManagerTeacher) (resp *ct.ManagerTeacher, err error)
	DeleteTeacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacherPrimaryKey, err error)

	CreateSupportTeacher(ctx context.Context, req *ct.CreateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error)
	GetSupportTeacherByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacher, err error)
	GetAllSupportTeachers(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (resp *ct.GetListManagerSupportTeacherResponse, err error)
	UpdateSupportTeacher(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error)
	DeleteSupportTeacher(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacherPrimaryKey, err error)

	CreateManager(ctx context.Context, req *ct.CreateManagerRequest) (resp *ct.Manager, err error)
	GetManagerByID(ctx context.Context, req *ct.ManagerPrimaryKey) (resp *ct.Manager, err error)
	GetAllManagers(ctx context.Context, req *ct.GetAllManagersRequest) (resp *ct.GetAllManagersResponse, err error)
	UpdateManager(ctx context.Context, req *ct.UpdateManagerRequest) (resp *ct.Manager, err error)
	DeleteManager(ctx context.Context, req *ct.ManagerPrimaryKey) (resp *ct.ManagerPrimaryKey, err error)

	CreateAdministrator(ctx context.Context, req *ct.CreateManagerAdministrator) (resp *ct.ManagerAdministrator, err error)
	GetAdministratorByID(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (resp *ct.ManagerAdministrator, err error)
	GetAllAdministrators(ctx context.Context, req *ct.GetListManagerAdministratorRequest) (resp *ct.GetListManagerAdministratorResponse, err error)
	//	UpdateAdministrator(ctx context.Context, req *ct.UpdateManagerAdministrator) (resp *ct.ManagerAdministrator, err error)
	DeleteAdministrator(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (resp *ct.ManagerAdministratorPrimaryKey, err error)

	ExportToCSV(ctx context.Context, req *ct.TableName) (resp *ct.Empty, err error)
}
