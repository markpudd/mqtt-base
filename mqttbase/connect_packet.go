package mqttbase

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
	FixedHeader  *FixedHeader
	Length       uint16
	UserNameFlag bool
	PasswordFlag bool
	WillRetain   bool
	Qos          int
	WillFlag     bool
	CleanSession bool
	KeepAlive    uint16
	Id           uint16
	ClientID     string
	WillTopic    string
	WillMessage  string
	Username     string
	Password     string
}

// PacketType - Returns packet type
func (p *ConnectPacket) PacketType() byte {
	return Connect
}

// NewConnectPacket - Creates a new Connect Packet
func NewConnectPacket() *ConnectPacket {
	packet := new(ConnectPacket)
	packet.FixedHeader = new(FixedHeader)
	packet.FixedHeader.cntrlPacketType = Connect
	return packet
}

func (p *ConnectPacket) MarshalVariableHeader() []byte {
	data := make([]byte, 0, 10)
	str, _ := EncodeString("MQTT")
	data = append(data, str...)
	data = append(data, ProtocolVersion)
	var flags byte
	if p.UserNameFlag {
		flags = flags | UsernameFlag
	}
	if p.PasswordFlag {
		flags = flags | PasswordFlag
	}
	if p.WillRetain {
		flags = flags | WillRetainFlag
	}
	if p.WillFlag {
		flags = flags | WillFlag
	}
	if p.CleanSession {
		flags = flags | CleanSession
	}
	switch p.Qos {
	case 1:
		flags = flags | WillQos1Flag
	case 2:
		flags = flags | WillQos2Flag
	case 3:
		flags = flags | WillQos1Flag | WillQos2Flag
	}
	data = append(data, flags)
	data = append(data, byte(p.KeepAlive>>8&0xFF))
	data = append(data, byte(p.KeepAlive&0xFF))
	return data
}

func (p *ConnectPacket) unmarshalVariableHeader(data []byte) error {
	str, err := UnencodeString(data[0:6])
	if err != nil || str != "MQTT" {
		return errors.New("Wrong protocol")
	}
	p.KeepAlive = (uint16(data[8]) << 8) | uint16(data[9])
	if (data[7] & UsernameFlag) != 0 {
		p.UserNameFlag = true
	}
	if (data[7] & PasswordFlag) != 0 {
		p.PasswordFlag = true
	}
	if (data[7] & WillRetainFlag) != 0 {
		p.WillRetain = true
	}
	if (data[7] & WillFlag) != 0 {
		p.WillFlag = true
	}
	if (data[7] & CleanSession) != 0 {
		p.CleanSession = true
	}
	if (data[7] & WillQos1Flag) != 0 {
		p.Qos = 0x1
	}
	if (data[7] & WillQos2Flag) != 0 {
		p.Qos = p.Qos | 0x2
	}
	return nil
}

func (p *ConnectPacket) Marshal() ([]byte, error) {
	var sd []byte
	if len(p.ClientID) > 0 {
		p.FixedHeader.remaingLength = uint32(10 + len(p.ClientID) + 2)
		var err error
		sd, err = EncodeString(p.ClientID)
		if err != nil {
			return nil, err
		}
	} else {
		p.FixedHeader.remaingLength = 12
		sd = []byte{0, 0}
	}
	fixedHeader := p.FixedHeader.Marshal()
	variableHeader := p.MarshalVariableHeader()
	data := make([]byte, 0, len(fixedHeader)+12)
	data = append(data, fixedHeader...)
	data = append(data, variableHeader...)
	data = append(data, sd...)
	return data, nil
}

func (p *ConnectPacket) unmarshal(data []byte) error {
	fh := new(FixedHeader)
	fh.unmarshal(data)
	p.FixedHeader = fh
	err := p.unmarshalVariableHeader(data[2:])
	if err != nil {
		return err
	}
	p.ClientID, err = UnencodeString(data[12:])
	if err != nil {
		return err
	}
	if p.ClientID == "" {
		p.ClientID = "NewId"
	}
	return nil
}
