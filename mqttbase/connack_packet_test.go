package mqttbase

import (
	"testing"
)

func TestConnackPacketConstructor(t *testing.T) {
	packet := NewConnackPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalConnack(t *testing.T) {
	packet := NewConnackPacket()
	packet.SessionPresent = false
	packet.ReturnCode = ConAccept
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data lentgh wrong should be 4 but was %d", len(data))
	} else {
		if data[0] != 32 ||
			data[1] != 2 ||
			data[2] != 0 ||
			data[3] != 0 {
			t.Errorf("Bytes  are {%d,%d,%d,%d}  but should be {32,2,0,0}",
				data[0], data[1], data[2], data[3])
		}
	}
}

func TestMarshalConnackSession(t *testing.T) {
	packet := NewConnackPacket()
	packet.SessionPresent = true
	packet.ReturnCode = ConAccept
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data lentgh wrong should be 4 but was %d", len(data))
	} else {
		if data[0] != 32 ||
			data[1] != 2 ||
			data[2] != 1 ||
			data[3] != 0 {
			t.Errorf("Bytes  are {%d,%d,%d,%d}  but should be {32,2,1,0}",
				data[0], data[1], data[2], data[3])
		}
	}
}

func TestUnmarshalConnack(t *testing.T) {
	data := []byte{32, 2, 1, 0}
	packet := new(ConnackPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if !packet.SessionPresent {
		t.Errorf("Session not checked")
	}
}

func TestUnmarshalConnackReturnType(t *testing.T) {
	data := []byte{32, 2, 0, 3}
	packet := new(ConnackPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.SessionPresent {
		t.Errorf("Session  checked")
	}
	if packet.ReturnCode != ConRefuseServerUnavailable {
		t.Errorf("Return code incorrect")
	}
}
