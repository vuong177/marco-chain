package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgCreateGauge = "create_gauge"
	TypeMsgAddToGauge  = "add_to_gauge"
)

var _ sdk.Msg = &MsgMintStable{}

// MsgMintStable creates a message to mint stable coin
func NewMsgMintStable() *MsgMintStable {
	return &MsgMintStable{}
}

func (m MsgMintStable) Route() string { return RouterKey }
func (m MsgMintStable) Type() string  { return TypeMsgCreateGauge }
func (m MsgMintStable) ValidateBasic() error {
	return nil
}

func (m MsgMintStable) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgMintStable) GetSigners() []sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{owner}
}
