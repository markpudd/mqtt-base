package mqttbase

import (
	"testing"
)

func TestPubrecPacketConstructor(t *testing.T) {
	packet := NewPubrecPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
	if packet.PacketType() != Pubrec {
		t.Errorf("Packet types should be Pubrec")
	}
}

func TestMarshalPubrec(t *testing.T) {
	packet := NewPubrecPacket()
	packet.Id = 0xa6f2
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data length wrong should be 4 but was %d", len(data))
	} else {
		expected := []byte{0x50, 0x02, 0xa6, 0xf2}
		for i := 0; i < 4; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalPubrec(t *testing.T) {
	data := []byte{0x50, 0x02, 0xa6, 0xf2}
	packet := new(PubrecPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.Id != 0xa6f2 {
		t.Errorf("Expecting id 0xa6f2   got %#x", packet.Id)
	}
}
