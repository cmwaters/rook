package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgJoin{}

func NewMsgJoin(creator sdk.AccAddress, gameId string) *MsgJoin {
	return &MsgJoin{
		Id:      uuid.New().String(),
		Creator: creator,
		GameId:  gameId,
	}
}

func (msg *MsgJoin) Route() string {
	return RouterKey
}

func (msg *MsgJoin) Type() string {
	return "Join"
}

func (msg *MsgJoin) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgJoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJoin) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
