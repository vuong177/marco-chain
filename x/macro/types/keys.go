package types

const (
	// ModuleName defines the module name
	ModuleName = "marco"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_marco"
)

var (
	// KeyCollateralAsset defines key for CollateralAsset
	KeyCollateralAsset = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
