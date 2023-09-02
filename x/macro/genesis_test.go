package marco_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vuong177/macro/testutil/keeper"
	"github.com/vuong177/macro/testutil/nullify"
	"github.com/vuong177/macro/x/marco"
	"github.com/vuong177/macro/x/marco/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarcoKeeper(t)
	marco.InitGenesis(ctx, *k, genesisState)
	got := marco.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
