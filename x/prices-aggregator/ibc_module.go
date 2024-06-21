package prices_aggregator

import (
	// "fmt"

	"fmt"
	"strconv"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/vuong177/macro/x/prices-aggregator/keeper"
	"github.com/vuong177/macro/x/prices-aggregator/types"
	// ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
)

var _ porttypes.IBCModule = IBCModule{}

type IBCModule struct {
	keeper keeper.Keeper
}

func ValidateChannelParams(
	ctx sdk.Context,
	keeper keeper.Keeper,
	order channeltypes.Order,
	portID string,
	channelID string,
) error {
	if order != channeltypes.UNORDERED {
		return errorsmod.Wrapf(channeltypes.ErrInvalidChannelOrdering, "expected %s channel, got %s ", channeltypes.UNORDERED, order)
	}

	// Require portID is the portID profiles module is bound to
	boundPort := keeper.GetPort(ctx)
	if boundPort != portID {
		return errorsmod.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	return nil
}

// OnChanOpenInit implements the IBCModule interface.
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	if err := ValidateChannelParams(ctx, im.keeper, order, portID, channelID); err != nil {
		return "", err
	}

	if strings.TrimSpace(version) == "" {
		version = types.Version
	}

	if version != types.Version {
		return "", errorsmod.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
	}

	// Claim channel capability passed back by IBC module
	if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	return version, nil
}

// OnChanOpenTry implements the IBCModule interface.
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	if err := ValidateChannelParams(ctx, im.keeper, order, portID, channelID); err != nil {
		return "", err
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !im.keeper.AuthenticateCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		err := im.keeper.ClaimCapability(ctx, channelCap, host.ChannelCapabilityPath(portID, channelID))
		if err != nil {
			return "", err
		}
	}

	return counterpartyVersion, nil
}

// OnChanOpenAck implements the IBCModule interface.
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface.
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface.
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return errorsmod.Wrap(errorstypes.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface.
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// no need to implement
	return nil
}

// OnRecvPacket implements the IBCModule interface.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	ack := channeltypes.NewResultAcknowledgement([]byte{byte(1)})

	var oracleResponse types.OracleResponsePacketData
	var ackErr error
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &oracleResponse); err != nil {
		ackErr = errorsmod.Wrapf(errorstypes.ErrInvalidType, "cannot unmarshal OracleResponsePacketData")
		ack = channeltypes.NewErrorAcknowledgement(ackErr)
	}

	// only attempt the application logic if the packet data
	// was successfully decoded
	if ack.Success() {
		err := im.keeper.OnRecvPacket(ctx, oracleResponse)
		if err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
			ackErr = err
		}
	}

	eventAttributes := []sdk.Attribute{
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(types.AttributeClientID, oracleResponse.ClientID),
		sdk.NewAttribute(types.AttributeResolveStatus, oracleResponse.ResolveStatus.String()),
	}

	if ackErr != nil {
		eventAttributes = append(eventAttributes, sdk.NewAttribute(types.AttributeKeyAckError, ackErr.Error()))
	} else {
		eventAttributes = append(eventAttributes, sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", ack.Success())))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOracleResponsePacket,
			eventAttributes...,
		),
	)

	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return errorsmod.Wrapf(errorstypes.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}

	if err := im.keeper.OnAcknowledgementPacket(ctx, ack, packet); err != nil {
		return err
	}

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeOracleRequestPacketAcknowledgement,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
				sdk.NewAttribute(types.AttributeKeyPacketSequence, strconv.FormatUint(packet.Sequence, 10)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeOracleRequestPacketAcknowledgement,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
				sdk.NewAttribute(types.AttributeKeyPacketSequence, strconv.FormatUint(packet.Sequence, 10)),
			),
		)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {

	return nil
}
