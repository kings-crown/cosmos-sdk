package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/kings-crown/cosmos-sdk/tree/main/x/auction/types"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	bankKeeper types.BankKeeper
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, bankKeeper types.BankKeeper) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) SubmitBid(ctx sdk.Context, bidder sdk.AccAddress, auctionId string, commitment string) error {
	store := ctx.KVStore(k.storeKey)

	auction, err := k.GetAuction(ctx, auctionId)
	if err != nil {
		return err
	}

	if auction.State != types.AuctionOpen {
		return sdkerrors.Wrap(types.ErrAuctionNotOpen, fmt.Sprintf("auction state: %s", auction.State))
	}

	key := types.GetBidKey(auctionId, bidder)
	bid := types.Bid{
		AuctionId:     auctionId,
		BidderAddress: bidder.String(),
		Commitment:    commitment,
	}

	bz := k.cdc.MustMarshalBinaryBare(&bid)
	store.Set(key, bz)

	return nil
}

func (k Keeper) RevealBid(ctx sdk.Context, bidder sdk.AccAddress, auctionId string, bidAmount sdk.Coins, nonce string) error {
	store := ctx.KVStore(k.storeKey)

	auction, err := k.GetAuction(ctx, auctionId)
	if err != nil {
		return err
	}

	if auction.State != types.AuctionRevealing {
		return sdkerrors.Wrap(types.ErrAuctionNotRevealing, fmt.Sprintf("auction state: %s", auction.State))
	}

	key := types.GetBidKey(auctionId, bidder)
	bz := store.Get(key)
	if bz == nil {
		return sdkerrors.Wrap(types.ErrBidNotFound, fmt.Sprintf("bid for auction: %s", auctionId))
	}

	var bid types.Bid
	k.cdc.MustUnmarshalBinaryBare(bz, &bid)

	if bid.Commitment != types.GenerateCommitment(bidAmount, nonce) {
		return sdkerrors.Wrap(types.ErrInvalidCommitment, "the revealed bid does not match the commitment")
	}

	bid.BidAmount = bidAmount
	bz = k.cdc.MustMarshalBinaryBare(&bid)
	store.Set(key, bz)

	return nil
}

func (k Keeper) GetAuction(ctx sdk.Context, auctionID string) (auction types.Auction, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.KeyPrefix(types.AuctionKey + auctionID))
	if value == nil {
		return auction, false
	}

	k.cdc.MustUnmarshalBinaryBare(value, &auction)
	return auction, true
}

func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&auction)
	store.Set(types.KeyPrefix(types.AuctionKey+auction.Id), bz)
}

func (k Keeper) GetBids(ctx sdk.Context, auctionID string) (bids []types.Bid) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.BidKey+auctionID))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bid)
		bids = append(bids, bid)
	}

	return bids
}

func (k Keeper) AddBid(ctx sdk.Context, auctionID string, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&bid)
	store.Set(types.KeyPrefix(types.BidKey+auctionID+bid.Bidder), bz)
}

func (k Keeper) GetAllAuctions(ctx sdk.Context) (auctions []types.Auction) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AuctionKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &auction)
		auctions = append(auctions, auction)
	}

	return auctions
}

func (k Keeper) GetAuctionIds(ctx sdk.Context) (ids []string) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AuctionKey))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		ids = append(ids, id)
	}

	return ids
}

func (k Keeper) CreateNewAuction(ctx sdk.Context, auction types.Auction) error {
	if _, found := k.GetAuction(ctx, auction.Id); found {
		return fmt.Errorf("auction with id %s already exists", auction.Id)
	}

	k.SetAuction(ctx, auction)
	return nil
}

func (k Keeper) DeleteAuction(ctx sdk.Context, auctionID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyPrefix(types.AuctionKey + auctionID))
}

func (k Keeper) TransferCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) error {
	err := k.bankKeeper.SendCoins(ctx, from, to, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetNextAuctionID(ctx sdk.Context) uint64 {
	var auctionID uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefix(types.AuctionIDKey))

	if bz == nil {
		auctionID = 1
	} else {
		k.cdc.MustUnmarshalBinaryBare(bz, &auctionID)
		auctionID++
	}

	bz = k.cdc.MustMarshalBinaryBare(&auctionID)
	store.Set(types.KeyPrefix(types.AuctionIDKey), bz)

	return auctionID
}

func (k Keeper) SetAuctionID(ctx sdk.Context, auctionID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryBare(&auctionID)
	store.Set(types.KeyPrefix(types.AuctionIDKey), bz)
}
