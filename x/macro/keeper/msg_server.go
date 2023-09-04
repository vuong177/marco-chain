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

func (s msgServer) MintStableCoin(goCtx context.Context, msg *types.MsgMintStableCoin) (*types.MsgMintStableCoinResponse, error) {
	return nil, nil
}

func (s msgServer) WithdrawCollateral(goCtx context.Context, msg *types.MsgWithdrawCollateral) (*types.MsgWithdrawCollateralResponse, error) {
	return nil, nil
}

func (s msgServer) Repay(goCtx context.Context, msg *types.MsgRepay) (*types.MsgRepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	accAddress, err := sdk.AccAddressFromBech32(msg.Repayer)
	if err != nil {
		return &types.MsgRepayResponse{}, err
	}
	err = s.handleRepay(ctx, accAddress, msg.Amount)
	if err != nil {
		return &types.MsgRepayResponse{}, err
	}
	return &types.MsgRepayResponse{}, nil
}