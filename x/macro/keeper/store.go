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

// TODO: need to handle collateralAsset denom, now we consider collateralAsset denom is ATOM.
func (k Keeper) pricingCollateralAsset(ctx sdk.Context, collateralAssets sdk.Coins) sdk.Dec {
	collateralAssetValue := sdk.NewDec(0)

	for _, coin := range collateralAssets {
		price := k.GetPrice(ctx, coin.Denom)
		collateralAssetValue = collateralAssetValue.Add(price.MulInt(coin.Amount))
	}

	return collateralAssetValue
}

func (k Keeper) calculateCollateralRate(ctx sdk.Context, collateralAsset sdk.Coins, stableCoinAmount sdk.Dec) (sdk.Dec, error) {
	if stableCoinAmount == sdk.NewDec(0) {
		return sdk.Dec{}, types.ErrEmptyMintedStableCoin
	}

	price := k.pricingCollateralAsset(ctx, collateralAsset)
	collateralRate := stableCoinAmount.Quo(price)

	return collateralRate, nil
}

// TODO : implement oracle module
func (k Keeper) GetPrice(ctx sdk.Context, denom string) sdk.Dec {
	return sdk.NewDec(13)
}
