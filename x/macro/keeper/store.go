package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

// GetBorrowerData get borrower data
func (k Keeper) GetBorrowerData(ctx sdk.Context, address sdk.AccAddress) (types.BorrowerData, bool) {
	var borrowerData types.BorrowerData
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetKeyBorrowerData(address))
	if bz == nil {
		return types.BorrowerData{}, false
	}

	k.cdc.Unmarshal(bz, &borrowerData)
	return borrowerData, true
}

// SetBorrowerData save collateral data
func (k Keeper) SetBorrowerData(ctx sdk.Context, address sdk.AccAddress, borrowerData types.BorrowerData) {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&borrowerData)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetKeyBorrowerData(address), bz)
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
	collateralRate := price.Quo(stableCoinAmount)

	return collateralRate, nil
}

// TODO : implement oracle module
func (k Keeper) GetPrice(ctx sdk.Context, denom string) sdk.Dec {
	return sdk.NewDec(13)
}
