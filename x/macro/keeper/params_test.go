package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/vuong177/macro/testutil/keeper"
	"github.com/vuong177/macro/x/marco/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MarcoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
