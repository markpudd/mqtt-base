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
	fixedHeader *FixedHeader
	id          uint16
	returnCodes []byte
}

// PacketType - Returns packet type
func (p *SubackPacket) PacketType() byte {
	return Subscribe
}

// NewSubackPacket - Creates a new Suback Packet
func NewSubackPacket() *SubackPacket {
	packet := new(SubackPacket)
	packet.fixedHeader = new(FixedHeader)
	packet.fixedHeader.cntrlPacketType = Suback
	packet.fixedHeader.remaingLength = 2
	// TODO - Fix this.......
	packet.returnCodes = make([]byte, 0, 32)
	return packet
}

func (p *SubackPacket) addReturnCode(code byte) {
	p.returnCodes = append(p.returnCodes, code)
}

func (p *SubackPacket) Marshal() ([]byte, error) {
	p.fixedHeader.remaingLength = uint32(len(p.returnCodes) + 2)
	fixedHeader := p.fixedHeader.Marshal()
	// 2 is variable header
	data := make([]byte, 0, len(p.returnCodes)+2)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.id>>8))
	data = append(data, byte(p.id))

	for _, code := range p.returnCodes {
		data = append(data, code)
	}
	return data, nil
}

func (p *SubackPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.fixedHeader = fh
	if fh.remaingLength < 2 {
		return errors.New("No subscribe ack packet id")
	}
	p.id = uint16(data[2])<<8 | uint16(data[3])
	// TODO - Fix this.......
	p.returnCodes = make([]byte, 0, 32)
	pos := 4
	for uint32(pos) < fh.remaingLength+2 {
		p.addReturnCode(data[pos])
		pos++
	}
	return nil
}
