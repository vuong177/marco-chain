package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/vuong177/macro/x/prices-aggregator/types"
)

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
