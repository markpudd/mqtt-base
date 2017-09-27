package mqttbase

import "errors"

type PubcompPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
}

// PacketType - Returns packet type
func (p *PubcompPacket) PacketType() byte {
	return Pubcomp
}

// NewSubackPacket - Creates a new Suback Packet
func NewPubcompPacket() *PubcompPacket {
	packet := new(PubcompPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Pubcomp
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *PubcompPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, 4)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.Id>>8))
	data = append(data, byte(p.Id))
	return data, nil
}

func (p *PubcompPacket) unmarshal(data []byte) error {
	if data[0] != 0x70 {
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
