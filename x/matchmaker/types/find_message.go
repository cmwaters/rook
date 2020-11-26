package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgFind{}

func NewMsgFind(creator sdk.AccAddress, gameId string) *MsgFind {
	return &MsgFind{
		Id:      uuid.New().String(),
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgFind) Route() string {
	return RouterKey
}

func (msg *MsgFind) Type() string {
	return "Find"
}

func (msg *MsgFind) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgFind) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFind) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
