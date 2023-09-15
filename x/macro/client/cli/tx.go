package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/vuong177/macro/x/macro/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(GetDepositCmd())
	cmd.AddCommand(GetMintStableCoinCmd())
	cmd.AddCommand(GetRepayCmd())
	cmd.AddCommand(GetBecomeRedemptionProviderCmd())

	return cmd
}

func GetDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit [coin]",
		Short:   "Deposit collateral asset ",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Example: fmt.Sprintf("%s tx macro deposit [coin]", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			coins, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddress := clientCtx.GetFromAddress().String()

			msg := types.NewMsgDeposit(
				fromAddress,
				coins,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetMintStableCoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "mint [amount]",
		Short:   "Mint new stable coin  ",
		Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Example: fmt.Sprintf("%s tx macro mint [amount]", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("can't parse uusd amount")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			minter := clientCtx.GetFromAddress().String()

			msg := types.NewMsgMintStable(
				minter,
				amount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetRepayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [borrower] [amount]",
		Short: "Repay debt to increase collateral ratio",
		Args:  cobra.ExactArgs(2),
		Example: fmt.Sprintf("%s tx macro repay [borrower] [amount]", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			borrower := args[0]
			amount, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return fmt.Errorf("can't parse uusd amount")
			}
			repayer := clientCtx.GetFromAddress().String()
			msg := types.NewMsgRepay(
				repayer,
				borrower,
				amount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetBecomeRedemptionProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "becomeredemptionprovider [amount]",
		Short: "Become a redemption provider",
		Args:  cobra.ExactArgs(0),
		Example: fmt.Sprintf("%s tx macro becomeredemptionprovider", version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			redemption_provider := clientCtx.GetFromAddress().String()
			msg := types.NewMsgBecomeRedemptionProvider(
				redemption_provider,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
