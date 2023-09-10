package types

// DefaultGenesis returns module's default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: Params{},
	}
}

// Validate performs basic genesis state validation, returning an error upon any failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
