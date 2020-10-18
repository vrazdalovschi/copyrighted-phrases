package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
	"testing"
)

// Test to validate full flow: Add, Get, Exists, Delete, GetOwner
func TestKeeper(t *testing.T) {
	ctx, keepers := CreateTestInput(t, false)
	accKeeper, keeper := keepers.AccountKeeper, keepers.CopyrightedPhraseKeeper
	creator := createFakeAccount(ctx, accKeeper)

	textValue := "My_Secret_text"

	//Check doesn't exists
	res := keeper.Exists(ctx, textValue)
	require.False(t, res)

	//Create
	keeper.CreateCopyrightedText(ctx, types.Texts{Value: textValue, Owner: creator})

	//Check created
	res = keeper.Exists(ctx, textValue)
	require.True(t, res)

	//Check valid
	r, err := keeper.GetCopyrightedText(ctx, textValue)
	require.NoError(t, err)
	require.Equal(t, textValue, r.Value)
	require.Equal(t, creator, r.Owner)

	//Check owner
	owner := keeper.GetOwner(ctx, textValue)
	require.Equal(t, creator, owner)

	//Check delete
	keeper.DeleteCopyrightedText(ctx, textValue)

	//Check doesn't exists
	res = keeper.Exists(ctx, textValue)
	require.False(t, res)
}

func createFakeAccount(ctx sdk.Context, am auth.AccountKeeper) sdk.AccAddress {
	_, _, addr := keyPubAddr()
	fundAccounts(ctx, am, addr, sdk.NewCoins(sdk.NewInt64Coin("denom", 100000)))
	return addr
}

func fundAccounts(ctx sdk.Context, am auth.AccountKeeper, addr sdk.AccAddress, coins sdk.Coins) {
	baseAcct := auth.NewBaseAccountWithAddress(addr)
	_ = baseAcct.SetCoins(coins)
	am.SetAccount(ctx, &baseAcct)
}

var keyCounter uint64 = 0

// we need to make this deterministic (same every test run), as encoded address size and thus gas cost,
// depends on the actual bytes (due to ugly CanonicalAddress encoding)
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	keyCounter++
	seed := make([]byte, 8)
	binary.BigEndian.PutUint64(seed, keyCounter)

	key := ed25519.GenPrivKeyFromSecret(seed)
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}
