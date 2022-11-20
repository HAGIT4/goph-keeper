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
	ListLoginPassKeywords(ctx context.Context, req *pb.ListLoginPassKeywordsRequest) (resp *pb.ListLoginPassKeywordsResponse, err error)

	SaveTextData(ctx context.Context, req *pb.SaveTextDataRequest) (resp *pb.SaveTextDataResponse, err error)
	GetTextData(ctx context.Context, req *pb.GetTextDataRequest) (resp *pb.GetTextDataResponse, err error)

	SaveCardData(ctx context.Context, req *pb.SaveCardDataRequest) (resp *pb.SaveCardDataResponse, err error)
	GetCardData(ctx context.Context, req *pb.GetCardDataRequest) (resp *pb.GetCardDataResponse, err error)
	ListCardDataKeywords(ctx context.Context, req *pb.ListCardDataKeywordsRequest) (resp *pb.ListCardDataKeywordsResponse, err error)
}

type agentGRPC struct {
	authClient      pb.AuthClient
	loginPassClient pb.LoginPassKeeperClient
	textdataClient  pb.TextDataKeeperClient
	cardDataClient  pb.CardDataKeeperClient
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
