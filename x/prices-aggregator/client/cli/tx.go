package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vuong177/macro/x/prices-aggregator/types"
)

func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "transfermiddleware",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		Short:                      fmt.Sprintf("Tx commands for the %s module", types.ModuleName),
	}

	// TODO: Add CLI Tx
	txCmd.AddCommand()

	return txCmd
}
