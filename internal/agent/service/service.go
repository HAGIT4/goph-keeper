package service

import (
	"context"

	keeperGRPC "github.com/hagit4/goph-keeper/internal/agent/grpc"
)

type AgentServiceInterface interface {
	RegisterUser(ctx context.Context) (err error)
	Login(ctx context.Context) (err error)

	SaveLoginPass(ctx context.Context) (err error)
	ListLoginPassKeywords(ctx context.Context) (err error)
	GetLoginPass(ctx context.Context) (err error)

	SaveTextData(ctx context.Context) (err error)
	GetTextData(ctx context.Context) (err error)
	ListTextDataKeywords(ctx context.Context) (err error)

	SaveCardData(ctx context.Context) (err error)
	GetCardData(ctx context.Context) (err error)
	ListCardDataKeywords(ctx context.Context) (err error)

	SaveBinary(ctx context.Context) (err error)
	GetBinary(ctx context.Context) (err error)
	ListBinaryKeywords(ctx context.Context) (err error)
}

type agentService struct {
	agentGRPC keeperGRPC.AgentGRPCinterface
	encKey    []byte
}

type AgentServiceOption func(as *agentService)

func NewAgentService(opts ...AgentServiceOption) (service *agentService) {
	service = &agentService{}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

func WithAgentGRPC(ag keeperGRPC.AgentGRPCinterface) AgentServiceOption {
	return func(as *agentService) {
		as.agentGRPC = ag
	}
}

func WithEncKey(encKey []byte) AgentServiceOption {
	return func(as *agentService) {
		as.encKey = encKey
	}
}
