package main

import (
	"managergetaway/api"
	"managergetaway/config"
	"managergetaway/pkg/grpc_client"
	"managergetaway/pkg/logger"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
)

// initDeps initializes the dependencies such as config, logger, and gRPC client
func initDeps() {
	var err error
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "customer-api-gateway")

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(cfg.HTTPPort)
}

// go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"
