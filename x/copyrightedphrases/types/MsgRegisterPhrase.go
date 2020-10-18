package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgRegisterText defines the RegisterText message
type MsgRegisterText struct {
	Text  string         `json:"text"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgRegisterText is the constructor function for MsgRegisterText
func NewMsgRegisterText(text string, owner sdk.AccAddress) MsgRegisterText {
	return MsgRegisterText{
		Text:  text,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgRegisterText) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterText) Type() string { return "register_text" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterText) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Owner.String())
	}
	if len(msg.Text) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Text cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterText) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterText) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
