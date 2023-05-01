package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAuction = "create_auction"
	TypeMsgSubmitBid     = "submit_bid"
	TypeMsgRevealBid     = "reveal_bid"
)

var (
	_ sdk.Msg = &MsgCreateAuction{}
	_ sdk.Msg = &MsgSubmitBid{}
	_ sdk.Msg = &MsgRevealBid{}
)

// MsgCreateAuction defines a new auction creation message
type MsgCreateAuction struct {
	Id      string         `json:"id" yaml:"id"`
	Seller  sdk.AccAddress `json:"seller" yaml:"seller"`
	MinBid  sdk.Coins      `json:"min_bid" yaml:"min_bid"`
	EndTime int64          `json:"end_time" yaml:"end_time"`
}

// NewMsgCreateAuction creates a new MsgCreateAuction instance
func NewMsgCreateAuction(id string, seller sdk.AccAddress, minBid sdk.Coins, endTime int64) MsgCreateAuction {
	return MsgCreateAuction{
		Id:      id,
		Seller:  seller,
		MinBid:  minBid,
		EndTime: endTime,
	}
}

// Route returns the module route for MsgCreateAuction
func (msg MsgCreateAuction) Route() string { return RouterKey }

// Type returns the message type for MsgCreateAuction
func (msg MsgCreateAuction) Type() string { return TypeMsgCreateAuction }

// ValidateBasic performs basic validation of MsgCreateAuction
func (msg MsgCreateAuction) ValidateBasic() error {
	if msg.Seller.Empty() {
		return errors.ErrInvalidAddress
	}
	if !msg.MinBid.IsValid() || msg.MinBid.IsZero() {
		return errors.ErrInvalidCoins
	}
	if msg.EndTime <= 0 {
		return errors.Wrap(errors.ErrInvalidRequest, "end time must be greater than zero")
	}
	return nil
}

// GetSignBytes returns the canonical byte representation of MsgCreateAuction
func (msg MsgCreateAuction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for MsgCreateAuction
func (msg MsgCreateAuction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}

// MsgSubmitBid defines a new bid submission message
type MsgSubmitBid struct {
	AuctionId string         `json:"auction_id" yaml:"auction_id"`
	Bidder    sdk.AccAddress `json:"bidder" yaml:"bidder"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgSubmitBid creates a new MsgSubmitBid instance
func NewMsgSubmitBid(auctionId string, bidder sdk.AccAddress, amount sdk.Coins) MsgSubmitBid {
	return MsgSubmitBid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Amount:    amount,
	}
}

// Route returns the module route for MsgSubmitBid
func (msg MsgSubmitBid) Route() string { return RouterKey }

// Type returns the message type for MsgSubmitBid
func (msg MsgSubmitBid) Type() string { return TypeMsgSubmitBid }

// ValidateBasic performs basic validation of MsgSubmitBid
func (msg MsgSubmitBid) ValidateBasic() error {
	if msg.Bidder.Empty() {
		return errors.ErrInvalidAddress
	}
	if !msg.Amount.IsValid() || msg; Amount.IsZero() {
		return errors.ErrInvalidCoins
	}
	return nil
}

// GetSignBytes returns the canonical byte representation of MsgSubmitBid
func (msg MsgSubmitBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for MsgSubmitBid
func (msg MsgSubmitBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Bidder}
}

// MsgRevealBid defines a new bid reveal message
type MsgRevealBid struct {
	AuctionId string         `json:"auction_id" yaml:"auction_id"`
	Bidder    sdk.AccAddress `json:"bidder" yaml:"bidder"`
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`
}

// NewMsgRevealBid creates a new MsgRevealBid instance
func NewMsgRevealBid(auctionId string, bidder sdk.AccAddress, amount sdk.Coins) MsgRevealBid {
	return MsgRevealBid{
		AuctionId: auctionId,
		Bidder:    bidder,
		Amount:    amount,
	}
}

// Route returns the module route for MsgRevealBid
func (msg MsgRevealBid) Route() string { return RouterKey }

// Type returns the message type for MsgRevealBid
func (msg MsgRevealBid) Type() string { return TypeMsgRevealBid }

// ValidateBasic performs basic validation of MsgRevealBid
func (msg MsgRevealBid) ValidateBasic() error {
	if msg.Bidder.Empty() {
		return errors.ErrInvalidAddress
	}
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return errors.ErrInvalidCoins
	}
	return nil
}

// GetSignBytes returns the canonical byte representation of MsgRevealBid
func (msg MsgRevealBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for MsgRevealBid
func (msg MsgRevealBid) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Bidder}
}
