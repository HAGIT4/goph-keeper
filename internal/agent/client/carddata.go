package client

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (ac *agentCli) AddCardDataCmd() {
	saveCardData := &cobra.Command{
		Use:   "carddata",
		Short: "save a card data",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.SaveCardData(ctx)
		},
	}
	ac.createCmd.AddCommand(saveCardData)

	getCardDataCmd := &cobra.Command{
		Use:   "carddata",
		Short: "Get card data",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			if err := ac.service.GetCardData(ctx); err != nil {
				fmt.Println(err)
			}
		},
	}
	ac.getCmd.AddCommand(getCardDataCmd)

	listCardDataCmd := &cobra.Command{
		Use:   "carddata",
		Short: "List card data keywords",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			_ = ac.service.ListCardDataKeywords(ctx)
		},
	}
	ac.listCmd.AddCommand(listCardDataCmd)
}
