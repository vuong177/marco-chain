package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

func (k Keeper) handleDeposit(ctx sdk.Context, depositAddress sdk.AccAddress, depositCoin sdk.Coin) error {
	k.validateDepositCoin(ctx, depositCoin.Denom)
	oldCollateralAssetData, found := k.GetBorrowerData(ctx, depositAddress)
	// if this's the first time user deposit, add save new CollateralData
	if !found {
		collateralAssetData := types.BorrowerData{
			CollateralAsset:  sdk.NewCoins(depositCoin),
			Borrowed: sdk.NewDec(0),
			IsRedemptionProvider: false,
		}
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositAddress, types.ModuleName, sdk.NewCoins(depositCoin))
		if err != nil {
			return err
		}
		k.SetBorrowerData(ctx, depositAddress, collateralAssetData)
		return nil
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositAddress, types.ModuleName, sdk.NewCoins(depositCoin))
	if err != nil {
		return err
	}
	// if user already deposit, calculate new collateral rate and set CollateralData
	newCollateralData := oldCollateralAssetData.CollateralAsset.Add(depositCoin)
	k.SetBorrowerData(
		ctx,
		depositAddress,
		types.BorrowerData{
			CollateralAsset:  newCollateralData,
			Borrowed: oldCollateralAssetData.Borrowed,
			IsRedemptionProvider: false,
		},
	)

	return nil
}

func (k Keeper) validateDepositCoin(ctx sdk.Context, denom string) error {
	return nil
}
