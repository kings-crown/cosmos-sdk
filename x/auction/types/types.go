package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Auction represents a Vickrey auction instance
type Auction struct {
	Id       string         `json:"id" yaml:"id"`
	Seller   sdk.AccAddress `json:"seller" yaml:"seller"`
	MinBid   sdk.Coins      `json:"min_bid" yaml:"min_bid"`
	EndTime  int64          `json:"end_time" yaml:"end_time"`
	IsClosed bool           `json:"is_closed" yaml:"is_closed"`
	Winner   sdk.AccAddress `json:"winner" yaml:"winner"`
}

// NewAuction creates a new Auction instance
func NewAuction(id string, seller sdk.AccAddress, minBid sdk.Coins, endTime int64) Auction {
	return Auction{
		Id:       id,
		Seller:   seller,
		MinBid:   minBid,
		EndTime:  endTime,
		IsClosed: false,
		Winner:   sdk.AccAddress{},
	}
}

// Bid represents a sealed bid in a Vickrey auction
type Bid struct {
	AuctionId string         `json:"auction_id" yaml:"auction_id"`
	Bidder    sdk.AccAddress `json:"bidder" yaml:"bidder"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
	Nonce     string         `json:"nonce" yaml:"nonce"`
}

// NewBid creates a new Bid instance
func NewBid(auctionId string, bidder sdk.AccAddress, amount sdk.Coins, nonce string) Bid {
	return Bid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Amount:    amount,
		Nonce:     nonce,
	}
}

// RevealBid represents a revealed bid in a Vickrey auction
type RevealBid struct {
	AuctionId  string         `json:"auction_id" yaml:"auction_id"`
	Bidder     sdk.AccAddress `json:"bidder" yaml:"bidder"`
	Amount     sdk.Coins      `json:"amount" yaml:"amount"`
	Nonce      string         `json:"nonce" yaml:"nonce"`
	IsRevealed bool           `json:"is_revealed" yaml:"is_revealed"`
}

// NewRevealBid creates a new RevealBid instance
func NewRevealBid(auctionId string, bidder sdk.AccAddress, amount sdk.Coins, nonce string, isRevealed bool) RevealBid {
	return RevealBid{
		AuctionId:  auctionId,
		Bidder:     bidder,
		Amount:     amount,
		Nonce:      nonce,
		IsRevealed: isRevealed,
	}
}
