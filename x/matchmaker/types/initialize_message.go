package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgInitialize{}

func NewMsgInitialize(creator sdk.AccAddress, players []sdk.AccAddress, config *GameConfig) *MsgInitialize {
	return &MsgInitialize{
		Id:      uuid.New().String(),
		Creator: creator,
		Players: players,
		Config:  config,
	}
}

func (msg *MsgInitialize) Route() string {
	return RouterKey
}

func (msg *MsgInitialize) Type() string {
	return "Initialize"
}

func (msg *MsgInitialize) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgInitialize) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitialize) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
