package client

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddBinaryCmd() {
	saveBinaryCmd := &cobra.Command{
		Use:   "bin",
		Short: "save bin",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := ac.service.SaveBinary(ctx)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	ac.createCmd.AddCommand(saveBinaryCmd)

	getBinaryCmd := &cobra.Command{
		Use:   "bin",
		Short: "get bin",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.GetBinary(ctx)
		},
	}
	ac.getCmd.AddCommand(getBinaryCmd)

	listBinaryKeywords := &cobra.Command{
		Use:   "bin",
		Short: "list binary keywords",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.ListBinaryKeywords(ctx)
		},
	}
	ac.listCmd.AddCommand(listBinaryKeywords)
}
