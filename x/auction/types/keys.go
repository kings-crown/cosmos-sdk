package types

const (
	// ModuleName is the name of the module
	ModuleName = "auction"

	// StoreKey is the store key string for the auction module
	StoreKey = ModuleName

	// RouterKey is the message route for the auction module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the auction module
	QuerierRoute = ModuleName
)

// Key Prefixes
const (
	AuctionKey = "auction-"
	BidKey     = "bid-"
)

// KeyPrefix creates a new key prefix for a given key type
func KeyPrefix(key string) []byte {
	return []byte(key)
}
