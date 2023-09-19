package keeper_test

import (
	"fmt"
	"testing"

	tmrand "github.com/cometbft/cometbft/libs/rand"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/stretchr/testify/suite"
	"github.com/vuong177/macro/app"
	"github.com/vuong177/macro/x/macro/keeper"
	"github.com/vuong177/macro/x/macro/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx         sdk.Context
	macroApp    *app.App
	macroKeeper keeper.Keeper
	govKeeper   govkeeper.Keeper
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.macroApp = app.Setup(suite.T())
	suite.ctx = suite.macroApp.BaseApp.NewContext(false, tmproto.Header{
		ChainID: fmt.Sprintf("test-chain-%s", tmrand.Str(4)),
		Height:  1,
	})
	suite.macroKeeper = suite.macroApp.MacroKeeper
	suite.govKeeper = suite.macroApp.GovKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.macroApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, keeper.NewQuerier(suite.macroKeeper))
	suite.queryClient = types.NewQueryClient(queryHelper)

	suite.msgServer = keeper.NewMsgServerImpl(suite.macroKeeper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
