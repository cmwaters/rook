package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgInvite{}

func NewMsgInvite(creator sdk.AccAddress, players []sdk.AccAddress, config []byte, contractID uint64) *MsgInvite {
	return &MsgInvite{
		Id:      uuid.New().String(),
		Creator: creator,
		Players: players,
		Config:  config,
		ContractId: contractID, 
	}
}

func (msg *MsgInvite) Route() string {
	return RouterKey
}

func (msg *MsgInvite) Type() string {
	return "Invite"
}

func (msg *MsgInvite) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgInvite) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInvite) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.ContractId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "contract id is required")
	}
	return nil
}
