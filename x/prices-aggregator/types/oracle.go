package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewOracleRequestPacketData contructs a new OracleRequestPacketData instance
func NewOracleRequestPacketData(
	clientID string,
	oracleScriptID uint64,
	calldata []byte,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) OracleRequestPacketData {
	return OracleRequestPacketData{
		ClientID:       clientID,
		OracleScriptID: oracleScriptID,
		Calldata:       calldata,
		AskCount:       askCount,
		MinCount:       minCount,
		FeeLimit:       feeLimit,
		PrepareGas:     prepareGas,
		ExecuteGas:     executeGas,
	}
}

// GetBytes returns the JSON marshalled interchain query packet data.
func (p OracleRequestPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}
