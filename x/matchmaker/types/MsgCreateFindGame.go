package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgFindGame{}

func NewMsgFindGame(creator sdk.AccAddress, game string, options string) *MsgFindGame {
	return &MsgFindGame{
		Id:      uuid.New().String(),
		Creator: creator,
		Game:    game,
		Options: options,
	}
}

func (msg *MsgFindGame) Route() string {
	return RouterKey
}

func (msg *MsgFindGame) Type() string {
	return "CreateFindGame"
}

func (msg *MsgFindGame) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgFindGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFindGame) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
