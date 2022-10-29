package server

import (
	"net"

	"github.com/hagit4/goph-keeper/internal/keeper/config"
	keeperGRCP "github.com/hagit4/goph-keeper/internal/keeper/grpc"
	"github.com/hagit4/goph-keeper/internal/keeper/service"
	"github.com/hagit4/goph-keeper/internal/keeper/storage/postgres"
	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type KeeperServerInterface interface{}

type keeperServer struct {
	grpcService keeperGRCP.ServerGRPCinterface
	grpcServer  *grpc.Server
	service     service.KeeperServiceInterface
}

func NewKeeperServer() (keeper *keeperServer, err error) {
	config, err := config.InitKeeperConfig()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("postgres", config.DbString)
	if err != nil {
		return nil, err
	}
	storage := postgres.NewKeeperPostgresStorage(
		postgres.WithConnection(db),
	)

	tm, err := service.NewTokenMaker(config.SymmKey)
	if err != nil {
		return nil, err
	}

	service := service.NewKeeperService(
		service.WithStorage(storage),
		service.WithTokenMaker(tm),
	)

	grpcService := keeperGRCP.NewServerGrpc(
		keeperGRCP.WithService(service),
	)

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, grpcService)
	pb.RegisterLoginPassKeeperServer(grpcServer, grpcService)
	reflection.Register(grpcServer)

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
