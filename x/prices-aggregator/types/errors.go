package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorUnknownProposalType             = errors.Register(ModuleName, 1, "unknown proposal type")
	ErrRequestIDNotAvailable             = errors.Register(ModuleName, 2, "Request ID not available")
	ErrInvalidVersion                    = errors.Register(ModuleName, 3, "invalid version")
	ErrUnrecognizePacket                 = errors.Register(ModuleName, 4, "Unrecognised packet")
	ErrorDuplicateAsset                  = errors.Register(ModuleName, 5, "err duplicate asset")
	ErrorAssetNotFound                   = errors.Register(ModuleName, 6, "err asset not found")
	ErrorAssetDenomNotFound              = errors.Register(ModuleName, 7, "err asset asset-denom not found")
	ErrorAssetSymbolNotFound             = errors.Register(ModuleName, 8, "err asset asset-symbol not found")
	ErrorOracleRequestNotFound           = errors.Register(ModuleName, 9, "oracle request not found")
	ErrorInvalidOracleResponse           = errors.Register(ModuleName, 10, "invalid oracle response")
	ErrorInvalidOracleResponseNotSuccess = errors.Register(ModuleName, 11, "oracle response not success")
	ErrorUpdatedTimeFromPast             = errors.Register(ModuleName, 12, "err asset asset-symbol not found")
)
