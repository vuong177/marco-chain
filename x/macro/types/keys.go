package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	KeyCollateralAssetData = []byte{0x01}

	// KeyCollateralAsset defines key for CollateralAsset
	KeyCollateralAssetRateSecondaryIndex = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetKeyCollateralAssetData(address sdk.AccAddress, tokenIndex []byte) []byte {
	bz := append(KeyCollateralAssetData, tokenIndex...)
	return append(bz, address...)
}

func GetKeyAddressCollateralAssetSecondaryIndex(address sdk.AccAddress, rate sdk.Dec) []byte {
	return append(KeyCollateralAssetRateSecondaryIndex, address...)
}
