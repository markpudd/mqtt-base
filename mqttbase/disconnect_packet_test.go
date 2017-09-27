package mqttbase

import (
	"testing"
)

func TestDisconnectConstructor(t *testing.T) {
	packet := NewDisconnectPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalDisconnect(t *testing.T) {
	packet := NewDisconnectPacket()
	data, _ := packet.Marshal()
	if len(data) != 2 {
		t.Errorf("Data length wrong should be 2 but was %d", len(data))
	} else {
		expected := []byte{0xe0, 0x00}
		for i := 0; i < 2; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalDisconnect(t *testing.T) {
	data := []byte{0xc0, 0x00}
	packet := new(DisconnectPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
}
