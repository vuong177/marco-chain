package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

func (k Keeper) handleDeposit(ctx sdk.Context, depositAddress sdk.AccAddress, depositCoin sdk.Coin) error {
	k.validateDepositCoin(ctx, depositCoin.Denom)
	oldCollateralAssetData, found := k.GetCollateralData(ctx, depositAddress)
	// if this's the first time user deposit, add save new CollateralData
	if !found {
		collateralAssetData := types.CollateralData{
			CollateralAsset:  sdk.NewCoins(depositCoin),
			MintedStableCoin: sdk.NewDec(0),
		}
		k.SetCollateralAsset(ctx, depositAddress, collateralAssetData)
		return nil
	}

	// if user already deposit, calculate new collateral rate and set CollateralData
	newCollateralAssetData := oldCollateralAssetData.CollateralAsset.Add(depositCoin)
	k.SetCollateralAsset(
		ctx,
		depositAddress,
		types.CollateralData{
			CollateralAsset:  newCollateralAssetData,
			MintedStableCoin: oldCollateralAssetData.MintedStableCoin,
		},
	)

	return nil
}

func (k Keeper) validateDepositCoin(ctx sdk.Context, denom string) error {
	return nil
}
