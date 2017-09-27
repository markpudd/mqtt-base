package mqttbase

import (
	"testing"
)

func TestUnSubackPacketConstructor(t *testing.T) {
	packet := NewUnSubackPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalUnSuback(t *testing.T) {
	packet := NewUnSubackPacket()
	packet.Id = 0xa6f2
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data length wrong should be 4 but was %d", len(data))
	} else {
		expected := []byte{0xb0, 0x02, 0xa6, 0xf2}
		for i := 0; i < 4; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalUnSuback(t *testing.T) {
	data := []byte{0xb0, 0x02, 0xa6, 0xf2}
	packet := new(UnSubackPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.Id != 0xa6f2 {
		t.Errorf("Expecting id 0xa6f2   got %#x", packet.Id)
	}
}
