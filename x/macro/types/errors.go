package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/macro module sentinel errors
var (
	ErrSample                     = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSmallRequestCollateralRate = sdkerrors.Register(ModuleName, 1, "request collateral rate too small")
	ErrEmptyDepositAsset          = sdkerrors.Register(ModuleName, 2, "empty deposit asset")
	ErrEmptyMintedStableCoin      = sdkerrors.Register(ModuleName, 3, "empty minted stable coin")
	ErrCanNotFindCollateralData   = sdkerrors.Register(ModuleName, 4, "can not find data of user")
)
