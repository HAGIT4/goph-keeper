package grpc

import (
	"github.com/hagit4/goph-keeper/internal/keeper/service"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
)

type ServerGRPCinterface interface{}

type serviceGRPC struct {
	pb.UnimplementedAuthServer
	pb.UnimplementedLoginPassKeeperServer
	service service.KeeperServiceInterface
}

var _ pb.AuthServer = (*serviceGRPC)(nil)

type serviceGRPCoption func(*serviceGRPC)

func NewServerGrpc(opts ...serviceGRPCoption) (grpcServer *serviceGRPC) {
	grpcServer = &serviceGRPC{}
	for _, opt := range opts {
		opt(grpcServer)
	}
	return grpcServer
}

func WithService(service service.KeeperServiceInterface) serviceGRPCoption {
	return func(sg *serviceGRPC) {
		sg.service = service
	}
}
