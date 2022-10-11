package client

import (
	"context"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddRegisterCommand() {
	registerUserCmd := &cobra.Command{
		Use:   "register",
		Short: "Register a new user",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.RegisterUser(ctx)
		},
	}
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login to a keeper",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.Login(ctx)
		},
	}
	ac.rootCmd.AddCommand(registerUserCmd)
	ac.rootCmd.AddCommand(loginCmd)
}
