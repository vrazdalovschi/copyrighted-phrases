package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
)

// Keeper of the copyrightedphrases store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a copyrightedphrases keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// CreateCopyrightedText register the copyrighted text.
func (k Keeper) CreateCopyrightedText(ctx sdk.Context, copyrightedText types.Texts) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.CopyrightedTextPrefix + copyrightedText.Value)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(copyrightedText)
	store.Set(key, value)
}

// GetCopyrightedText returns the copyrighted text information
func (k Keeper) GetCopyrightedText(ctx sdk.Context, key string) (types.Texts, error) {
	store := ctx.KVStore(k.storeKey)
	var copyrightedText types.Texts
	byteKey := []byte(types.CopyrightedTextPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &copyrightedText)
	if err != nil {
		return copyrightedText, err
	}
	return copyrightedText, nil
}

// SetCopyrightedText sets a copyrightedText
func (k Keeper) SetCopyrightedText(ctx sdk.Context, name string, copyrightedText types.Texts) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(copyrightedText)
	key := []byte(types.CopyrightedTextPrefix + name)
	store.Set(key, bz)
}

// DeleteCopyrightedText deletes a copyrighted text
func (k Keeper) DeleteCopyrightedText(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.CopyrightedTextPrefix + key))
}

//
// Functions used by querier
//

func listCopyrightedText(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	key := path[0]
	var copyrightedTextList []types.Texts
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.CopyrightedTextPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var copyrightedText types.Texts
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &copyrightedText)
		if copyrightedText.Owner.String() == key {
			copyrightedTextList = append(copyrightedTextList, copyrightedText)
		}
	}
	res := codec.MustMarshalJSONIndent(k.cdc, copyrightedTextList)
	return res, nil
}

func getCopyrightedText(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	copyrightedText, err := k.GetCopyrightedText(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, copyrightedText)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get owner of the item
func (k Keeper) GetOwner(ctx sdk.Context, key string) sdk.AccAddress {
	copyrightedText, _ := k.GetCopyrightedText(ctx, key)
	return copyrightedText.Owner
}

// Check if the key exists in the store
func (k Keeper) Exists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.CopyrightedTextPrefix + key))
}

// Get an iterator over all names in which the keys are the names and the values are the copyrightedText
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.CopyrightedTextPrefix))
}
