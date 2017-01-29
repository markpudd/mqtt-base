package mqtt

import (
	"errors"
)

// ProtocolVersion - ProtocolVersion contstant
const ProtocolVersion byte = 0x04

// CleanSession - clean session bit
const CleanSession byte = 0x02

// WillFlag - will flag bit
const WillFlag byte = 0x04

// WillQos1Flag - will qos 1 flag bit
const WillQos1Flag byte = 0x08

// WillQos2Flag - will qos 2 flag bit
const WillQos2Flag byte = 0x10

// WillRetainFlag - will retain flag bit
const WillRetainFlag byte = 0x20

// PasswordFlag - will password flag bit
const PasswordFlag byte = 0x40

// UsernameFlag - will username flag bit
const UsernameFlag byte = 0x80

// ConnectPacket - ConnectPacket data structure
type ConnectPacket struct {
	fixedHeader  *FixedHeader
	length       uint16
	userNameFlag bool
	passwordFlag bool
	willRetain   bool
	qos          int
	willFlag     bool
	cleanSession bool
	keepAlive    uint16
	id           uint16
	clientID     string
	willTopic    string
	willMessage  string
	username     string
	password     string
}

// PacketType - Returns packet type
func (p *ConnectPacket) PacketType() byte {
	return Connect
}

// NewConnectPacket - Creates a new Connect Packet
func NewConnectPacket() *ConnectPacket {
	packet := new(ConnectPacket)
	packet.fixedHeader = new(FixedHeader)
	packet.fixedHeader.cntrlPacketType = Connect
	return packet
}

func (p *ConnectPacket) marshalVariableHeader() []byte {
	data := make([]byte, 0, 10)
	str, _ := EncodeString("MQTT")
	data = append(data, str...)
	data = append(data, ProtocolVersion)
	var flags byte
	if p.userNameFlag {
		flags = flags | UsernameFlag
	}
	if p.passwordFlag {
		flags = flags | PasswordFlag
	}
	if p.willRetain {
		flags = flags | WillRetainFlag
	}
	if p.willFlag {
		flags = flags | WillFlag
	}
	if p.cleanSession {
		flags = flags | CleanSession
	}
	switch p.qos {
	case 1:
		flags = flags | WillQos1Flag
	case 2:
		flags = flags | WillQos2Flag
	case 3:
		flags = flags | WillQos1Flag | WillQos2Flag
	}
	data = append(data, flags)
	data = append(data, byte(p.keepAlive>>8&0xFF))
	data = append(data, byte(p.keepAlive&0xFF))
	return data
}

func (p *ConnectPacket) unmarshalVariableHeader(data []byte) error {
	str, err := UnencodeString(data[0:6])
	if err != nil || str != "MQTT" {
		return errors.New("Wrong protocol")
	}
	p.keepAlive = (uint16(data[8]) << 8) | uint16(data[9])
	if (data[7] & UsernameFlag) != 0 {
		p.userNameFlag = true
	}
	if (data[7] & PasswordFlag) != 0 {
		p.passwordFlag = true
	}
	if (data[7] & WillRetainFlag) != 0 {
		p.willRetain = true
	}
	if (data[7] & WillFlag) != 0 {
		p.willFlag = true
	}
	if (data[7] & CleanSession) != 0 {
		p.cleanSession = true
	}
	if (data[7] & WillQos1Flag) != 0 {
		p.qos = 0x1
	}
	if (data[7] & WillQos2Flag) != 0 {
		p.qos = p.qos | 0x2
	}
	return nil
}

func (p *ConnectPacket) marshal() ([]byte, error) {
	var sd []byte
	if len(p.clientID) > 0 {
		p.fixedHeader.remaingLength = uint32(10 + len(p.clientID) + 2)
		var err error
		sd, err = EncodeString(p.clientID)
		if err != nil {
			return nil, err
		}
	} else {
		p.fixedHeader.remaingLength = 12
		sd = []byte{0, 0}
	}
	fixedHeader := p.fixedHeader.marshal()
	variableHeader := p.marshalVariableHeader()
	data := make([]byte, 0, len(fixedHeader)+12)
	data = append(data, fixedHeader...)
	data = append(data, variableHeader...)
	data = append(data, sd...)
	return data, nil
}

func (p *ConnectPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.fixedHeader = fh
	err := p.unmarshalVariableHeader(data[2:])
	if err != nil {
		return err
	}
	p.clientID, err = UnencodeString(data[12:])
	if err != nil {
		return err
	}
	if p.clientID == "" {
		p.clientID = "NewId"
	}
	return nil
}
