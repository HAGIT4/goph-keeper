package server

import (
	"net"

	keeperGRCP "github.com/hagit4/goph-keeper/internal/keeper/grpc"
	"github.com/hagit4/goph-keeper/internal/keeper/service"
	"github.com/hagit4/goph-keeper/internal/keeper/storage"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc"
)

type KeeperServerInterface interface{}

type keeperServer struct {
	grpcService keeperGRCP.ServerGRPCinterface
	grpcServer  *grpc.Server
	service     service.KeeperServiceInterface
}

func NewKeeperServer() (keeper *keeperServer, err error) {
	storage := storage.NewKeeperPostgresStorage()

	service := service.NewKeeperService(
		service.WithStorage(storage),
	)

	grpcService := keeperGRCP.NewServerGrpc(
		keeperGRCP.WithService(service),
	)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, grpcService)

	keeper = &keeperServer{
		grpcService: grpcService,
		grpcServer:  grpcServer,
		service:     service,
	}
	return keeper, nil
}

func (ks *keeperServer) Serve() (err error) {
	conn, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}
	if err = ks.grpcServer.Serve(conn); err != nil {
		return err
	}
	return nil
}
