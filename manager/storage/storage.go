package storage

import (
	"context"
	ct "crmservice/genproto/manager_service"
)

type StorageI interface {
	CloseDB()
	Manager_student() Manager_crmRepoRepoI
	Managerteacher() Managerteacher_RepoRepoI
	Managersupportteacher() Managersupportteacher_RepoRepoI
	Manageradministratio() ManageradministrationRepo_RepoRepoI
	Managergroup() ManagerGroupService_RepoRepoI
}

type Manager_crmRepoRepoI interface {
	CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (resp *ct.ManagerStudent, err error)
	GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudent, err error)
	GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (resp *ct.GetListManagerStudentResponse, err error)
	UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (resp *ct.ManagerStudent, err error)
	DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (resp *ct.ManagerStudentPrimaryKey, err error)
}

type Managerteacher_RepoRepoI interface {
	Createteacher(ctx context.Context, req *ct.CreateManagerTeacher) (resp *ct.ManagerTeacher, err error)
	GetteacherByID(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacher, err error)
	GetAllteacher(ctx context.Context, req *ct.GetListManagerTeacherRequest) (resp *ct.GetListManagerTeacherResponse, err error)
	Updateteacher(ctx context.Context, req *ct.UpdateManagerTeacher) (resp *ct.ManagerTeacher, err error)
	Deleteteacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (resp *ct.ManagerTeacherPrimaryKey, err error)
}

type Managersupportteacher_RepoRepoI interface {
	Createsupportteacher(ctx context.Context, req *ct.CreateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error)
	GetsupportteacherByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacher, err error)
	GetAllsupportteacher(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (resp *ct.GetListManagerSupportTeacherResponse, err error)
	Updatesupportteacher(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (resp *ct.ManagerSupportTeacher, err error)
	Deletesupportteacher(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (resp *ct.ManagerSupportTeacherPrimaryKey, err error)
}

type ManageradministrationRepo_RepoRepoI interface {
	Createadministration(ctx context.Context, req *ct.CreateManagerAdministration) (resp *ct.ManagerAdministration, err error)
	GetadministrationrByID(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (resp *ct.ManagerAdministration, err error)
	GetAlladministration(ctx context.Context, req *ct.GetListManagerAdministrationRequest) (resp *ct.GetListManagerAdministrationResponse, err error)
	Updateadministration(ctx context.Context, req *ct.UpdateManagerAdministration) (resp *ct.ManagerAdministration, err error)
	Deleteadministration(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (resp *ct.ManagerAdministrationPrimaryKey, err error)
}

type ManagerGroupService_RepoRepoI interface {
	CreateGroup(ctx context.Context, req *ct.CreateGroupRequest) (resp *ct.GroupID, err error)
	GetGroupByID(ctx context.Context, req *ct.GroupID) (resp *ct.Group, err error)

	//GetAlladministration(ctx context.Context, req *ct.GetListManagerAdministrationRequest) (resp *ct.GetListManagerAdministrationResponse, err error)
	//Updateadministration(ctx context.Context, req *ct.UpdateManagerAdministration) (resp *ct.ManagerAdministration, err error)
	//Deleteadministration(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (resp *ct.ManagerAdministrationPrimaryKey, err error)

	CreateTask(ctx context.Context, req *ct.CreateTaskRequest) (resp *ct.TaskID, err error)
	GetTaskByID(ctx context.Context, req *ct.TaskID) (resp *ct.Task, err error)

	CreateLesson(ctx context.Context, req *ct.CreateLessonRequest) (resp *ct.LessonID, err error)
	GetLessonByID(ctx context.Context, req *ct.LessonID) (resp *ct.Lesson, err error)

	CreateType(ctx context.Context, req *ct.CreateTypeRequest) (resp *ct.TypeID, err error)
	GetTypeByID(ctx context.Context, req *ct.TypeID) (resp *ct.Type, err error)
	GetAllTasks(ctx context.Context, req *ct.GetAllTasksRequest) (resp *ct.GetAllTasksResponse, err error)
	UpdateTask(ctx context.Context, req *ct.UpdateTaskRequest) (resp *ct.Task, err error)
	//DeleteTask(ctx context.Context, req *ct.TaskID) (resp *ct.TaskID, err error)

	CreateSchedules(ctx context.Context, req *ct.CreateSchedulesRequest) (resp *ct.SchedulesID, err error)
	GetSchedulesByID(ctx context.Context, req *ct.SchedulesID) (resp *ct.Schedules, err error)

	CreateJournal(ctx context.Context, req *ct.CreateJournalRequest) (resp *ct.JournalID, err error)
	GetJournalByID(ctx context.Context, req *ct.JournalID) (resp *ct.Journal, err error)

	GetTeacherInfo(ctc context.Context, req *ct.TeacherGetTeacherInfoRequest) (resp *ct.TeacherGetTeacherInfoResponse, err error)
	GetSupporTeacherInfo(ctc context.Context, req *ct.TeacherGetTeacherInfoRequest) (resp *ct.TeacherGetSupportInfoResponse, err error)
}
