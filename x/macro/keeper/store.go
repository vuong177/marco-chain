package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

// GetLastGaugeID returns ID used last time
func (k Keeper) GetCollateralData(ctx sdk.Context, address sdk.AccAddress, tokenIndex []byte) types.CollateralData {
	var collateralData types.CollateralData
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetKeyCollateralAssetData(address, tokenIndex))
	if bz == nil {
		return types.CollateralData{}
	}

	k.cdc.Unmarshal(bz, &collateralData)
	return collateralData
}

// SetCollateralAsset save collateral asset used by `address`
func (k Keeper) SetCollateralAsset(ctx sdk.Context, address sdk.AccAddress, tokenIndex []byte, collateralData types.CollateralData) {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&collateralData)
	if err != nil {
		panic(err)
	}
	collateralRate := k.calculateCollateralRate(ctx, collateralData.CollateralAsset, collateralData.MintedStableCoin)
	k.setAbsoluteRatePositionSecondaryIndex(ctx, address, collateralRate)
	store.Set(types.GetKeyCollateralAssetData(address, tokenIndex), bz)
}

// setAbsoluteRatePositionSecondaryIndex save secondary index (key | AbsoluteCollateralRate | tokenIndex | Address| )
func (k Keeper) setAbsoluteRatePositionSecondaryIndex(ctx sdk.Context, address sdk.AccAddress, collateralRate sdk.Dec) {
	key := types.GetKeyAddressCollateralAssetSecondaryIndex(address, collateralRate)
	store := ctx.KVStore(k.storeKey)

	store.Set(key, []byte{})
}

func (k Keeper) calculateCollateralRate(ctx sdk.Context, collateralAsset sdk.Coin, stableCoin sdk.Dec) sdk.Dec {
	price := k.GetPrice(ctx, collateralAsset.Denom)
	collateralRate := stableCoin.Quo(price.MulInt(collateralAsset.Amount))

	return collateralRate
}

// TODO : implement oracle module
func (k Keeper) GetPrice(ctx sdk.Context, denom string) sdk.Dec {
	return sdk.NewDec(5)
}
