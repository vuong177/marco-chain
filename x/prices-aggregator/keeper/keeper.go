package keeper

import (
	// "github.com/vuong177/macro/x/prices-aggregator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	paramSpace     paramtypes.Subspace

	// ibc keeper
	// portKeeper    types.PortKeeper
	// channelKeeper types.ChannelKeeper
	// scopedKeeper  types.ScopedKeeper
}