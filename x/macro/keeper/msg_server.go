package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/macro/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (s msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}
	err = s.handleDeposit(ctx, accAddress, msg.DepositCoin)
	if err != nil {
		return &types.MsgDepositResponse{}, err
	}
	return &types.MsgDepositResponse{}, nil
}

func (s msgServer) MintStableCoin(goCtx context.Context, msg *types.MsgMintStableCoin) (*types.MsgMintStableCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddress, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return &types.MsgMintStableCoinResponse{}, err
	}
	err = s.handleMintStableCoin(ctx, accAddress, msg.RequestAmount)
	if err != nil {
		return &types.MsgMintStableCoinResponse{}, err
	}
	return &types.MsgMintStableCoinResponse{}, nil
}

func (s msgServer) WithdrawCollateral(goCtx context.Context, msg *types.MsgWithdrawCollateral) (*types.MsgWithdrawCollateralResponse, error) {
	return nil, nil
}

func (s msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	repayerAddress, err := sdk.AccAddressFromBech32(msg.Repayer)
	if err != nil {
		return &types.MsgRepayResponse{}, err
	}
	borrowerAddress, err := sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return &types.MsgRepayResponse{}, err
	}
	err = s.handleRepay(ctx, repayerAddress, borrowerAddress, msg.Amount)
	if err != nil {
		return &types.MsgRepayResponse{}, err
	}
	return &types.MsgRepayResponse{}, nil
}

func (s msgServer) BecomeRedemptionProvider(goCtx context.Context, msg *types.MsgBecomeRedemptionProvider) (*types.MsgBecomeRedemptionProvider, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddress, err := sdk.AccAddressFromBech32(msg.RedemptionProvider)
	if err != nil {
		return &types.MsgBecomeRedemptionProvider{}, err
	}
	err = s.handleBecomeRedemptionProvide(ctx, accAddress)
	if err != nil {
		return &types.MsgBecomeRedemptionProvider{}, err
	}
	return &types.MsgBecomeRedemptionProvider{}, nil
}

func (s msgServer) Redeem(goCtx context.Context, msg *types.MsgRedeem) (*types.MsgRedeemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddress, err := sdk.AccAddressFromBech32(msg.Redeemer)
	if err != nil {
		return &types.MsgRedeemResponse{}, err
	}
	err = s.handleRedeem(ctx, accAddress, msg.Amount, msg.DenomRedeem)
	if err != nil {
		return &types.MsgRedeemResponse{}, err
	}
	return &types.MsgRedeemResponse{}, nil
}