package grpc_client

import (
	"fmt"
	"managergetaway/config"
	"managergetaway/genproto/manager_service"
	"managergetaway/genproto/superadmin_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClientI ...
type GrpcClientI interface {
	AdministrationService() manager_service.ManagerAdministrationServiceServer
	StudentService() manager_service.ManagerStudentServiceServer
	TeacherService() manager_service.ManagerTeacherServiceServer
	GroupServiceC() manager_service.ManagerGroupServiceServer
	SupportTeacherService() manager_service.ManagerSupportTeacherServiceServer
	SuperadminService() superadmin_service.ManagerServiceServer
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %v port:%v err: %v",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}
	connOrder, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("order service dial hsot:%v port :%v err:%v",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"administration_service": manager_service.NewManagerAdministrationServiceClient(connUser),
			"student_service":        manager_service.NewManagerStudentServiceClient(connUser),
			"teacher_service":        manager_service.NewManagerTeacherServiceClient(connUser),
			"groupr_service":         manager_service.NewManagerGroupServiceClient(connUser),
			"supportteacher_service": manager_service.NewManagerSupportTeacherServiceClient(connUser),
			"admin_service":          superadmin_service.NewManagerServiceClient(connOrder),
		},
	}, nil
}

func (g *GrpcClient) AdministrationService() manager_service.ManagerAdministrationServiceClient {
	return g.connections["administration_service"].(manager_service.ManagerAdministrationServiceClient)
}

func (g *GrpcClient) StudentService() manager_service.ManagerStudentServiceClient {
	return g.connections["student_service"].(manager_service.ManagerStudentServiceClient)
}

func (g *GrpcClient) TeacherService() manager_service.ManagerTeacherServiceClient {
	return g.connections["teacher_service"].(manager_service.ManagerTeacherServiceClient)
}

func (g *GrpcClient) GroupServiceC() manager_service.ManagerGroupServiceClient {
	return g.connections["groupr_service"].(manager_service.ManagerGroupServiceClient)
}

func (g *GrpcClient) SupportTeacherService() manager_service.ManagerSupportTeacherServiceClient {
	return g.connections["supportteacher_service"].(manager_service.ManagerSupportTeacherServiceClient)
}

func (g *GrpcClient) SuperadminService() superadmin_service.ManagerServiceClient {
	return g.connections["admin_service"].(superadmin_service.ManagerServiceClient)
}
