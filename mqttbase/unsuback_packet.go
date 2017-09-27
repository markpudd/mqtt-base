package mqttbase

import "errors"

type UnSubackPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
}

// PacketType - Returns packet type
func (p *UnSubackPacket) PacketType() byte {
	return UnSuback
}

// NewSubackPacket - Creates a new Suback Packet
func NewUnSubackPacket() *UnSubackPacket {
	packet := new(UnSubackPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = UnSuback
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *UnSubackPacket) Marshal() ([]byte, error) {
	p.FixedHeader.remaingLength = 2
	fixedHeader := p.FixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, 4)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.Id>>8))
	data = append(data, byte(p.Id))
	return data, nil
}

func (p *UnSubackPacket) unmarshal(data []byte) error {
	if data[0] != 0xb0 {
		return errors.New("Wrong packet type")
	}
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	if fh.remaingLength != 2 {
		return errors.New("Unsuback length wrong")
	}
	p.Id = uint16(data[2])<<8 | uint16(data[3])
	return nil
}
