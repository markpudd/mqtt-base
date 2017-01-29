package mqttbase

import (
	"testing"
)

func TestPublishPacketConstructor(t *testing.T) {
	packet := NewPublishPacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.fixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestMarshalPublish(t *testing.T) {
	packet := NewPublishPacket()
	packet.topicName = "a/b"
	packet.id = 0x3acd
	packet.data = []byte("TEST_DATA")
	data, _ := packet.Marshal()
	if len(data) != 18 {
		t.Errorf("Data length wrong should be 18 but was %d", len(data))
	} else {
		expected := []byte{0x30, 16, 0, 3, 'a', '/', 'b', 0x3a, 0xcd, 'T', 'E',
			'S', 'T', '_', 'D', 'A', 'T', 'A'}
		for i := 0; i < 8; i++ {
			if data[i] != expected[i] {
				t.Errorf("Expected %#x at %d but got %#x",
					expected[i], i, data[i])
			}
		}

	}
}

func TestUnmarshalPublish(t *testing.T) {
	data := []byte{0x30, 16, 0, 3, 'a', '/', 'b', 0x3a, 0xcd, 'T', 'E',
		'S', 'T', '_', 'D', 'A', 'T', 'A'}
	packet := new(PublishPacket)
	packet.unmarshal(data)
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
