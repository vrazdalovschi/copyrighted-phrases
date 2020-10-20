package copyrightedphrases

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/keeper"
	"github.com/vrazdalovschi/copyrighted-phrases/x/copyrightedphrases/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgRegisterText:
			return handleMsgRegisterText(ctx, k, msg)
		case types.MsgDeleteText:
			return handleMsgDeleteText(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// Handle a message to delete name
func handleMsgDeleteText(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgDeleteText) (*sdk.Result, error) {
	if !keeper.Exists(ctx, msg.Text) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Text)
	}
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Text)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	keeper.DeleteCopyrightedText(ctx, msg.Text)
	return &sdk.Result{}, nil
}

// Handle a message to register text
func handleMsgRegisterText(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgRegisterText) (*sdk.Result, error) {
	if keeper.Exists(ctx, msg.Text) {
		keeper.Logger(ctx).Error(fmt.Sprintf("Text %s is already copyrighted.", msg.Text))
		return nil, sdkerrors.Wrap(types.ErrAlreadyCopyrighted, msg.Text)
	}
	keeper.Logger(ctx).Debug(fmt.Sprintf("Register text %s for owner %v", msg.Text, msg.Owner))

	keeper.CreateCopyrightedText(ctx, types.Texts{Value: msg.Text, Owner: msg.Owner, Block: ctx.BlockHeight()})
	return &sdk.Result{}, nil
}
