package types

// IBC transfer events
const (
	EventTypeOracleRequestPacket                = "oracle-request-packet"
	EventTypeOracleRequestPacketAcknowledgement = "oracle-request-packet-ack"
	EventTypeOracleResponsePacket               = "oracle-response-packet"

	AttributeKeyPacketSequence = "sequence"
	AttributeKeyAckSuccess     = "success"
	AttributeKeyAckError       = "error"
)
