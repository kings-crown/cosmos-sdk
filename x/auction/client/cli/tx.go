package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/your-github-username/your-chain-name/x/auction/types"
)

// GetTxCmd returns the transaction commands for the auction module
func GetTxCmd() *cobra.Command {
	auctionTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auction transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	auctionTxCmd.AddCommand(
		flags.PostCommands(
			GetCmdSubmitBid(),
			GetCmdRevealBid(),
		)...,
	)

	return auctionTxCmd
}

// GetCmdSubmitBid implements the submit bid command
func GetCmdSubmitBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-bid [auction-id] [commitment]",
		Short: "Submit a sealed bid for an auction",
		Long:  "Submit a sealed bid for an auction, with the commitment being the hash of the bid amount and a random nonce",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auctionID := args[0]
			commitment := args[1]

			msg := types.NewMsgSubmitBid(clientCtx.GetFromAddress(), auctionID, commitment)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdRevealBid implements the reveal bid command
func GetCmdRevealBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-bid [auction-id] [bid-amount] [nonce]",
		Short: "Reveal a previously submitted bid for an auction",
		Long:  "Reveal a previously submitted bid for an auction, including the bid amount and the nonce used when creating the commitment",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auctionID := args[0]
			bidAmountStr := args[1]
			nonce := args[2]

			bidAmount, err := sdk.ParseCoinsNormalized(bidAmountStr)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevealBid(clientCtx.GetFromAddress(), auctionID, bidAmount, nonce)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
