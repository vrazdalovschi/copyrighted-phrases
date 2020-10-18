package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
)

// NewQuerier creates a new querier for copyrightedphrases clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryGetCopyrightedText:
			return getCopyrightedText(ctx, path[1:], k)
		case types.QueryListCopyrightedText:
			return listCopyrightedText(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown copyrightedphrases query endpoint")
		}
	}
}
