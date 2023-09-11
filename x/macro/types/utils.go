package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	MinimunCollateralRate           = sdk.NewDecWithPrec(150, 2)
	MaximumCollateralRate           = sdk.NewDecWithPrec(300, 2)
	ThresholdPartialLiquidationRate = sdk.NewDecWithPrec(125, 2)
)

const (
	StableCoinDenom = "uusd"
)

func DecToDecBytes(dec sdk.Dec) []byte {
	return dec.BigInt().Bytes()
}
