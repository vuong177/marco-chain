package types

const (
	// ModuleName store the name of module
	ModuleName = "prices-aggregator"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// Version defines the current version the IBC oracle module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to.
	PortID = "oracle"
)

var (
	PortKey                = KeyPrefix("band-oracle-port-")
	AssetsCountKey         = KeyPrefix("assets-count-")
	AssetsStoreByDenomKey  = KeyPrefix("assets-store-by-denom-")
	AssetsStoreBySymbolKey = KeyPrefix("assets-store-by-symbol-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetAssetByDenomKey(denom string) []byte {
	key := append(AssetsStoreByDenomKey, []byte(denom)...)
	return key
}

func GetAssetBySymbolKey(symbol string) []byte {
	key := append(AssetsStoreBySymbolKey, []byte(symbol)...)
	return key
}
