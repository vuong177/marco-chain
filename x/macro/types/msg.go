package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgMintStable      = "mint_stable_coin"
	TypeWithdrawCollateral = "withdraw_collateral"
)

var _ sdk.Msg = &MsgMintStableCoin{}

// MsgMintStable creates a message to mint stable coin
func NewMsgMintStable() *MsgMintStableCoin {
	return &MsgMintStableCoin{}
}

func (m MsgMintStableCoin) Route() string { return RouterKey }
func (m MsgMintStableCoin) Type() string  { return TypeMsgMintStable }
func (m MsgMintStableCoin) ValidateBasic() error {
	return nil
}

func (m MsgMintStableCoin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgMintStableCoin) GetSigners() []sdk.AccAddress {
	minter, _ := sdk.AccAddressFromBech32(m.Minter)
	return []sdk.AccAddress{minter}
}

var _ sdk.Msg = &MsgWithdrawCollateral{}

// MsgMintStable creates a message to mint stable coin
func NewWithdrawCollateral() *MsgWithdrawCollateral {
	return &MsgWithdrawCollateral{}
}

func (m MsgWithdrawCollateral) Route() string { return RouterKey }
func (m MsgWithdrawCollateral) Type() string  { return TypeMsgMintStable }

func (m MsgWithdrawCollateral) ValidateBasic() error {
	return nil
}

func (m MsgWithdrawCollateral) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgWithdrawCollateral) GetSigners() []sdk.AccAddress {
	minter, _ := sdk.AccAddressFromBech32(m.Minter)
	return []sdk.AccAddress{minter}
}
