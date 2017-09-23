package mqttbase

import (
	"testing"
)

func TestPingRespPacketConstructor(t *testing.T) {
	packet := NewPingRespPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
		if packet.PacketType() != Pingresp {
			t.Errorf("Packet types should be Pingresp")
		}
	}
}

func TestMarshalPingResp(t *testing.T) {
	packet := NewPingRespPacket()
	data, _ := packet.Marshal()
	if len(data) != 2 {
		t.Errorf("Data length wrong should be 2 but was %d", len(data))
	} else {
		if data[0] != 0xd0 ||
			data[1] != 0 {
			t.Errorf("Bytes  are {%d,%d}  but should be {0xd0,0}",
				data[0], data[1])
		}

	}
}

func TestUnmarshalPingResp(t *testing.T) {
	data := []byte{0xd0, 0}
	packet := new(PingRespPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
}
