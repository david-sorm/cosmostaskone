package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddTokensLock{}

func NewMsgAddTokensLock(creator string, amount string, denom string, address string) *MsgAddTokensLock {
	return &MsgAddTokensLock{
		Creator: creator,
		Amount:  amount,
		Denom:   denom,
		Address: address,
	}
}

func (msg *MsgAddTokensLock) Route() string {
	return RouterKey
}

func (msg *MsgAddTokensLock) Type() string {
	return "AddTokensLock"
}

func (msg *MsgAddTokensLock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddTokensLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTokensLock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
