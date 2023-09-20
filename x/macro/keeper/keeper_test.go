package keeper_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto/ed25519"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
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
	testAccs    []sdk.AccAddress
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
	suite.testAccs = CreateRandomAccounts(3)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := ed25519.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

// Mints and sends coins to a user account
func (suite *KeeperTestSuite) FundAccount(acc sdk.AccAddress, amount sdk.Coin) {
	amountCoins := sdk.NewCoins(amount)
	err := suite.macroApp.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, amountCoins)
	suite.Require().NoError(err)
	err = suite.macroApp.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, acc, amountCoins)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestHandleRepay() {
	testcases := []struct {
		name               string
		notSetDataBorrower bool
		amountBorrowed     sdkmath.LegacyDec
		amountRepay        sdkmath.LegacyDec
		balanceOfRepayer   sdkmath.Int
		expectedErr        bool
	}{
		{
			name:               "Err: can not find collateral data of borrower",
			notSetDataBorrower: true,
			balanceOfRepayer:   sdkmath.NewIntFromUint64(100),
			expectedErr:        true,
		},
	}
	for _, tc := range testcases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			repayerAddress, borrowerAddress := suite.testAccs[0], suite.testAccs[1]
			if !tc.notSetDataBorrower {
				borrowerData := types.BorrowerData{
					Address:  borrowerAddress.String(),
					Borrowed: tc.amountBorrowed,
				}
				suite.macroApp.MacroKeeper.SetBorrowerData(suite.ctx, borrowerAddress, borrowerData)
			}
			suite.FundAccount(repayerAddress, sdk.NewCoin(types.StableCoinDenom, sdkmath.Int(tc.balanceOfRepayer)))
			err := suite.macroApp.MacroKeeper.HandleRepay(suite.ctx, repayerAddress, borrowerAddress, tc.amountRepay)
			if tc.expectedErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
