package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
)

// RegisterLegacyAminoCodec registers the necessary x/auction interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAuction{}, "auction/CreateAuction", nil)
	cdc.RegisterConcrete(&MsgSubmitBid{}, "auction/SubmitBid", nil)
	cdc.RegisterConcrete(&MsgRevealBid{}, "auction/RevealBid", nil)
	cdc.RegisterConcrete(&MsgEndAuction{}, "auction/EndAuction", nil)
}

// RegisterInterfaces registers the x/auction interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAuction{},
		&MsgSubmitBid{},
		&MsgRevealBid{},
		&MsgEndAuction{},
	)

	registry.RegisterInterface(
		"yourproject.yourusername.auction.v1beta1.Auction",
		(*Auction)(nil),
		&Auction{},
	)

	registry.RegisterInterface(
		"yourproject.yourusername.auction.v1beta1.Bid",
		(*Bid)(nil),
		&Bid{},
	)
}

// NewCodec returns a new codec for the x/auction module
func NewCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	RegisterLegacyAminoCodec(cdc)
	std.RegisterLegacyAminoCodec(cdc)
	return cdc
}

// NewAppCodec returns a new app codec for the x/auction module
func NewAppCodec() codec.Marshaler {
	return codec.NewAminoCodec(NewCodec())
}
