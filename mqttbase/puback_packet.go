package mqttbase

// PubackPacket - Publish ack packet structure
type PubackPacket struct {
	fixedHeader *FixedHeader
	id          uint16
}

// PacketType - Returns packet type
func (p *PubackPacket) PacketType() byte {
	return Puback
}

// NewPubackPacket - Creates a new Publish ack Packet
func NewPubackPacket() *PubackPacket {
	packet := new(PubackPacket)
	packet.fixedHeader = new(FixedHeader)
	packet.fixedHeader.cntrlPacketType = Puback
	packet.fixedHeader.remaingLength = 2
	return packet
}

func (p *PubackPacket) Marshal() ([]byte, error) {
	return []byte{0x40, 2, byte(p.id >> 8), byte(p.id)}, nil
}

func (p *PubackPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.fixedHeader = fh
	p.id = uint16(data[2])<<8 | uint16(data[3])
	return nil
}
