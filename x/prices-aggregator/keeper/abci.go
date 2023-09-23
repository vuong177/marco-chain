package keeper

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

// BeginBlocker of prices-aggregator module.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, ctx.BlockTime(), telemetry.MetricKeyBeginBlocker)
	if ctx.BlockHeight()%types.Int64Five == types.Int64Zero {
		// Send Oracle request
		k.handleSendOracleRequest(ctx)
	}
}
