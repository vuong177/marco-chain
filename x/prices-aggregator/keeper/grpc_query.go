package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryAssetByDenom(goCtx context.Context, req *types.QueryAssetByDenomRequest) (*types.QueryAssetByDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	asset, found := k.GetAssetByDenom(ctx, req.Denom)
	if !found {
		return nil, types.ErrorAssetDenomNotFound
	}

	return &types.QueryAssetByDenomResponse{
		Asset: asset,
	}, nil
}
