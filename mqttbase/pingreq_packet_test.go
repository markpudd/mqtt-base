package mqttbase

import (
	"testing"
)

func TestPingReqPacketConstructor(t *testing.T) {
	packet := NewPingReqPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
	if packet.PacketType() != Pingreq {
		t.Errorf("Packet types should be Pingreq")
	}
}

func TestMarshalPingReq(t *testing.T) {
	packet := NewPingReqPacket()
	data, _ := packet.Marshal()
	if len(data) != 2 {
		t.Errorf("Data length wrong should be 2 but was %d", len(data))
	} else {
		if data[0] != 0xc0 ||
			data[1] != 0 {
			t.Errorf("Bytes  are {%d,%d}  but should be {0xc0,0}",
				data[0], data[1])
		}
	}
}

func TestUnmarshalPingReq(t *testing.T) {
	data := []byte{0xc0, 0}
	packet := new(PingReqPacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	}
}
