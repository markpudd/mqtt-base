package mqttbase

// Packet - Interface for Marshalling and unmarshalling packets
type Packet interface {
	PacketType() byte
	unmarshal(data []byte) error
	Marshal() ([]byte, error)
}
