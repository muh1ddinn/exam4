package grpc

import (
	"crmservice/config"
	"crmservice/genproto/manager_service"
	"crmservice/grpc/client"
	"crmservice/grpc/service"
	"crmservice/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServciceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	manager_service.RegisterManagerStudentServiceServer(grpcServer, service.NewManagerService(cfg, log, strg, srvc))
	manager_service.RegisterManagerTeacherServiceServer(grpcServer, service.NewManagerteacherService(cfg, log, strg, srvc))
	manager_service.RegisterManagerSupportTeacherServiceServer(grpcServer, service.NewManagersupportteacherService(cfg, log, strg, srvc))
	manager_service.RegisterManagerAdministrationServiceServer(grpcServer, service.NewManageradministrationService(cfg, log, strg, srvc))
	manager_service.RegisterManagerGroupServiceServer(grpcServer, service.NewManagerServicegroup(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
