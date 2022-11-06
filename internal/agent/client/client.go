package client

import (
	"github.com/hagit4/goph-keeper/internal/agent/config"
	keeperGRPC "github.com/hagit4/goph-keeper/internal/agent/grpc"
	in "github.com/hagit4/goph-keeper/internal/agent/grpc/interceptor"
	"github.com/hagit4/goph-keeper/internal/agent/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type agentClient struct {
	cli     *agentCli
	conn    *grpc.ClientConn
	service service.AgentServiceInterface
}

func NewAgentClient() (client *agentClient, err error) {
	config, err := config.InitAgentConfig()
	if err != nil {
		return nil, err
	}

	authInterceptor := in.NewAuthInterceptor()
	conn, err := grpc.Dial(config.KeeperAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(authInterceptor),
	)
	if err != nil {
		return nil, err
	}

	agentGRPC := keeperGRPC.NewAgentGRCP(
		keeperGRPC.WithClientConn(conn),
	)
	service := service.NewAgentService(
		service.WithAgentGRPC(agentGRPC),
		service.WithEncKey(config.EncKey),
	)
	cli := NewAgentCli(
		WithAgentService(service),
	)
	cli.InitCommands()

	client = &agentClient{
		cli:     cli,
		conn:    conn,
		service: service,
	}
	return client, nil
}

func (a *agentClient) Execute() (err error) {
	if err = a.cli.Execute(); err != nil {
		return err
	}
	return nil
}
