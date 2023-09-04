package keeper

import (
	"fmt"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/vuong177/macro/x/macro/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// handle first time mint/borrow uusd
func (k Keeper) handleMintStableCoin(ctx sdk.Context, minterAddress sdk.AccAddress, requestedAmount sdk.Int) error {
	collateralData, found := k.GetCollateralData(ctx, minterAddress)
	if !found {
		return types.ErrEmptyDepositAsset
	}
	coins := sdk.NewCoins(
		sdk.NewCoin(
			types.StableCoinDenom,
			requestedAmount,
		),
	)
	newMintedStableCoin := collateralData.MintedStableCoin.Add(sdk.NewDecFromInt(requestedAmount))
	rate, err := k.calculateCollateralRate(ctx, collateralData.CollateralAsset, newMintedStableCoin)
	if err != nil {
		return err
	}

	if types.MinimunCollateralRate.GTE(rate) {
		return types.ErrSmallRequestCollateralRate
	}

	// mint uusd and send it to minterAddress
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minterAddress, coins)
	if err != nil {
		return err
	}

	// save this infomation in module state
	// TODO: need to handle collateralAsset denom, now we consider collateralAsset denom is ATOM.
	collateralData.MintedStableCoin = collateralData.MintedStableCoin.Add(sdk.NewDecFromInt(requestedAmount))

	// set new SetCollateralData
	k.SetCollateralData(ctx, minterAddress, collateralData)

	return nil
}

// handleRepay handle repay process: user repay uUSD debt to increase collateral ratio
func (k Keeper) handleRepay(ctx sdk.Context, repayerAddress sdk.AccAddress, amount sdkmath.Int) error {
	collateralData, found := k.GetCollateralData(ctx, repayerAddress)
	if !found {
		return types.ErrCanNotFindDataOfUser
	}
	// check if amount is greater than amount of stablecoin minted, then assign amount to stablecoin minted
	if amount.GT(collateralData.MintedStableCoin.RoundInt()) {
		amount = collateralData.MintedStableCoin.RoundInt()
	}
	// burn amount of stablecoin of repayer
	coinsBurn := sdk.NewCoins(
		sdk.NewCoin(
			types.StableCoinDenom,
			amount,
		),
	)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, repayerAddress, types.ModuleName, coinsBurn)
	if err != nil {
		return fmt.Errorf("could not send coins from account %s to module %s. err: %s", repayerAddress, types.ModuleName, err.Error())
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to burn stablecoin in repay process %s", err.Error()))
		return fmt.Errorf("could not burn %v stablecoin in module account . err: %s", amount ,err.Error())
	}
	// update data of repayer in store
	collateralData.MintedStableCoin.Sub(sdkmath.LegacyDec(amount))

	k.UpdateReward(ctx, repayerAddress)

	//TODO: emit the event, I think we need to calculate collateral ratio of user after repay here?
	// Set CollateralData
	k.SetCollateralData(ctx, repayerAddress, collateralData)

	return nil
}

// TODO: Update this function when we implement reward module
// UpdateReward udpate the reward of user based on the change of amount of stablecoin of user
func (k *Keeper) UpdateReward(ctx sdk.Context, userAddress sdk.Address) {

}