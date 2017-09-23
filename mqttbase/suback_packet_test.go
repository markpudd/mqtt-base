package mqttbase

import (
	"testing"
)

func TestSubackPacketConstructor(t *testing.T) {
	packet := NewSubackPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestReturnCode(t *testing.T) {
	packet := NewSubackPacket()
	packet.AddReturnCode(0x01)
	if len(packet.ReturnCodes) != 1 || packet.ReturnCodes[0] != 0x01 {
		t.Errorf("Return code added but not in list")
	}
}

func TestMarshalSuback(t *testing.T) {
	packet := NewSubackPacket()
	packet.AddReturnCode(SuccessQos0)
	packet.AddReturnCode(SuccessQos1)
	packet.AddReturnCode(SuccessQos2)
	packet.AddReturnCode(Failure)
	packet.Id = 0xa6f2
	data, _ := packet.Marshal()
	if len(data) != 8 {
		t.Errorf("Data lentgh wrong should be 5 but was %d", len(data))
	} else {
		expected := []byte{0x90, 0x06, 0xa6, 0xf2, 0x00, 0x01, 0x02, 0xf0}
		for i := 0; i < 8; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalSuback(t *testing.T) {
	data := []byte{0x90, 0x06, 0xa6, 0xf2, 0x00, 0x01, 0x02, 0xf0}
	packet := new(SubackPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.Id != 0xa6f2 {
		t.Errorf("Expecting id 0xa6f2   got %#x", packet.Id)
	}
	if len(packet.ReturnCodes) != 4 {
		t.Errorf("Expecting 4 return codes got %d", len(packet.ReturnCodes))
	} else {
		if packet.ReturnCodes[0] != SuccessQos0 ||
			packet.ReturnCodes[1] != SuccessQos1 ||
			packet.ReturnCodes[2] != SuccessQos2 ||
			packet.ReturnCodes[3] != Failure {
			t.Errorf("Bytes  are {%d,%d,%d,%d}  but should be {0x00,0x01,0x02,0xf0}",
				data[0], data[1], data[2], data[3])
		}
	}

}
