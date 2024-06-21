package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
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

func (k Keeper) AddAssetPricesTrackingList(goCtx context.Context, msg *types.MsgAddAssetPricesTrackingList) (*types.MsgAddAssetPricesTrackingListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	id, err := k.AddAsset(ctx, msg.Denom, msg.Symbol)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddAssetPricesTrackingListResponse{Id: id}, nil
}

func (k Keeper) DeleteAssetPricesTrackingList(goCtx context.Context, msg *types.MsgDeleteAssetPricesTrackingList) (*types.MsgDeleteAssetPricesTrackingListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	err := k.DeleteAsset(ctx, msg.Denom)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteAssetPricesTrackingListResponse{}, nil
}
