package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
)

func GetCmdListCopyrightedTexts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-copyrighted-texts [address]",
		Short: "list all copyrighted texts for a given sdk.AccAddress",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListCopyrightedText, key), nil)
			if err != nil {
				fmt.Printf("could not list Texts for account %s\n%s\n", key, err.Error())
				return nil
			}
			var out []types.Texts
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetCopyrightedText(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-copyrighted-text [text]",
		Short: "Query a copyrighted text by text",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetCopyrightedText, key), nil)
			if err != nil {
				fmt.Printf("could not resolve copyrighted text %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Texts
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
