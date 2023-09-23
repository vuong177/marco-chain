package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys store keys.
var (
	Int64Zero = int64(0)
	Int64Five = int64(5)

	KeyAskCount   = []byte("AskCount")
	KeyMinCount   = []byte("MinCount")
	KeyFeeLimit   = []byte("FeeLimit")
	KeyPrepareGas = []byte("PrepareGas")
	KeyExecuteGas = []byte("ExecuteGas")
	KeyChannelID  = []byte("ChannelID")
)

var _ paramtypes.ParamSet = &Params{}

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyAskCount, &p.AskCount, validateUint64),
		paramtypes.NewParamSetPair(KeyMinCount, &p.MinCount, validateUint64),
		paramtypes.NewParamSetPair(KeyFeeLimit, &p.FeeLimit, validateFeeLimit),
		paramtypes.NewParamSetPair(KeyPrepareGas, &p.PrepareGas, validateUint64),
		paramtypes.NewParamSetPair(KeyExecuteGas, &p.ExecuteGas, validateUint64),
		paramtypes.NewParamSetPair(KeyChannelID, &p.ChannelId, validateString),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateUint64(p.AskCount); err != nil {
		return err
	}
	if err := validateUint64(p.MinCount); err != nil {
		return err
	}
	if err := validateUint64(p.PrepareGas); err != nil {
		return err
	}
	if err := validateUint64(p.ExecuteGas); err != nil {
		return err
	}
	if err := validateFeeLimit(p.FeeLimit); err != nil {
		return err
	}
	if err := validateString(p.ChannelId); err != nil {
		return err
	}
	return nil
}

func validateUint64(i interface{}) error {
	if _, ok := i.(uint64); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateString(i interface{}) error {
	if _, ok := i.(string); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateFeeLimit(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}
