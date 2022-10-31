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
}

type agentService struct {
	agentGRPC keeperGRPC.AgentGRPCinterface
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
