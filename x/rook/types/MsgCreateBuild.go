package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgBuild{}

func NewMsgBuild(creator sdk.AccAddress, settlement string, position string) *MsgBuild {
  return &MsgBuild{
    Id: uuid.New().String(),
		Creator: creator,
    Settlement: settlement,
    Position: position,
	}
}

func (msg *MsgBuild) Route() string {
  return RouterKey
}

func (msg *MsgBuild) Type() string {
  return "CreateBuild"
}

func (msg *MsgBuild) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgBuild) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgBuild) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
