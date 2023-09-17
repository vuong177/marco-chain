package keeper

import (
	"encoding/binary"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

// GetAssetsCount get the total number of assets
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

// SetAssetsCount set the total number of assets
func (k Keeper) SetAssetsCount(ctx sdk.Context, count uint64) {
	store := ctx.KVStore(k.storeKey)
	key := store.Get(types.AssetsCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(key, bz)
}

// AddAsset add new asset to store
func (k Keeper) AddAsset(ctx sdk.Context, denom string, symbol string) (uint64, error) {
	symbol = strings.ToUpper(symbol)
	store := ctx.KVStore(k.storeKey)
	keyDenom := types.GetAssetByDenomKey(denom)
	keySymbol := types.GetAssetBySymbolKey(symbol)

	if store.Has(keyDenom) || store.Has(keySymbol) {
		return 0, errorsmod.Wrapf(types.ErrorDuplicateAsset, "duplicate asset denom")
	}

	ID := k.GetAssetsCount(ctx)

	asset := types.Asset{
		Id:     ID,
		Denom:  denom,
		Symbol: symbol,
	}

	bz := k.cdc.MustMarshal(&asset)

	store.Set(keyDenom, bz)
	store.Set(keySymbol, bz)

	k.SetAssetsCount(ctx, ID+1)

	return ID, nil
}

func (k Keeper) GetAssetByDenom(ctx sdk.Context, denom string) (types.Asset, bool) {
	store := ctx.KVStore(k.storeKey)
	keyDenom := types.GetAssetByDenomKey(denom)

	bz := store.Get(keyDenom)
	if bz == nil {
		return types.Asset{}, false
	}

	var asset types.Asset
	k.cdc.MustUnmarshal(bz, &asset)

	return asset, true
}

func (k Keeper) GetAssetBySymbol(ctx sdk.Context, symbol string) (types.Asset, bool) {
	store := ctx.KVStore(k.storeKey)
	keySymbol := types.GetAssetBySymbolKey(symbol)

	bz := store.Get(keySymbol)
	if bz == nil {
		return types.Asset{}, false
	}

	var asset types.Asset
	k.cdc.MustUnmarshal(bz, &asset)

	return asset, true
}

func (k Keeper) DeleteAsset(ctx sdk.Context, denom string) error {
	asset, found := k.GetAssetByDenom(ctx, denom)
	if !found {
		return types.ErrorAssetDenomNotFound
	}

	store := ctx.KVStore(k.storeKey)
	keyDenom := types.GetAssetByDenomKey(asset.Denom)
	keySymbol := types.GetAssetBySymbolKey(asset.Symbol)

	store.Delete(keyDenom)
	store.Delete(keySymbol)

	return nil
}
