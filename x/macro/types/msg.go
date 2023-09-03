package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgMintStable      = "mint_stable_coin"
	TypeWithdrawCollateral = "withdraw_collateral"
	TypeMsgDeposit         = "deposit"
)

var _ sdk.Msg = &MsgMintStableCoin{}

// MsgMintStable creates a message to mint stable coin
func NewMsgMintStable(minter string, requestAmount sdk.Int) *MsgMintStableCoin {
	return &MsgMintStableCoin{
		Minter:        minter,
		RequestAmount: requestAmount,
	}
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

// MsgMintStable creates a message to mint stable coin
func NewMsgDeposit(fromAdress string, coin sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		FromAddress: fromAdress,
		DepositCoin: coin,
	}
}

func (m MsgDeposit) Route() string { return RouterKey }
func (m MsgDeposit) Type() string  { return TypeMsgDeposit }

func (m MsgDeposit) ValidateBasic() error {
	return nil
}

func (m MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgDeposit) GetSigners() []sdk.AccAddress {
	from, _ := sdk.AccAddressFromBech32(m.FromAddress)
	return []sdk.AccAddress{from}
}
