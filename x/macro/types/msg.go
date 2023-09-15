package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgMintStable      = "mint_stable_coin"
	TypeWithdrawCollateral = "withdraw_collateral"
	TypeMsgDeposit         = "deposit"
	TypeMsgRepay 		   = "repay"
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

var _ sdk.Msg = &MsgRepay{}

// MsgRepay creates a message to mint stable coin
func NewMsgRepay(repayer string, borrower string, amount sdkmath.LegacyDec) *MsgRepay {
	return &MsgRepay{
		repayer, borrower, amount,
	}
}

func (m MsgRepay) Route() string { return RouterKey }
func (m MsgRepay) Type() string  { return TypeMsgRepay }

func (m MsgRepay) ValidateBasic() error {
	return nil
}

func (m MsgRepay) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgRepay) GetSigners() []sdk.AccAddress {
	minter, _ := sdk.AccAddressFromBech32(m.Repayer)
	return []sdk.AccAddress{minter}
}

var _ sdk.Msg = &MsgBecomeRedemptionProvider{}

// NewMsgBecomeRedemptionProvider creates a message to mint stable coin
func NewMsgBecomeRedemptionProvider(redemption_provider string) *MsgBecomeRedemptionProvider {
	return &MsgBecomeRedemptionProvider{
		redemption_provider,
	}
}

func (m MsgBecomeRedemptionProvider) Route() string { return RouterKey }
func (m MsgBecomeRedemptionProvider) Type() string  { return TypeMsgRepay }

func (m MsgBecomeRedemptionProvider) ValidateBasic() error {
	return nil
}

func (m MsgBecomeRedemptionProvider) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgBecomeRedemptionProvider) GetSigners() []sdk.AccAddress {
	minter, _ := sdk.AccAddressFromBech32(m.RedemptionProvider)
	return []sdk.AccAddress{minter}
}