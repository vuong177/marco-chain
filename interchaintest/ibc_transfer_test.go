package interchaintest

import (
	"context"
	"fmt"
	"testing"

	"cosmossdk.io/math"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v7/relayer"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestHyperspace features
// * sets up a Polkadot parachain
// * sets up a Cosmos chain
// * sets up the Hyperspace relayer
// * Funds a user wallet on both chains
// * Pushes a wasm client contract to the Cosmos chain
// * create client, connection, and channel in relayer
// * start relayer
// * send transfer over ibc
func TestMacroGaiaIBCTransfer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctx := context.Background()

	// Create chain factory with Macro and Gaia
	numVals := 1
	numFullNodes := 1

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          "Macro",
			ChainConfig:   macroConfig,
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "gaia",
			Version:       "v9.1.0",
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	macro, gaia := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)
	pathMacroGaia := "macro-gaia"
	// Create relayer factory to utilize the go-relayer
	client, network := interchaintest.DockerSetup(t)

	r := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		interchaintestrelayer.CustomDockerImage(IBCRelayerImage, IBCRelayerVersion, "100:1000"),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100")).Build(t, client, network)

	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain().
		AddChain(macro).
		AddChain(gaia).
		AddRelayer(r, "rly").
		AddLink(interchaintest.InterchainLink{
			Chain1:  macro,
			Chain2:  gaia,
			Relayer: r,
			Path:    pathMacroGaia,
		})

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	err = ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: false,

		// This can be used to write to the block database which will index all block data e.g. txs, msgs, events, etc.
		// BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Start the relayer
	require.NoError(t, r.StartRelayer(ctx, eRep, pathMacroGaia))
	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occurred while stopping the relayer: %s", err))
			}
		},
	)
	genesisWalletAmount := int64(10_000_000)
	// Create some user accounts on both chains
	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), genesisWalletAmount, macro, gaia)

	// Wait a few blocks for relayer to start and for user accounts to be created
	err = testutil.WaitForBlocks(ctx, 5, macro, gaia)
	require.NoError(t, err)

	// Get our Bech32 encoded user addresses
	macroUser, gaiaUser := users[0], users[1]

	macroUserAddr := macroUser.FormattedAddress()
	gaiaUserAddr := gaiaUser.FormattedAddress()

	// Get original account balances
	macroOrigBal, err := macro.GetBalance(ctx, macroUserAddr, macro.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, macroOrigBal)

	gaiaOrigBal, err := gaia.GetBalance(ctx, gaiaUserAddr, gaia.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, genesisWalletAmount, gaiaOrigBal)

	// Compose an IBC transfer and send from macro -> Gaia
	transferAmount := math.NewInt(1000)
	transfer := ibc.WalletAmount{
		Address: gaiaUserAddr,
		Denom:   macro.Config().Denom,
		Amount:  transferAmount,
	}

	channel, err := ibc.GetTransferChannel(ctx, r, eRep, macro.Config().ChainID, gaia.Config().ChainID)
	require.NoError(t, err)

	transferTx, err := macro.SendIBCTransfer(ctx, channel.ChannelID, macroUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	macroHeight, err := macro.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know the transfer was successful
	_, err = testutil.PollForAck(ctx, macro, macroHeight, macroHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Get the IBC denom for stake on Gaia
	macroTokenDenom := transfertypes.GetPrefixedDenom(channel.Counterparty.PortID, channel.Counterparty.ChannelID, macro.Config().Denom)
	macroIBCDenom := transfertypes.ParseDenomTrace(macroTokenDenom).IBCDenom()

	// Assert that the funds are no longer present in user acc on macro and are in the user acc on Gaia
	macroUpdateBal, err := macro.GetBalance(ctx, macroUserAddr, macro.Config().Denom)
	require.NoError(t, err)
	require.True(t, macroUpdateBal.Equal(macroOrigBal.Sub(transferAmount)))

	gaiaUpdateBal, err := gaia.GetBalance(ctx, gaiaUserAddr, macroIBCDenom)
	require.NoError(t, err)
	require.Equal(t, transferAmount, gaiaUpdateBal)

	// Compose an IBC transfer and send from Gaia -> macro
	transfer = ibc.WalletAmount{
		Address: macroUserAddr,
		Denom:   macroIBCDenom,
		Amount:  transferAmount,
	}

	transferTx, err = gaia.SendIBCTransfer(ctx, channel.Counterparty.ChannelID, gaiaUserAddr, transfer, ibc.TransferOptions{})
	require.NoError(t, err)

	gaiaHeight, err := gaia.Height(ctx)
	require.NoError(t, err)

	// Poll for the ack to know the transfer was successful
	_, err = testutil.PollForAck(ctx, gaia, gaiaHeight, gaiaHeight+10, transferTx.Packet)
	require.NoError(t, err)

	// Assert that the funds are now back on macro and not on Gaia
	macroUpdateBal, err = macro.GetBalance(ctx, macroUserAddr, macro.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, macroOrigBal, macroUpdateBal)

	gaiaUpdateBal, err = gaia.GetBalance(ctx, gaiaUserAddr, macroIBCDenom)
	require.NoError(t, err)
	require.Equal(t, int64(0), gaiaUpdateBal)
}

func fundUsers(t *testing.T, ctx context.Context, fundAmount int64, composable, macrod ibc.Chain) (ibc.Wallet, ibc.Wallet) {
	users := interchaintest.GetAndFundTestUsers(t, ctx, "user", fundAmount, composable, macrod)
	polkadotUser, cosmosUser := users[0], users[1]
	err := testutil.WaitForBlocks(ctx, 2, composable, macrod) // Only waiting 1 block is flaky for parachain
	require.NoError(t, err, "cosmos or polkadot chain failed to make blocks")

	// Check balances are correct
	polkadotUserAmount, err := composable.GetBalance(ctx, polkadotUser.FormattedAddress(), composable.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundAmount, polkadotUserAmount, "Initial polkadot user amount not expected")
	parachainUserAmount, err := composable.GetBalance(ctx, polkadotUser.FormattedAddress(), "")
	require.NoError(t, err)
	require.Equal(t, fundAmount, parachainUserAmount, "Initial parachain user amount not expected")
	cosmosUserAmount, err := macrod.GetBalance(ctx, cosmosUser.FormattedAddress(), macrod.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundAmount, cosmosUserAmount, "Initial cosmos user amount not expected")

	return polkadotUser, cosmosUser
}
