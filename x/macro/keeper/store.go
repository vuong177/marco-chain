package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

// GetLastGaugeID returns ID used last time
func (k Keeper) GetCollateralAsset(ctx sdk.Context) sdk.Coins {
	var collateralAsset string
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.KeyCollateralAsset)
	if bz == nil {
		return sdk.Coins{}
	}

	k.cdc.Unmarshal(bz, &collateralAsset)
	return collateralAsset
}

// SetLastGaugeID save collateral asset used
func (k Keeper) SetCollateralAsset(ctx sdk.Context, ID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyCollateralAsset, sdk.Uint64ToBigEndian(ID))
}
