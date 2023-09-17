package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeAddAssetPricesTrackingList = "add_asset_prices_tracking"
)

var _ sdk.Msg = &MsgAddAssetPricesTrackingList{}

// Route Implements Msg.
func (msg MsgAddAssetPricesTrackingList) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgAddAssetPricesTrackingList) Type() string { return TypeAddAssetPricesTrackingList }

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgAddAssetPricesTrackingList) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgAddParachainIBCTokenInfo message.
func (msg *MsgAddAssetPricesTrackingList) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *MsgAddAssetPricesTrackingList) ValidateBasic() error {
	// validate authority
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrap(err, "invalid authority address")
	}

	return nil
}
