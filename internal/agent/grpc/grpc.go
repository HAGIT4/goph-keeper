package grcp

import (
	"context"

	pb "github.com/hagit4/goph-keeper/pkg/pb/goph-keeper"
	"google.golang.org/grpc"
)

type AgentGRPCinterface interface {
	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error)
	Login(ctx context.Context, req *pb.LoginRequest) (resp *pb.LoginResponse, err error)

	SaveLoginPass(ctx context.Context, req *pb.SaveLoginPassRequest) (resp *pb.SaveLoginPassResponse, err error)
	GetLoginPass(ctx context.Context, req *pb.GetLoginPassRequest) (resp *pb.GetLoginPassResponse, err error)
}

type agentGRPC struct {
	authClient      pb.AuthClient
	loginPassClient pb.LoginPassKeeperClient
}

type AgentGRPCoption func(ag *agentGRPC)

func NewAgentGRCP(opts ...AgentGRPCoption) (grpcClient *agentGRPC) {
	grpcClient = &agentGRPC{}
	for _, opt := range opts {
		opt(grpcClient)
	}
	return grpcClient
}

func WithClientConn(cc grpc.ClientConnInterface) AgentGRPCoption {
	return func(ag *agentGRPC) {
		ag.authClient = pb.NewAuthClient(cc)
		ag.loginPassClient = pb.NewLoginPassKeeperClient(cc)
	}
}
