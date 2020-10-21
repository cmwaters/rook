package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgMove{}

func NewMsgMove(creator sdk.AccAddress, quantity string, from string, to string) *MsgMove {
  return &MsgMove{
    Id: uuid.New().String(),
		Creator: creator,
    Quantity: quantity,
    From: from,
    To: to,
	}
}

func (msg *MsgMove) Route() string {
  return RouterKey
}

func (msg *MsgMove) Type() string {
  return "CreateMove"
}

func (msg *MsgMove) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgMove) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgMove) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
