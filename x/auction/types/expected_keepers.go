package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected bank keeper functions used by the auction module
type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// AuctionKeeper defines the expected auction keeper functions used by the auction module
type AuctionKeeper interface {
	GetAuction(ctx sdk.Context, auctionID string) (Auction, bool)
	SetAuction(ctx sdk.Context, auction Auction)
	DeleteAuction(ctx sdk.Context, auctionID string)
	GetBids(ctx sdk.Context, auctionID string) []Bid
	AddBid(ctx sdk.Context, auctionID string, bid Bid)
	GetAllAuctions(ctx sdk.Context) []Auction
	GetAuctionIds(ctx sdk.Context) []string
	CreateNewAuction(ctx sdk.Context, auction Auction) error
	TransferCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error
}
