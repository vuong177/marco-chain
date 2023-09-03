package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

// GetLastGaugeID returns ID used last time
func (k Keeper) GetCollateralData(ctx sdk.Context, address sdk.AccAddress) (types.CollateralData, bool) {
	var collateralData types.CollateralData
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetKeyCollateralAssetData(address))
	if bz == nil {
		return types.CollateralData{}, false
	}

	k.cdc.Unmarshal(bz, &collateralData)
	return collateralData, true
}

// SetCollateralAsset save collateral asset used by `address`
func (k Keeper) SetCollateralAsset(ctx sdk.Context, address sdk.AccAddress, collateralData types.CollateralData) {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&collateralData)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetKeyCollateralAssetData(address), bz)
}

func (k Keeper) calculateCollateralRate(ctx sdk.Context, collateralAsset sdk.Coin, stableCoinAmount sdk.Dec) (sdk.Dec, error) {
	if stableCoinAmount == sdk.NewDec(0) {
		return sdk.Dec{}, types.ErrEmptyMintedStableCoin
	}

	price := k.GetPrice(ctx, collateralAsset.Denom)
	collateralRate := stableCoinAmount.Quo(price.MulInt(collateralAsset.Amount))

	return collateralRate, nil
}

// TODO : implement oracle module
func (k Keeper) GetPrice(ctx sdk.Context, denom string) sdk.Dec {
	return sdk.NewDec(13)
}
