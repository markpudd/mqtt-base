package mqttbase

import "errors"

// SuccessQos0 - SuccessQos0 return code
const SuccessQos0 byte = 0x00

// SuccessQos1 - SuccessQos1 return code
const SuccessQos1 byte = 0x01

// SuccessQos2 - SuccessQos2 return code
const SuccessQos2 byte = 0x02

// Failure - Failure return code
const Failure byte = 0xF0

// SubackPacket - Subcription ack packet structure
type SubackPacket struct {
	FixedHeader *FixedHeader
	Id          uint16
	ReturnCodes []byte
}

// PacketType - Returns packet type
func (p *SubackPacket) PacketType() byte {
	return Suback
}

// NewSubackPacket - Creates a new Suback Packet
func NewSubackPacket() *SubackPacket {
	packet := new(SubackPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Suback
	packet.FixedHeader.remaingLength = 2
	// TODO - Fix this.......
	packet.ReturnCodes = make([]byte, 0, 32)
	return packet
}

func (p *SubackPacket) AddReturnCode(code byte) {
	p.ReturnCodes = append(p.ReturnCodes, code)
}

func (p *SubackPacket) Marshal() ([]byte, error) {
	p.FixedHeader.remaingLength = uint32(len(p.ReturnCodes) + 2)
	fixedHeader := p.FixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, len(p.ReturnCodes)+2)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.Id>>8))
	data = append(data, byte(p.Id))

	for _, code := range p.ReturnCodes {
		data = append(data, code)
	}
	return data, nil
}

func (p *SubackPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	if fh.remaingLength < 2 {
		return errors.New("No subscribe ack packet id")
	}
	p.Id = uint16(data[2])<<8 | uint16(data[3])
	// TODO - Fix this.......
	p.ReturnCodes = make([]byte, 0, 32)
	pos := 4
	for uint32(pos) < fh.remaingLength+2 {
		p.AddReturnCode(data[pos])
		pos++
	}
	return nil
}
