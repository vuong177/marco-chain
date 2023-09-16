package types

const (
	// ModuleName store the name of module
	ModuleName = "prices-aggregator"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// RouterKey is the message route for slashing.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName

	// Version defines the current version the IBC oracle module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to.
	PortID = "oracle"
)

var (
	PortKey = KeyPrefix("band-oracle-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
