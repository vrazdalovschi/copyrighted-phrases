package copyrightedphrases_test

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/keeper"
	ctypes "github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
	"testing"
)

func TestHandleMsgRegisterText(t *testing.T) {
	textKey := "CreateText"
	t.Run("create success", func(t *testing.T) {
		ctx, k := keeper.CreateTestInput(t, false)
		account := createFakeAccount(ctx, k.AccountKeeper)
		require.False(t, k.CopyrightedPhraseKeeper.Exists(ctx, textKey))

		h := copyrightedphrases.NewHandler(k.CopyrightedPhraseKeeper)
		res, err := h(ctx, ctypes.MsgRegisterText{Text: textKey, Owner: account})
		require.NoError(t, err)
		require.Equal(t, &sdk.Result{}, res)
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, textKey))
	})
	t.Run("create failed, already exists", func(t *testing.T) {
		ctx, k := keeper.CreateTestInput(t, false)
		account := createFakeAccount(ctx, k.AccountKeeper)
		k.CopyrightedPhraseKeeper.CreateCopyrightedText(ctx, ctypes.Texts{Value: textKey, Owner: account})
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, textKey))

		h := copyrightedphrases.NewHandler(k.CopyrightedPhraseKeeper)
		res, err := h(ctx, ctypes.MsgRegisterText{Text: textKey, Owner: account})
		require.Error(t, err)
		require.Nil(t, res)
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, textKey))
	})
}

func TestHandler_MsgDeleteText(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		ctx, k := keeper.CreateTestInput(t, false)
		account := createFakeAccount(ctx, k.AccountKeeper)
		k.CopyrightedPhraseKeeper.CreateCopyrightedText(ctx, ctypes.Texts{Value: "MyText", Owner: account})
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))

		h := copyrightedphrases.NewHandler(k.CopyrightedPhraseKeeper)
		res, err := h(ctx, ctypes.MsgDeleteText{Text: "MyText", Owner: account})
		require.NoError(t, err)
		require.Equal(t, &sdk.Result{}, res)
		require.False(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))
	})
	t.Run("delete fail different accounts", func(t *testing.T) {
		ctx, k := keeper.CreateTestInput(t, false)
		k.CopyrightedPhraseKeeper.CreateCopyrightedText(ctx, ctypes.Texts{Value: "MyText", Owner: createFakeAccount(ctx, k.AccountKeeper)})
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))

		h := copyrightedphrases.NewHandler(k.CopyrightedPhraseKeeper)

		res, err := h(ctx, ctypes.MsgDeleteText{Text: "MyText", Owner: createFakeAccount(ctx, k.AccountKeeper)})
		require.Error(t, err)
		require.Nil(t, res)
		require.True(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))
	})
	t.Run("delete not existing", func(t *testing.T) {
		ctx, k := keeper.CreateTestInput(t, false)
		require.False(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))
		h := copyrightedphrases.NewHandler(k.CopyrightedPhraseKeeper)

		res, err := h(ctx, ctypes.MsgDeleteText{Text: "MyText", Owner: createFakeAccount(ctx, k.AccountKeeper)})
		require.Error(t, err)
		require.Nil(t, res)
		require.False(t, k.CopyrightedPhraseKeeper.Exists(ctx, "MyText"))
	})
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
