package keeper

import (
	"fmt"

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

	// set new SetCollateralAssetData
	k.SetCollateralAsset(ctx, minterAddress, collateralData)

	return nil
}
