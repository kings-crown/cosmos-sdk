package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/kings-crown/cosmos-sdk/tree/main/x/auction/types"
)

const (
	DefaultCommitDuration   = int64(100)
	DefaultRevealDuration   = int64(100)
	DefaultMinimumBidAmount = int64(1)
	DefaultAuctionTimeout   = int64(200)
)

var (
	KeyCommitDuration   = []byte("CommitDuration")
	KeyRevealDuration   = []byte("RevealDuration")
	KeyMinimumBidAmount = []byte("MinimumBidAmount")
	KeyAuctionTimeout   = []byte("AuctionTimeout")
)

type Params struct {
	CommitDuration   int64
	RevealDuration   int64
	MinimumBidAmount sdk.Coins
	AuctionTimeout   int64
}

func NewParams(commitDuration int64, revealDuration int64, minBidAmount sdk.Coins, auctionTimeout int64) Params {
	return Params{
		CommitDuration:   commitDuration,
		RevealDuration:   revealDuration,
		MinimumBidAmount: minBidAmount,
		AuctionTimeout:   auctionTimeout,
	}
}

func DefaultParams() Params {
	return Params{
		CommitDuration:   DefaultCommitDuration,
		RevealDuration:   DefaultRevealDuration,
		MinimumBidAmount: sdk.NewCoins(sdk.NewInt64Coin("stake", DefaultMinimumBidAmount)),
		AuctionTimeout:   DefaultAuctionTimeout,
	}
}

func (p Params) Validate() error {
	if p.CommitDuration <= 0 {
		return types.ErrInvalidCommitDuration
	}
	if p.RevealDuration <= 0 {
		return types.ErrInvalidRevealDuration
	}
	if !p.MinimumBidAmount.IsAllPositive() {
		return types.ErrInvalidMinimumBidAmount
	}
	if p.AuctionTimeout <= 0 {
		return types.ErrInvalidAuctionTimeout
	}
	return nil
}

func (k Keeper) GetParams(ctx sdk.Context) (params Params) {
	k.paramstore.Get(ctx, KeyCommitDuration, &params.CommitDuration)
	k.paramstore.Get(ctx, KeyRevealDuration, &params.RevealDuration)
	k.paramstore.Get(ctx, KeyMinimumBidAmount, &params.MinimumBidAmount)
	k.paramstore.Get(ctx, KeyAuctionTimeout, &params.AuctionTimeout)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params Params) {
	k.paramstore.Set(ctx, KeyCommitDuration, params.CommitDuration)
	k.paramstore.Set(ctx, KeyRevealDuration, params.RevealDuration)
	k.paramstore.Set(ctx, KeyMinimumBidAmount, params.MinimumBidAmount)
	k.paramstore.Set(ctx, KeyAuctionTimeout, params.AuctionTimeout)
}
