package mqttbase

// PubackPacket - Publish ack packet structure
type PubackPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
}

// PacketType - Returns packet type
func (p *PubackPacket) PacketType() byte {
	return Puback
}

// NewPubackPacket - Creates a new Publish ack Packet
func NewPubackPacket() *PubackPacket {
	packet := new(PubackPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Puback
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *PubackPacket) Marshal() ([]byte, error) {
	return []byte{0x40, 2, byte(p.Id >> 8), byte(p.Id)}, nil
}

func (p *PubackPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	p.Id = uint16(data[2])<<8 | uint16(data[3])
	return nil
}
