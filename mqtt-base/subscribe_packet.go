package mqtt

import "errors"

// TopicFilter - TopicFilter  structure
type TopicFilter struct {
	filter string
	qos    byte
}

// SubscribePacket - Subscribe packet structure
type SubscribePacket struct {
	fixedHeader *FixedHeader
	id          uint16
	topics      []*TopicFilter
}

// PacketType - Returns packet type
func (p *SubscribePacket) PacketType() byte {
	return Subscribe
}

// NewSubscribePacket - Creates a new Subscribe Packet
func NewSubscribePacket() *SubscribePacket {
	packet := new(SubscribePacket)
	packet.fixedHeader = new(FixedHeader)
	packet.fixedHeader.cntrlPacketType = Subscribe
	packet.fixedHeader.remaingLength = 2
	// TODO - Fix this.......
	packet.topics = make([]*TopicFilter, 0, 32)
	return packet
}

func (p *SubscribePacket) addTopic(topic *TopicFilter) {
	p.topics = append(p.topics, topic)
}

func (p *SubscribePacket) marshal() ([]byte, error) {
	totalLen := 0
	for _, topic := range p.topics {
		// lentgh of string plus 2 (len) + qos
		totalLen = totalLen + len(topic.filter) + 3
	}
	p.fixedHeader.remaingLength = uint32(totalLen + 2)
	fixedHeader := p.fixedHeader.marshal()
	// 2 is variable header
	data := make([]byte, 0, len(fixedHeader)+totalLen+2)
	data = append(data, fixedHeader...)
	// append ID
	data = append(data, byte(p.id>>8))
	data = append(data, byte(p.id))

	for _, topic := range p.topics {
		str, _ := EncodeString(topic.filter)
		data = append(data, str...)
		data = append(data, topic.qos)
	}
	return data, nil
}

func (p *SubscribePacket) unmarshal(data []byte) error {
	var err error
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.fixedHeader = fh
	if fh.remaingLength < 2 {
		return errors.New("No subscribe packet id")
	}
	p.id = uint16(data[2])<<8 | uint16(data[3])
	// TODO - Fix this.......
	p.topics = make([]*TopicFilter, 0, 32)
	pos := 4
	for uint32(pos) < fh.remaingLength+2 {
		topic := new(TopicFilter)
		topic.filter, err = UnencodeString(data[pos:])
		if err != nil {
			return err
		}
		pos = pos + len(topic.filter) + 2
		topic.qos = data[pos]
		p.addTopic(topic)
		pos++
	}
	return nil
}
