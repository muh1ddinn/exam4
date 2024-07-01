package client

import "crmservice/config"

type ServciceManagerI interface{}

type grpcClients struct{}

func NewGrpcClients(cfg config.Config) (ServciceManagerI, error) {

	return &grpcClients{}, nil
}
