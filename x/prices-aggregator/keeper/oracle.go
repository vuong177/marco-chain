package keeper

import (
	"encoding/binary"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

// GetNextClientID get next clientID
// TODO: testing
func (k Keeper) GetNextClientID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ClientIDCountKey)
	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetClientIDCount set client ID count
// TODO: testing
func (k Keeper) SetClientIDCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(types.ClientIDCountKey, bz)
}

// SetOracleRequestByClientID store oracle request by client ID
// TODO: testing
func (k Keeper) SetOracleRequestByClientID(ctx sdk.Context, oracleRequest types.OracleRequestPacketData) error {
	clientID, err := strconv.ParseUint(oracleRequest.ClientID, 10, 64)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetOracleRequestByClientIDKey(clientID)

	bz := k.cdc.MustMarshal(&oracleRequest)

	store.Set(key, bz)

	return nil
}

// GetOracleRequestByClientID get oracle request by client ID
// TODO: testing
func (k Keeper) GetOracleRequestByClientID(ctx sdk.Context, clientID uint64) (types.OracleRequestPacketData, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetOracleRequestByClientIDKey(clientID)

	bz := store.Get(key)
	if bz == nil {
		return types.OracleRequestPacketData{}, false
	}

	var oracleRequest types.OracleRequestPacketData
	k.cdc.MustUnmarshal(bz, &oracleRequest)

	return oracleRequest, true
}
