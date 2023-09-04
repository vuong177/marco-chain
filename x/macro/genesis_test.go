package macro_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vuong177/macro/testutil/keeper"
	"github.com/vuong177/macro/testutil/nullify"
	"github.com/vuong177/macro/x/macro"
	"github.com/vuong177/macro/x/macro/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MacroKeeper(t)
	macro.InitGenesis(ctx, *k, genesisState)
	got := macro.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
