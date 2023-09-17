package keeper

import (
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
	utils "github.com/vuong177/macro/x/prices-aggregator/utils"
)

// OnRecvPacket receive OracleResponsePacketData
func (k Keeper) OnRecvPacket(ctx sdk.Context, oracleResponse types.OracleResponsePacketData) error {
	// Get OracleRequestPacketData
	clientID, err := strconv.ParseUint(oracleResponse.ClientID, 10, 64)
	if err != nil {
		return err
	}

	// Get and delete oracle request
	oracleRequest, found := k.GetOracleRequestByClientID(ctx, clientID)
	if !found {
		return types.ErrorOracleRequestNotFound
	}
	k.DeleteOracleRequest(ctx, clientID)

	// Check if oracleResponse valid
	if oracleResponse.ResolveStatus != types.RESOLVE_STATUS_SUCCESS {
		return types.ErrorInvalidOracleResponseNotSuccess
	}

	var fetchPriceRequest types.FetchPriceRequest
	if err := utils.Decode(oracleRequest.GetCalldata(), &fetchPriceRequest); err != nil {
		return errorsmod.Wrap(errorstypes.ErrUnknownRequest, "cannot decode the fetchPriceRequest oracleRequest packet")
	}

	var fetchPriceResponse types.FetchPriceResponse
	if err := utils.Decode(oracleResponse.GetResult(), &fetchPriceResponse); err != nil {
		return errorsmod.Wrap(errorstypes.ErrUnknownRequest, "cannot decode the fetchPriceResponse oracleResponse packet")
	}

	if len(fetchPriceRequest.Symbols) != len(fetchPriceResponse.Rates) {
		return types.ErrorInvalidOracleResponse
	}

	for i, symbol := range fetchPriceRequest.Symbols {
		// get asset
		asset, found := k.GetAssetBySymbol(ctx, symbol)
		if !found {
			return types.ErrorAssetSymbolNotFound
		}
		// check time
		resolveTime := time.Unix(oracleResponse.ResolveTime, 0)
		if resolveTime.Before(asset.UpdateTime) {
			return types.ErrorUpdatedTimeFromPast
		}
		// update
		exchangeRates := sdk.NewDecFromInt(sdk.NewIntFromUint64(fetchPriceResponse.Rates[i])).QuoInt64(int64(fetchPriceRequest.Multiplier))
		asset.ExchangeRates = exchangeRates
		asset.UpdateTime = resolveTime

		k.SetAsset(ctx, asset)
	}

	return nil
}

// OnAcknowledgementPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain. If the acknowledgement
// failed then nothing occurs. If the acknowledgement success, then
// the sender store packet send information in store so that can retrieve later (when receive a OracleResponsePacketData)
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, ack channeltypes.Acknowledgement, packet channeltypes.Packet) error {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck types.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return err
		}

		var oracleRequest types.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(packet.GetData(), &oracleRequest); err != nil {
			return err
		}

		k.SetOracleRequestByClientID(ctx, oracleRequest)
		return nil
	default:
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	}
}
