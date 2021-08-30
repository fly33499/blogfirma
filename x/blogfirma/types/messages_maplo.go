package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMaplo{}

func NewMsgCreateMaplo(creator string, index string, title string, body string) *MsgCreateMaplo {
	return &MsgCreateMaplo{
		Creator: creator,
		Index:   index,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreateMaplo) Route() string {
	return RouterKey
}

func (msg *MsgCreateMaplo) Type() string {
	return "CreateMaplo"
}

func (msg *MsgCreateMaplo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMaplo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMaplo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMaplo{}

func NewMsgUpdateMaplo(creator string, index string, title string, body string) *MsgUpdateMaplo {
	return &MsgUpdateMaplo{
		Creator: creator,
		Index:   index,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdateMaplo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMaplo) Type() string {
	return "UpdateMaplo"
}

func (msg *MsgUpdateMaplo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMaplo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMaplo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMaplo{}

func NewMsgDeleteMaplo(creator string, index string) *MsgDeleteMaplo {
	return &MsgDeleteMaplo{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteMaplo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMaplo) Type() string {
	return "DeleteMaplo"
}

func (msg *MsgDeleteMaplo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMaplo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMaplo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
