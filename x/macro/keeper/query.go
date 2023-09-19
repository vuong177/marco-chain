package keeper

import (
	"github.com/vuong177/macro/x/macro/types"
)

var _ types.QueryServer = Keeper{}

// Querier defines a wrapper around the x/macro keeper providing gRPC method
// handlers.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}
