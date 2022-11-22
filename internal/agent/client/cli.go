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
	defaultCreateCmd = cobra.Command{
		Use:   "create",
		Short: "create entry",
		Long:  "create entry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: create + data type")
		},
	}
	defaultListCmd = cobra.Command{
		Use:   "list",
		Short: "list entries",
		Long:  "list entries",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: list + data type")
		},
	}
	defaultGetCmd = cobra.Command{
		Use:   "get",
		Short: "get entry info",
		Long:  "get entry info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Usage: get + data type")
		},
	}
)

type agentCli struct {
	rootCmd   *cobra.Command
	createCmd *cobra.Command
	listCmd   *cobra.Command
	getCmd    *cobra.Command
	service   service.AgentServiceInterface
}

type agentCLIoption func(ac *agentCli)

func NewAgentCli(opts ...agentCLIoption) (cli *agentCli) {
	defaultRootCmd.AddCommand(&defaultCreateCmd)
	defaultRootCmd.AddCommand(&defaultListCmd)
	defaultRootCmd.AddCommand(&defaultGetCmd)
	cli = &agentCli{
		rootCmd:   &defaultRootCmd,
		createCmd: &defaultCreateCmd,
		listCmd:   &defaultListCmd,
		getCmd:    &defaultGetCmd,
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
	ac.AddCardDataCmd()
	ac.AddTextDataCmd()
	ac.AddBinaryCmd()
}

func (ac *agentCli) Execute() (err error) {
	if err = ac.rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
