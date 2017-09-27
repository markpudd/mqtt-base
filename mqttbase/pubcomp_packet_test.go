package mqttbase

import (
	"testing"
)

func TestPubcompPacketConstructor(t *testing.T) {
	packet := NewPubcompPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
	if packet.PacketType() != Pubcomp {
		t.Errorf("Packet types should be Pubrel")
	}
}

func TestMarshalPubcomp(t *testing.T) {
	packet := NewPubcompPacket()
	packet.Id = 0xa6f2
	data, _ := packet.Marshal()
	if len(data) != 4 {
		t.Errorf("Data length wrong should be 4 but was %d", len(data))
	} else {
		expected := []byte{0x70, 0x02, 0xa6, 0xf2}
		for i := 0; i < 4; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}
	}
}

func TestUnmarshalPubcomp(t *testing.T) {
	data := []byte{0x70, 0x02, 0xa6, 0xf2}
	packet := new(PubcompPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
	if packet.Id != 0xa6f2 {
		t.Errorf("Expecting id 0xa6f2   got %#x", packet.Id)
	}
}
