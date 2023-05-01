package auction

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/kings-crown/cosmos-sdk/tree/main/x/auction/keeper"
	"github.com/kings-crown/cosmos-sdk/tree/main/x/auction/types"
)

func NewHandler(k keeper.Keeper, bk bank.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case *types.MsgSubmitBid:
			return handleMsgSubmitBid(ctx, k, bk, msg)
		case *types.MsgRevealBid:
			return handleMsgRevealBid(ctx, k, bk, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized auction Msg type: %T", msg))
		}
	}
}

func handleMsgSubmitBid(ctx sdk.Context, k keeper.Keeper, bk bank.Keeper, msg *types.MsgSubmitBid) (*sdk.Result, error) {
	bidder, err := sdk.AccAddressFromBech32(msg.BidderAddress)
	if err != nil {
		return nil, err
	}

	err = k.SubmitBid(ctx, bidder, msg.AuctionId, msg.Commitment)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgRevealBid(ctx sdk.Context, k keeper.Keeper, bk bank.Keeper, msg *types.MsgRevealBid) (*sdk.Result, error) {
	bidder, err := sdk.AccAddressFromBech32(msg.BidderAddress)
	if err != nil {
		return nil, err
	}

	err = k.RevealBid(ctx, bidder, msg.AuctionId, msg.BidAmount, msg.Nonce)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
