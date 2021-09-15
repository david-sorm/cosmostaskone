package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveTokenLock{}

func NewMsgRemoveTokenLock(creator string, id string) *MsgRemoveTokenLock {
	return &MsgRemoveTokenLock{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRemoveTokenLock) Route() string {
	return RouterKey
}

func (msg *MsgRemoveTokenLock) Type() string {
	return "RemoveTokenLock"
}

func (msg *MsgRemoveTokenLock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveTokenLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveTokenLock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
