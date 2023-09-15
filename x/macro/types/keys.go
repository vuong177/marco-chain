package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "macro"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_macro"
)

var (
	// KeyBorrowerData defines key for BorrowerData
	KeyBorrowerData = []byte{0x01}

	// KeyCollateralAsset defines key for CollateralAsset
	KeyCollateralAssetRateSecondaryIndex = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetKeyBorrowerData(address sdk.AccAddress) []byte {
	bz := append(KeyBorrowerData)
	return append(bz, address...)
}

func GetKeyAddressCollateralAssetSecondaryIndex(address sdk.AccAddress, rate sdk.Dec) []byte {
	return append(KeyCollateralAssetRateSecondaryIndex, address...)
}
