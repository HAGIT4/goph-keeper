package client

import (
	"context"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddLoginPassCmd() {
	saveLoginPass := &cobra.Command{
		Use:   "save",
		Short: "save a loginpass",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.SaveLoginPass(ctx)
		},
	}
	ac.rootCmd.AddCommand(saveLoginPass)
}
