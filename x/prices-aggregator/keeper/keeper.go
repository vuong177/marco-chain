package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   storetypes.StoreKey
	paramSpace paramtypes.Subspace

	// ibc keeper
	// portKeeper    types.PortKeeper
	// channelKeeper types.ChannelKeeper
	// scopedKeeper  types.ScopedKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ps paramtypes.Subspace,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: ps,
	}
}
