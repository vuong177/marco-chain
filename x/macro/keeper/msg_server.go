package keeper

import (
	"context"

	"github.com/vuong177/macro/x/marco/types"
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
