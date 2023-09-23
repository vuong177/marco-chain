package keeper

import (
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
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

	// Check if oracleResponse succes
	if oracleResponse.ResolveStatus != types.RESOLVE_STATUS_SUCCESS {
		return types.ErrorInvalidOracleResponseNotSuccess
	}

	// Get fetch price request
	var fetchPriceRequest types.FetchPriceRequest
	if err := utils.Decode(oracleRequest.GetCalldata(), &fetchPriceRequest); err != nil {
		return errorsmod.Wrap(errorstypes.ErrUnknownRequest, "cannot decode the fetchPriceRequest oracleRequest packet")
	}

	// Get fetch price response
	var fetchPriceResponse types.FetchPriceResponse
	if err := utils.Decode(oracleResponse.GetResult(), &fetchPriceResponse); err != nil {
		return errorsmod.Wrap(errorstypes.ErrUnknownRequest, "cannot decode the fetchPriceResponse oracleResponse packet")
	}

	// Check response valid
	if len(fetchPriceRequest.Symbols) != len(fetchPriceResponse.Rates) {
		return types.ErrorInvalidOracleResponse
	}

	// update exchange rates
	for i, symbol := range fetchPriceRequest.Symbols {
		// get asset
		asset, found := k.GetAssetBySymbol(ctx, symbol)
		if !found {
			return types.ErrorAssetSymbolNotFound
		}
		// check time
		resolveTime := time.Unix(oracleResponse.ResolveTime, 0)
		if resolveTime.Before(asset.UpdateTime) {
			k.Logger(ctx).Error("Error resolveTime %v before UpdateTime %v", resolveTime, asset.UpdateTime)
			continue
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
		return nil
	}
}

// SendPacket send oracle request packet
func (k Keeper) SendPacket(ctx sdk.Context, oracleRequest types.OracleRequestPacketData, sourcePort string, sourceChannel string) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute * 5).UnixNano()

	sequence, err := k.channelKeeper.SendPacket(
		ctx,
		channelCap,
		sourcePort,
		sourceChannel,
		clienttypes.ZeroHeight(),
		uint64(timeoutTimestamp),
		oracleRequest.GetBytes(),
	)
	if err != nil {
		return 0, err
	}

	return sequence, nil
}

func (k Keeper) handleSendOracleRequest(ctx sdk.Context) (uint64, error) {
	params := k.GetParams(ctx)
	// Get oracleRequest ClientID
	clientID := k.GetNextClientID(ctx)
	// Create calldata
	var symbol []string
	k.IterateAssetList(ctx, func(asset types.Asset) (stop bool) {
		symbol = append(symbol, asset.Symbol)
		return false
	})
	encodedCallData := utils.MustEncode(types.FetchPriceRequest{Symbols: symbol, Multiplier: 1000000})
	// Create OracleRequest packet
	oracleRequest := types.OracleRequestPacketData{
		ClientID:       strconv.FormatUint(clientID, 10),
		OracleScriptID: 1,
		Calldata:       encodedCallData,
		AskCount:       params.AskCount,
		MinCount:       params.MinCount,
		FeeLimit:       params.FeeLimit,
		PrepareGas:     params.PrepareGas,
		ExecuteGas:     params.ExecuteGas,
	}

	seq, err := k.SendPacket(ctx, oracleRequest, types.PortID, params.ChannelId)
	if err != nil {
		return seq, err
	}

	k.SetClientIDCount(ctx, clientID+1)

	return seq, nil
}
