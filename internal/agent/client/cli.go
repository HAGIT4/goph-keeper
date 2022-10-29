package client

import (
	"fmt"

	"github.com/hagit4/goph-keeper/internal/agent/service"
	"github.com/spf13/cobra"
)

var (
	defaultRootCmd = cobra.Command{
		Use:   "keep",
		Short: "password keeper",
		Long:  "Password keeper long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi keeper")
		},
	}
)

type agentCli struct {
	rootCmd *cobra.Command
	service service.AgentServiceInterface
}

type agentCLIoption func(ac *agentCli)

func NewAgentCli(opts ...agentCLIoption) (cli *agentCli) {
	cli = &agentCli{
		rootCmd: &defaultRootCmd,
	}
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}

func WithAgentService(as service.AgentServiceInterface) agentCLIoption {
	return func(ac *agentCli) {
		ac.service = as
	}
}

func (ac *agentCli) InitCommands() {
	ac.AddRegisterCommand()
	ac.AddLoginPassCmd()
}

func (ac *agentCli) Execute() (err error) {
	if err = ac.rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
