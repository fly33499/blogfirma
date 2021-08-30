package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateHello{}

func NewMsgCreateHello(creator string, title string, body string) *MsgCreateHello {
	return &MsgCreateHello{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreateHello) Route() string {
	return RouterKey
}

func (msg *MsgCreateHello) Type() string {
	return "CreateHello"
}

func (msg *MsgCreateHello) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHello) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHello) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateHello{}

func NewMsgUpdateHello(creator string, id uint64, title string, body string) *MsgUpdateHello {
	return &MsgUpdateHello{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdateHello) Route() string {
	return RouterKey
}

func (msg *MsgUpdateHello) Type() string {
	return "UpdateHello"
}

func (msg *MsgUpdateHello) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateHello) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateHello) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteHello{}

func NewMsgDeleteHello(creator string, id uint64) *MsgDeleteHello {
	return &MsgDeleteHello{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteHello) Route() string {
	return RouterKey
}

func (msg *MsgDeleteHello) Type() string {
	return "DeleteHello"
}

func (msg *MsgDeleteHello) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteHello) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteHello) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
