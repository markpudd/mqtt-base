package mqttbase

import (
	"errors"
)

// ConAccept - Connection Accept
const ConAccept byte = 0x00

// ConRefuseBadProtocol - Bad protocol
const ConRefuseBadProtocol byte = 0x01

// ConRefuseIDReject - rejected id
const ConRefuseIDReject byte = 0x02

// ConRefuseServerUnavailable - server unavailble
const ConRefuseServerUnavailable byte = 0x03

// ConRefuseBadCred - bad credential
const ConRefuseBadCred byte = 0x04

// ConRefuseNoAuth - no auth
const ConRefuseNoAuth byte = 0x05

// ConnackPacket - ConnackPacket stucture
type ConnackPacket struct {
	FixedHeader    *FixedHeader
	SessionPresent bool
	ReturnCode     byte
}

// PacketType - Returns packet type
func (p *ConnackPacket) PacketType() byte {
	return Connack
}

// NewConnackPacket - Creates a new Connack Packet
func NewConnackPacket() *ConnackPacket {
	packet := new(ConnackPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Connack
	packet.FixedHeader.remaingLength = 2
	return packet
}

func (p *ConnackPacket) Marshal() ([]byte, error) {
	fixedHeader := p.FixedHeader.Marshal()
	data := make([]byte, 0, 4)
	data = append(data, fixedHeader...)
	if p.SessionPresent {
		data = append(data, 1)
	} else {
		data = append(data, 0)
	}
	data = append(data, p.ReturnCode)
	return data, nil
}

func (p *ConnackPacket) unmarshal(data []byte) error {
	if len(data) != 4 {
		return errors.New("Bad Connack size")
	}
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	if data[2] == 1 {
		p.SessionPresent = true
	} else {
		p.SessionPresent = false
	}
	p.ReturnCode = data[3]
	return nil
}
