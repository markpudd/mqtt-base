package mqttbase

import "errors"

type PubrelPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
}

// PacketType - Returns packet type
func (p *PubrelPacket) PacketType() byte {
	return Pubrel
}

// NewSubackPacket - Creates a new Suback Packet
func NewPubrelPacket() *PubrelPacket {
	packet := new(PubrelPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Pubrel
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *PubrelPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, 4)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.Id>>8))
	data = append(data, byte(p.Id))
	return data, nil
}

func (p *PubrelPacket) unmarshal(data []byte) error {
	if data[0] != 0x62 {
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
