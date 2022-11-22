package client

import (
	"context"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddLoginPassCmd() {
	saveLoginPass := &cobra.Command{
		Use:   "loginpass",
		Short: "save loginpass",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.SaveLoginPass(ctx)
		},
	}
	ac.createCmd.AddCommand(saveLoginPass)

	listLoginPassKeywordsCmd := &cobra.Command{
		Use:   "loginpass",
		Short: "List loginpass keywors",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.ListLoginPassKeywords(ctx)
		},
	}
	ac.listCmd.AddCommand(listLoginPassKeywordsCmd)

	getLoginPassCmd := &cobra.Command{
		Use:   "loginpass",
		Short: "Get loginpass data",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.GetLoginPass(ctx)
		},
	}
	ac.getCmd.AddCommand(getLoginPassCmd)
}
