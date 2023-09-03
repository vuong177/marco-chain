package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func DecToDecBytes(dec sdk.Dec) []byte {
	return dec.BigInt().Bytes()
}
