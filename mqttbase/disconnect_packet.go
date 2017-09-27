package mqttbase

import "errors"

type DisconnectPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
}

// PacketType - Returns packet type
func (p *DisconnectPacket) PacketType() byte {
	return Disconnect
}

// NewSubackPacket - Creates a new Suback Packet
func NewDisconnectPacket() *DisconnectPacket {
	packet := new(DisconnectPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Disconnect
	packet.FixedHeader.remaingLength = 0
	return packet
}

func (p *DisconnectPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, 2)
	data = append(data, fixedHeader...)
	return data, nil
}

func (p *DisconnectPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	if fh.remaingLength != 0 {
		return errors.New("Unsuback length wrong")
	}
	return nil
}
