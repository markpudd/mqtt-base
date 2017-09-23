package mqttbase

type PingRespPacket struct {
	FixedHeader *FixedHeader
}

func NewPingRespPacket() *PingRespPacket {
	packet := new(PingRespPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Pingresp
	packet.FixedHeader.remaingLength = 0
	return packet
}

func (p *PingRespPacket) PacketType() byte {
	return Pingresp
}

func (p *PingRespPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	data := make([]byte, 0, 2)
	data = append(data, fixedHeader...)
	data = append(data, 0)
	return data, nil

}

func (p *PingRespPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	return nil
}
