package mqtt

import "errors"

// PublishPacket - Publish packet structure
type PublishPacket struct {
	fixedHeader *FixedHeader
	topicName   string
	id          uint16
	data        []byte
}

// PacketType - Returns packet type
func (p *PublishPacket) PacketType() byte {
	return Publish
}

// NewPublishPacket - Creates a new Publish Packet
func NewPublishPacket() *PublishPacket {
	packet := new(PublishPacket)
	packet.fixedHeader = new(FixedHeader)
	packet.fixedHeader.cntrlPacketType = Publish
	packet.fixedHeader.remaingLength = 2
	return packet
}

func (p *PublishPacket) marshal() ([]byte, error) {
	totalLength := len(p.topicName) + 2 + 2 + len(p.data)
	p.fixedHeader.remaingLength = uint32(totalLength)
	fixedHeader := p.fixedHeader.marshal()
	data := make([]byte, 0, len(fixedHeader)+totalLength+2)
	data = append(data, fixedHeader...)
	// append ID
	str, _ := EncodeString(p.topicName)
	data = append(data, str...)
	data = append(data, byte(p.id>>8))
	data = append(data, byte(p.id))
	data = append(data, p.data...)

	return data, nil
}

func (p *PublishPacket) unmarshal(data []byte) error {
	var err error
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.fixedHeader = fh
	if fh.remaingLength < 2 {
		return errors.New("No remaining length")
	}
	pos := 2
	p.topicName, err = UnencodeString(data[pos:])
	if err != nil {
		return err
	}
	pos = pos + len(p.topicName) + 2
	p.id = uint16(data[pos])<<8 | uint16(data[pos+1])
	p.data = data[pos+2:]
	return nil
}
