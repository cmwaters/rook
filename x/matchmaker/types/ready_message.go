package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgReady{}

func NewMsgReady(creator sdk.AccAddress, gameId string) *MsgReady {
	return &MsgReady{
		Id:      uuid.New().String(),
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgReady) Route() string {
	return RouterKey
}

func (msg *MsgReady) Type() string {
	return "Ready"
}

func (msg *MsgReady) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgReady) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReady) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
