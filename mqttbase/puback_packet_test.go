package mqttbase

import (
	"testing"
)

func TestPubackPacketConstructor(t *testing.T) {
	packet := NewPubackPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.fixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalPuback(t *testing.T) {
	packet := NewPubackPacket()
	packet.id = 0x3acd
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data length wrong should be 4 but was %d", len(data))
	} else {
		expected := []byte{0x40, 2, 0x3a, 0xcd}
		for i := 0; i < 4; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalPubackReturnType(t *testing.T) {
	data := []byte{0x40, 2, 0x3a, 0xcd}
	packet := new(PubackPacket)
	packet.unmarshal(data)
	if packet.fixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.id != 0x3acd {
		t.Errorf("Packet id is %#x but should 0x3acd", packet.id)
	}
}
