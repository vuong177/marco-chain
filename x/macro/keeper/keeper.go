package keeper

import (
	sdkmath "cosmossdk.io/math"
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
	borrowerData, found := k.GetBorrowerData(ctx, minterAddress)
	if !found {
		return types.ErrEmptyDepositAsset
	}
	coins := sdk.NewCoins(
		sdk.NewCoin(
			types.StableCoinDenom,
			requestedAmount,
		),
	)
	newMintedStableCoin := borrowerData.Borrowed.Add(sdk.NewDecFromInt(requestedAmount))
	rate, err := k.calculateCollateralRate(ctx, borrowerData.CollateralAsset, newMintedStableCoin)
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
	borrowerData.Borrowed = borrowerData.Borrowed.Add(sdk.NewDecFromInt(requestedAmount))

	// set new SetBorrowerData
	k.SetBorrowerData(ctx, borrowerData)

	return nil
}

// handleRepay handle repay process: repayer pay amount of uUSD for borrower's debt to increase borrower's collateral ratio
func (k Keeper) handleRepay(ctx sdk.Context, repayerAddress sdk.AccAddress, borrowerAddress sdk.AccAddress, uusdAmount sdkmath.LegacyDec) error {
	borrowerCollateralData, found := k.GetBorrowerData(ctx, borrowerAddress)
	if !found {
		return types.ErrCanNotFindCollateralData
	}
	// check if uusdAmount repay is greater than amount of uusd borrowed
	if uusdAmount.GT(borrowerCollateralData.Borrowed) {
		uusdAmount = borrowerCollateralData.Borrowed
	}
	// burn amount of stablecoin of repayer
	// round amount up to avoid under charging
	coinsBurn := sdk.NewCoins(
		sdk.NewCoin(
			types.StableCoinDenom,
			uusdAmount.Ceil().TruncateInt(),
		),
	)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, repayerAddress, types.ModuleName, coinsBurn)
	if err != nil {
		return fmt.Errorf("could not send coins from account %s to module %s. err: %s", repayerAddress, types.ModuleName, err.Error())
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, coinsBurn)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Failed to burn stablecoin in repay process %s", err.Error()))
		return fmt.Errorf("could not burn %v stablecoin in module account . err: %s", uusdAmount ,err.Error())
	}
	// update data of borrower in store
	borrowerCollateralData.Borrowed.Sub(uusdAmount)

	//TODO: emit the event, I think we need to calculate collateral ratio of user after repay here?
	// Set CollateralData
	k.SetBorrowerData(ctx, borrowerCollateralData)

	return nil
}

// handleBecomeRedemptionProvide update data of borrower to become a redemption provider
func (k Keeper) handleBecomeRedemptionProvide(ctx sdk.Context, borrower sdk.AccAddress) error {
	borrowerData, found := k.GetBorrowerData(ctx, borrower)
	if !found {
		// TODO: Update name types of errors
		return types.ErrCanNotFindCollateralData
	}
	// TODO: Is there any condition to become a redemption provider?
	borrowerData.IsRedemptionProvider = true
	k.SetBorrowerData(ctx, borrowerData)

	return nil
}

// Redeemer doesn't have to be a borrower

// handleRedeem handle process when redeemer redeem stablecoin to get stToken
func (k Keeper) handleRedeem(ctx sdk.Context, redeemer sdk.AccAddress, uusdAmount sdkmath.LegacyDec, denomRedeem string) error {
	redemptionProvider, err := k.getRedemptionProvider(ctx, uusdAmount, denomRedeem)
	if err != nil {
		return err
	}
	// redeemer pay redemptionProvider's debt to get stToken
	err = k.handleRepay(ctx, redeemer, sdk.AccAddress(redemptionProvider.Address), uusdAmount)

	// TODO: Need to check the calculations below
	// Calculate amount of collateral redeemed with redemption fee  = 0.05%
	assetPrice := k.GetPrice(ctx, denomRedeem)
	collateralAmount := uusdAmount.Mul(sdk.MustNewDecFromStr("100").Sub(types.RedemptionFee)).Quo(assetPrice).Quo(sdk.MustNewDecFromStr("100"))
	collateralAssetRedeemed := sdk.NewCoins(sdk.NewCoin(denomRedeem, collateralAmount.TruncateInt()))
	redemptionProvider.CollateralAsset = redemptionProvider.CollateralAsset.Sub(collateralAssetRedeemed...)
	
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(redemptionProvider.Address), types.ModuleName, collateralAssetRedeemed)
	if err != nil {
		return fmt.Errorf("could not send coins from redemption provider %s to module %s. err: %s", redemptionProvider.Address, types.ModuleName, err.Error())
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, redeemer, collateralAssetRedeemed)
	if err != nil {
		return fmt.Errorf("could not send coins from module %s to redeemer %s. err: %s", types.ModuleName, redeemer, err.Error())
	}

	k.SetBorrowerData(ctx, redemptionProvider)

	return nil
}

// TODO: Need to discuss the conditions below

// The provider's collateral asset is the highest in list of redemption providers
// The provider's debt must be equal to or above the requested uUSD amount
// The provider's collateral ratio must be at least 125%. (partially liquidation rate)

// getRedemptionProvider get the redemption provider for redeem progress
func (k Keeper) getRedemptionProvider(ctx sdk.Context, amount sdkmath.LegacyDec, denomRedeem string) (types.BorrowerData, error) {
	redemptionProvider := types.BorrowerData{}
	highestCollateralAsset := sdkmath.NewInt(0)

	redemptionProviders := k.GetAllRedemptionProviders(ctx)
	if len(redemptionProviders) == 0 {
		return types.BorrowerData{}, fmt.Errorf("There is no redemption provider now")
	}
	seedsOfRedemptionProvider := []types.BorrowerData{}
	for _, rp := range redemptionProviders {
		collateral_ratio, _ := k.calculateCollateralRate(ctx, rp.CollateralAsset, rp.Borrowed)
		if collateral_ratio.LT(types.ThresholdPartialLiquidationRate) {
			continue
		}
		if rp.Borrowed.LT(sdkmath.LegacyDec(amount)) {
			continue
		}
		seedsOfRedemptionProvider = append(seedsOfRedemptionProvider, rp)
		// TODO: in case we have two redemption providers have the same amount of collateral asset, how to solve it?
		// maybe we should passing it because after redeem the first one, another one will be the highest
		if rp.CollateralAsset.AmountOf(denomRedeem).GT(highestCollateralAsset) {
			highestCollateralAsset = rp.CollateralAsset.AmountOf(denomRedeem)
			redemptionProvider = rp
		}
	}
	if len(seedsOfRedemptionProvider) == 0 {
		return types.BorrowerData{}, fmt.Errorf("There is no redemption provider eligible for redeem progress")
	}

	return redemptionProvider, nil
}
