package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/your-github-username/your-chain-name/x/auction/types"
)

// GetQueryCmd returns the cli query commands for the auction module
func GetQueryCmd(queryRoute string) *cobra.Command {
	auctionQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the auction module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	auctionQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdGetAuction(queryRoute),
		)...,
	)

	return auctionQueryCmd
}

// GetCmdGetAuction implements the query auction command
func GetCmdGetAuction(queryRoute string) *cobra.Command {
	return &cobra.Command{
		Use:   "get [auction-id]",
		Short: "Query details of an auction by auction ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			auctionID := args[0]

			res, err := queryClient.Auction(cmd.Context(), &types.QueryGetAuctionRequest{AuctionId: auctionID})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res.Auction)
		},
	}
}
