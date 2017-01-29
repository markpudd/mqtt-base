package mqtt

import (
	"testing"
)

type TestProcessor struct {
	buffer         []byte
	length         int
	readPos        int
	recievedPacket Packet
}

func (tp *TestProcessor) GetByte() (byte, error) {
	ret := tp.buffer[tp.readPos]
	tp.readPos++
	if tp.readPos == tp.length {
		tp.readPos = 0
	}
	return ret, nil
}

func (tp *TestProcessor) Process(p *Packet) bool {
	tp.recievedPacket = *p
	return false
}

func TestRecievePublishPacket(t *testing.T) {
	tp := new(TestProcessor)
	tp.buffer = []byte{0x30, 16, 0, 3, 'a', '/', 'b', 0x3a, 0xcd, 'T', 'E',
		'S', 'T', '_', 'D', 'A', 'T', 'A'}
	tp.length = len(tp.buffer)
	tp.readPos = 0
	dr := NewDataReader(tp)
	for _, b := range tp.buffer {
		dr.RecieveByte(b)
	}

	packetIf := tp.recievedPacket
	if packetIf == nil || packetIf.PacketType() != Publish {
		t.Errorf("Packet not publish")
	} else {
		var packet *PublishPacket
		packet = packetIf.(*PublishPacket)
		if packet.fixedHeader == nil {
			t.Errorf("Fixed header nil")
		} else {
			if packet.fixedHeader.cntrlPacketType != Publish {
				t.Errorf("Packet not Publish")
			}
		}
		if packet.id != 0x3acd {
			t.Errorf("Packet id is %#x but should 0x3acd", packet.id)
		}
		if packet.topicName != "a/b" {
			t.Errorf("Packet topic name is %s but should be a/b", packet.topicName)
		}
		if len(packet.data) != 9 {
			t.Errorf("Packet data length should be 9 but is %d", len(packet.data))
		} else {
			expected := []byte("TEST_DATA")
			for i := 0; i < 9; i++ {
				if packet.data[i] != expected[i] {
					t.Errorf("Expected %#x at %d but got %#x",
						expected[i], i, packet.data[i])
				}
			}
		}
	}
}

func TestUnmarshalPublishPacket(t *testing.T) {
	buffer := []byte{0x30, 16, 0, 3, 'a', '/', 'b', 0x3a, 0xcd, 'T', 'E',
		'S', 'T', '_', 'D', 'A', 'T', 'A'}
	packet, _ := Unmarshal(buffer)
	if (*packet).PacketType() != Publish {
		t.Errorf("Packet not publish")
	}
}

func TestUnmarshalConnectPacket(t *testing.T) {
	buffer := []byte{21, 14, 0, 4, 'M', 'Q', 'T', 'T', 4, 0, 1, 2, 0, 4, 't', 'e', 's', 't'}
	packet, _ := Unmarshal(buffer)
	if (*packet).PacketType() != Connect {
		t.Errorf("Packet not publish")
	}
}

func TestUnmarshalConnackPacket(t *testing.T) {
	buffer := []byte{32, 2, 0, 3}
	packet, _ := Unmarshal(buffer)
	if (*packet).PacketType() != Connack {
		t.Errorf("Packet not publish")
	}
}

func TestUnmarshalPubackPacket(t *testing.T) {
	buffer := []byte{0x40, 2, 0x3a, 0xcd}
	packet, _ := Unmarshal(buffer)
	if (*packet).PacketType() != Puback {
		t.Errorf("Packet not publish")
	}
}
