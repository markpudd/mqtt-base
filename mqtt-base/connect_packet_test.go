package mqtt

import (
	"strings"
	"testing"
)

func TestConstructorConnectPacket(t *testing.T) {
	packet := NewConnectPacket()
	if packet.PacketType() != Connect {
		t.Errorf("Packet types should be Connect")
	}
	if packet.userNameFlag {
		t.Errorf("Username flag should be false")
	}
	if packet.passwordFlag {
		t.Errorf("Password flag should be false")
	}
	if packet.willRetain {
		t.Errorf("will Retain flag should be false")
	}
	if packet.willFlag {
		t.Errorf("will flag should be false")
	}
	if packet.cleanSession {
		t.Errorf("cleanSession flag should be false")
	}
	if packet.qos > 0 {
		t.Errorf("qos flag should be 0")
	}
}

func TestMarshalConnectHeader(t *testing.T) {
	packet := NewConnectPacket()
	packet.keepAlive = 258

	data := packet.marshalVariableHeader()

	if len(data) != 10 {
		t.Errorf("Data length %d when should be 10", len(data))
	} else {
		if data[0] != 0 {
			t.Errorf("First byte %d when should be 0", data[0])
		}
		if data[1] != 4 {
			t.Errorf("Second byte  %d when should be 4", data[1])
		}
		if data[2] != 'M' ||
			data[3] != 'Q' ||
			data[4] != 'T' ||
			data[5] != 'T' {
			t.Errorf("Bytes 3-6 should are {%d,%d,%d,%d}  but should be {M,Q,T,T}",
				data[2], data[3], data[4], data[5])
		}
		if data[6] != 0x04 {
			t.Errorf("Protocol version byte  %d when should be 0x04", data[6])
		}
		if data[7]&0x01 != 0 {
			t.Errorf("Protocol flag reserved bit 1 but should be 0")
		}
		if data[8] != 1 && data[9] != 2 {
			t.Errorf("Keep Alive should be 1,2 but is %d,%d", data[8], data[9])
		}
	}
}

func TestMarshalUsernameFlag(t *testing.T) {
	packet := NewConnectPacket()
	packet.userNameFlag = true
	data := packet.marshalVariableHeader()
	if data[7] != UsernameFlag {
		t.Errorf("UsernameFlag not set or other flags set %d", data[7])
	}
}

func TestMarshalPasswordFlag(t *testing.T) {
	packet := NewConnectPacket()
	packet.passwordFlag = true
	data := packet.marshalVariableHeader()
	if data[7] != PasswordFlag {
		t.Errorf("PasswordFlag not set or other flags set %d", data[7])
	}
}

func TestMarshalWillRetainFlag(t *testing.T) {
	packet := NewConnectPacket()
	packet.willRetain = true
	data := packet.marshalVariableHeader()
	if data[7] != WillRetainFlag {
		t.Errorf("WillRetainFlag not set or other flags set %d", data[7])
	}
}

func TestMarshalWillFlag(t *testing.T) {
	packet := NewConnectPacket()
	packet.willFlag = true
	data := packet.marshalVariableHeader()
	if data[7] != WillFlag {
		t.Errorf("WillFlag not set or other flags set %d", data[7])
	}
}

func TestMarshalCleanSession(t *testing.T) {
	packet := NewConnectPacket()
	packet.cleanSession = true
	data := packet.marshalVariableHeader()
	if data[7] != CleanSession {
		t.Errorf("CleanSession not set or other flags set %d", data[7])
	}
}

func TestQOSVale(t *testing.T) {
	packet := NewConnectPacket()
	packet.qos = 1
	data := packet.marshalVariableHeader()
	if data[7] != WillQos1Flag {
		t.Errorf("WillQos1Flag not set or other flags set %d", data[7])
	}
	packet.qos = 2
	data = packet.marshalVariableHeader()
	if data[7] != WillQos2Flag {
		t.Errorf("WillQos2Flag not set or other flags set %d", data[7])
	}
	packet.qos = 3
	data = packet.marshalVariableHeader()
	if data[7] != (WillQos1Flag | WillQos2Flag) {
		t.Errorf("WillQos1Flag and WillQos2Flag not set or other flags set %d", data[7])
	}

}

func TestUnmarshalConnectHeader(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, 0, 1, 2}
	packet := NewConnectPacket()
	err := packet.unmarshalVariableHeader(data)
	if err != nil {
		t.Errorf("Bad header %s", err)
	}

	if packet.PacketType() != Connect {
		t.Errorf("Packet type %d when should be %d", packet.PacketType(), Connect)
	}

	if packet.keepAlive != 258 {
		t.Errorf("keepAlive %d when should be 258", packet.keepAlive)
	}
}

func TestUnmarshalBadHeader(t *testing.T) {
	data := []byte{0, 4, 'P', 'Q', 'T', 'T', 4, 0, 1, 2}
	packet := NewConnectPacket()
	err := packet.unmarshalVariableHeader(data)
	if err == nil {
		t.Errorf("Should be error in protocol")
	}
}

func TestUnmarshalUsernameFlag(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, UsernameFlag, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.userNameFlag != true {
		t.Errorf("userNameFlag not set on unmarshal %d", data[7])
	}
}

func TestUnmarshalPasswordFlag(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, PasswordFlag, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.passwordFlag != true {
		t.Errorf("passwordFlag not set on unmarshal")
	}
}

func TestUnmarshalWillRetainFlag(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, WillRetainFlag, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.willRetain != true {
		t.Errorf("willRetain not set on unmarshal")
	}
}

func TestUnmarshalWillFlag(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, WillFlag, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.willFlag != true {
		t.Errorf("willFlag not set on unmarshal")
	}
}

func TestUnmarshalCleanSession(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, CleanSession, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.cleanSession != true {
		t.Errorf("cleanSession not set on unmarshal")
	}
}

func TestUnmarshalQOSVale(t *testing.T) {
	data := []byte{0, 4, 'M', 'Q', 'T', 'T', 4, WillQos1Flag, 1, 2}
	packet := NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.qos != 1 {
		t.Errorf("qos is %d but should be 1 on unmarshal", packet.qos)
	}

	data = []byte{0, 4, 'M', 'Q', 'T', 'T', 4, WillQos2Flag, 1, 2}
	packet = NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.qos != 2 {
		t.Errorf("qos is %d but should be 2 on unmarshal", packet.qos)
	}

	data = []byte{0, 4, 'M', 'Q', 'T', 'T', 4, WillQos1Flag | WillQos2Flag, 1, 2}
	packet = NewConnectPacket()
	packet.unmarshalVariableHeader(data)
	if packet.qos != 3 {
		t.Errorf("qos is %d but should be 3 on unmarshal", packet.qos)
	}

}

func TestPacketConstructor(t *testing.T) {
	packet := NewConnectPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.fixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalConnectServerId(t *testing.T) {
	packet := NewConnectPacket()
	data, _ := packet.marshal()

	// Should be 2+10+2
	if len(data) != 14 {
		t.Errorf("Data length %d when should be 14", len(data))
	} else {
		if data[1] != 12 {
			t.Errorf("Length is %d when should be 12", data[1])
		}
		if data[12] != 0 || data[13] != 0 {
			t.Errorf("Client Id is %d,%d when should be 0,0", data[12], data[13])
		}
	}
}

func TestMarshalConnectServerIdSet(t *testing.T) {
	packet := NewConnectPacket()
	packet.clientID = "afjdj3c"
	data, _ := packet.marshal()

	// Should be 2+10+9
	if len(data) != 21 {
		t.Errorf("Data length %d when should be 21", len(data))
	} else {
		if data[1] != 19 {
			t.Errorf("Length is %d when should be 12", data[1])
		}
	}
}

func TestUnmarshalConnect(t *testing.T) {
	data := []byte{21, 14, 0, 4, 'M', 'Q', 'T', 'T', 4, 0, 1, 2, 0, 0}
	packet := new(ConnectPacket)
	packet.unmarshal(data)
	if packet.fixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.clientID == "" {
		t.Errorf("clientID header empty")
	}
}

func TestUnmarshalConnectWithId(t *testing.T) {
	data := []byte{21, 14, 0, 4, 'M', 'Q', 'T', 'T', 4, 0, 1, 2, 0, 4, 't', 'e', 's', 't'}
	packet := new(ConnectPacket)
	packet.unmarshal(data)
	if packet.fixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if strings.Compare(packet.clientID, "test") != 0 {
		t.Errorf("clientID should be test but is %s", packet.clientID)
	}
}
