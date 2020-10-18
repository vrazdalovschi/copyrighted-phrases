package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	copyrightedphrasesTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	copyrightedphrasesTxCmd.AddCommand(flags.PostCommands(
		GetCmdRegisterText(cdc),
		GetCmdDeleteCopyrightedText(cdc),
	)...)

	return copyrightedphrasesTxCmd
}
