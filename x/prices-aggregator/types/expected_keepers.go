package types

import (
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	// connectiontypes "github.com/cosmos/ibc-go/v7/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	// ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
)

// PortKeeper defines the expected IBC port keeper.
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

// ScopedKeeper defines the expected scoped keeper.
type ScopedKeeper interface {
	AuthenticateCapability(ctx sdk.Context, capability *capabilitytypes.Capability, name string) bool
	ClaimCapability(ctx sdk.Context, capability *capabilitytypes.Capability, name string) error
	GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool)
}

// ChannelKeeper defines the expected IBC channel keeper.
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, sourcePort string, sourceChannel string,
		timeoutHeight clienttypes.Height, timeoutTimestamp uint64, data []byte) (uint64, error)
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
}
