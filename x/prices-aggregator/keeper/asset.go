package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// GetCommentCount get the total number of comment
func (k Keeper) GetAssetsCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AssetsCountKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) AddAsset(ctx sdk.Context, denom string, symbol string) (uint64, error) {

}
