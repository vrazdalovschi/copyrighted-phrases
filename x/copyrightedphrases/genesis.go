package copyrightedphrases

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/keeper"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	for _, record := range data.CopyrightedTextRecords {
		keeper.SetCopyrightedText(ctx, record.Value, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	var records []types.Texts
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		copyrightedText, _ := k.GetCopyrightedText(ctx, name)
		records = append(records, copyrightedText)

	}
	return types.GenesisState{CopyrightedTextRecords: records}
}
