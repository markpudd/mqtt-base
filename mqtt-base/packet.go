package mqtt

// Packet - Interface for marshalling and unmarshalling packets
type Packet interface {
	PacketType() byte
	unmarshal(data []byte) error
	marshal() ([]byte, error)
}
