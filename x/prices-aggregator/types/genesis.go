package types

import (
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// DefaultGenesis returns module's default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{},
		PortId: PortID,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	return gs.Params.Validate()
}
