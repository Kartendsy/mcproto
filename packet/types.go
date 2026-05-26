package packet

// ConnectionState untuk tracking player state
type ConnectionState int

const (
	StateHandshake ConnectionState = 0
	StateStatus    ConnectionState = 1
	StateLogin     ConnectionState = 2
	StatePlay      ConnectionState = 3
)

// Handshake State Packets
const (
	HandshakePacketID = 0x00
)

// Status State Packets
const (
	StatusResponsePacketID = 0x00
	StatusPingPacketID     = 0x01
)

// Login State Packets
const (
	LoginStartPacketID           = 0x00
	LoginDisconnectPacketID      = 0x00
	LoginEncryptionRequestID     = 0x01
	LoginSuccessPacketID         = 0x02
	LoginSetCompressionPacketID  = 0x03
)

// Play State Packets
const (
	PlayKeepAlivePacketID              = 0x00
	PlayJoinGamePacketID               = 0x01
	PlayChatPacketID                   = 0x02
	PlayPlayerPositionAndLookPacketID  = 0x04
	PlayPlayerLookPacketID             = 0x05
	PlayPlayerPositionPacketID         = 0x06
	PlayMapChunkDataPacketID           = 0x21
	PlayPlayerListHeaderFooterPacketID = 0x48
	PlayDisconnectPacketID             = 0x40
)

// Helper function untuk convert state to string
func (s ConnectionState) String() string {
	switch s {
	case StateHandshake:
		return "Handshake"
	case StateStatus:
		return "Status"
	case StateLogin:
		return "Login"
	case StatePlay:
		return "Play"
	default:
		return "Unknown"
	}
}
