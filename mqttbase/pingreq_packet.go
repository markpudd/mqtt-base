package mqttbase

type PingReqPacket struct {
	FixedHeader *FixedHeader
}

func NewPingReqPacket() *PingReqPacket {
	packet := new(PingReqPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Pingreq
	packet.FixedHeader.remaingLength = 0
	return packet
}

func (p *PingReqPacket) PacketType() byte {
	return Pingreq
}

func (p *PingReqPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	data := make([]byte, 0, 2)
	data = append(data, fixedHeader...)
	data = append(data, 0)
	return data, nil

}

func (p *PingReqPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	return nil
}
