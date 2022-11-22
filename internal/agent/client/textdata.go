package client

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddTextDataCmd() {
	saveTextDataCmd := &cobra.Command{
		Use:   "text",
		Short: "save text",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if err := ac.service.SaveTextData(ctx); err != nil {
				fmt.Println(err)
			}
		},
	}
	ac.createCmd.AddCommand(saveTextDataCmd)

	listTextDataKeywordsCmd := &cobra.Command{
		Use:   "text",
		Short: "List text keywors",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if err := ac.service.ListTextDataKeywords(ctx); err != nil {
				fmt.Println(err)
			}
		},
	}
	ac.listCmd.AddCommand(listTextDataKeywordsCmd)

	getTextDataCmd := &cobra.Command{
		Use:   "text",
		Short: "Get text data",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if err := ac.service.GetTextData(ctx); err != nil {
				fmt.Println(err)
			}
		},
	}
	ac.getCmd.AddCommand(getTextDataCmd)
}
