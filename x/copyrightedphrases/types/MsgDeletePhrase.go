package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteText{}

type MsgDeleteText struct {
	Text  string         `json:"text" yaml:"text"`
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
}

func NewMsgDeleteCopyrightedText(text string, owner sdk.AccAddress) MsgDeleteText {
	return MsgDeleteText{
		Text:  text,
		Owner: owner,
	}
}

func (msg MsgDeleteText) Route() string {
	return RouterKey
}

func (msg MsgDeleteText) Type() string {
	return "DeleteText"
}

func (msg MsgDeleteText) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

func (msg MsgDeleteText) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteText) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
