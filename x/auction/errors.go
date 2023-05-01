package auction

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ErrInvalidAuctionID is an error for when an invalid auction ID is provided.
var ErrInvalidAuctionID = sdkerrors.Register(ModuleName, 1, "Invalid auction ID")

// ErrAuctionNotFound is an error for when an auction is not found.
var ErrAuctionNotFound = sdkerrors.Register(ModuleName, 2, "Auction not found")

// ErrInvalidBid is an error for when an invalid bid is submitted.
var ErrInvalidBid = sdkerrors.Register(ModuleName, 3, "Invalid bid")

// ErrBidNotFound is an error for when a bid is not found.
var ErrBidNotFound = sdkerrors.Register(ModuleName, 4, "Bid not found")

// ErrInvalidReveal is an error for when an invalid bid reveal is submitted.
var ErrInvalidReveal = sdkerrors.Register(ModuleName, 5, "Invalid bid reveal")

// ErrRevealNotFound is an error for when a reveal is not found.
var ErrRevealNotFound = sdkerrors.Register(ModuleName, 6, "Reveal not found")

// ErrAuctionAlreadyEnded is an error for when an action is taken on an already ended auction.
var ErrAuctionAlreadyEnded = sdkerrors.Register(ModuleName, 7, "Auction already ended")

// ErrAuctionNotEnded is an error for when an action is taken on an auction that has not yet ended.
var ErrAuctionNotEnded = sdkerrors.Register(ModuleName, 8, "Auction has not ended")

// ErrInvalidCommitment is an error for when an invalid commitment is provided.
var ErrInvalidCommitment = sdkerrors.Register(ModuleName, 9, "Invalid commitment")

func FormatError(err error, id string) error {
	return fmt.Errorf("%w: %s", err, id)
}
