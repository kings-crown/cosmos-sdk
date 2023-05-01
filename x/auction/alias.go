package auction

import (
	"github.com/your-github-username/your-chain-name/x/auction/client/cli"
	"github.com/your-github-username/your-chain-name/x/auction/client/rest"
	"github.com/your-github-username/your-chain-name/x/auction/types"
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	TypeMsgSubmitBid  = types.TypeMsgSubmitBid
	TypeMsgRevealBid  = types.TypeMsgRevealBid
	DefaultParamspace = types.DefaultParamspace
)

type (
	GenesisState        = types.GenesisState
	MsgSubmitBid        = types.MsgSubmitBid
	MsgRevealBid        = types.MsgRevealBid
	Params              = types.Params
	QuerierRoute        = types.QuerierRoute
	QueryEndpoints      = types.QueryEndpoints
	NewMsgSubmitBid     = types.NewMsgSubmitBid
	NewMsgRevealBid     = types.NewMsgRevealBid
	NewGenesisState     = types.NewGenesisState
	NewParams           = types.NewParams
	DefaultGenesisState = types.DefaultGenesisState
	DefaultParams       = types.DefaultParams
	ValidateGenesis     = types.ValidateGenesis
)

var (
	NewQuerier       = keeper.NewQuerier
	RegisterCodec    = types.RegisterCodec
	ModuleCdc        = types.ModuleCdc
	ErrInvalidBid    = types.ErrInvalidBid
	ErrInvalidReveal = types.ErrInvalidReveal
)

type (
	Keeper               = keeper.Keeper
	Auction              = types.Auction
	CLIQueryEndpointsCmd = cli.QueryEndpointsCmd
	CLITxCmd             = cli.TxCmd
	RESTHandlers         = rest.Handlers
)
