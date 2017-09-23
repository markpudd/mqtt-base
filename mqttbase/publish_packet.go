package mqttbase

import "errors"

// PublishPacket - Publish packet structure
type PublishPacket struct {
	FixedHeader *FixedHeader
	TopicName   string
	Id          uint16
	Data        []byte
}

// PacketType - Returns packet type
func (p *PublishPacket) PacketType() byte {
	return Publish
}

// NewPublishPacket - Creates a new Publish Packet
func NewPublishPacket() *PublishPacket {
	packet := new(PublishPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Publish
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *PublishPacket) Marshal() ([]byte, error) {
	totalLength := len(p.TopicName) + 2 + 2 + len(p.Data)
	p.FixedHeader.remaingLength = uint32(totalLength)
	fixedHeader := p.FixedHeader.Marshal()
	data := make([]byte, 0, len(fixedHeader)+totalLength+2)
	data = append(data, fixedHeader...)
	// append ID
	str, _ := EncodeString(p.TopicName)
	data = append(data, str...)
	data = append(data, byte(p.Id>>8))
	data = append(data, byte(p.Id))
	data = append(data, p.Data...)

	return data, nil
}

func (p *PublishPacket) unmarshal(data []byte) error {
	var err error
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	if fh.remaingLength < 2 {
		return errors.New("No remaining length")
	}
	pos := 2
	p.TopicName, err = UnencodeString(data[pos:])
	if err != nil {
		return err
	}
	pos = pos + len(p.TopicName) + 2
	p.Id = uint16(data[pos])<<8 | uint16(data[pos+1])
	p.Data = data[pos+2:]
	return nil
}
