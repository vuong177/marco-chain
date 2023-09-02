package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CoinsBytesToCoins(bz []byte) sdk.Coins {
	dec := sdk.NewCoin(new(big.Int).SetBytes(bz), sdk.Precision)
	if dec.IsNil() {
		return sdk.ZeroDec()
	}
	return dec
}
