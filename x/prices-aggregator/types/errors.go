package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorUnknownProposalType = errors.Register(ModuleName, 1, "unknown proposal type")
	ErrRequestIDNotAvailable = errors.Register(ModuleName, 2, "Request ID not available")
	ErrInvalidVersion        = errors.Register(ModuleName, 3, "invalid version")
	ErrUnrecognizePacket     = errors.Register(ModuleName, 4, "Unrecognised packet")
	ErrorDuplicateAsset      = errors.Register(ModuleName, 5, "err duplicate asset")
	ErrorAssetDenomNotFound  = errors.Register(ModuleName, 6, "err asset asset denom not found")
)
