package grpc

import (
	"superadmin/config"
	"superadmin/genproto/superadmin_service"
	"superadmin/grpc/client"
	"superadmin/grpc/service"
	"superadmin/storage"

	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServciceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	superadmin_service.RegisterManagerServiceServer(grpcServer, service.NewSuperadminrService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)
	return
}
