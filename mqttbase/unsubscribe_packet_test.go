package mqttbase

import (
	"testing"
)

func TestUnsubscribePacketConstructor(t *testing.T) {
	packet := NewUnsubscribePacket()
	if packet == nil {
		t.Errorf("Packet is nil")
	} else {
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header is nil")
		}
	}
}

func TestAddTopicUnsub(t *testing.T) {
	packet := NewSubscribePacket()
	topic1 := new(TopicFilter)
	topic1.Filter = "a/b"
	topic1.Qos = Qos1
	packet.AddTopic(topic1)
	if len(packet.Topics) != 1 || packet.Topics[0] != topic1 {
		t.Errorf("Topic added but not in list")
	}
}

func TestMarshalUnsubscribe(t *testing.T) {
	packet := NewUnsubscribePacket()
	topic1 := new(TopicFilter)
	topic1.Filter = "a/b"
	topic1.Qos = Retain
	topic2 := new(TopicFilter)
	topic2.Filter = "c/d"
	topic2.Qos = Qos2

	packet.AddTopic(topic1)
	packet.AddTopic(topic2)
	packet.Id = 0x3a74

	data, _ := packet.Marshal()
	if len(data) != 16 {
		t.Errorf("Data lentgh wrong should be 16 but was %d", len(data))
	} else {
		if data[0] != 0xa2 {
			t.Errorf("Packet should start with id 0xa2 but starts with %#x",
				data[0])
		}
		if data[1] != 14 {
			t.Errorf("Packet length should be 14 but is %d",
				data[1])
		}
		if data[2] != 0x3a || data[3] != 0x74 {
			t.Errorf("ID Bytes  are {%#x,%#x}  but should be {0x3a,0x74}",
				data[2], data[3])
		}
		if data[4] != 0 || data[5] != 3 {
			t.Errorf("Length of topic 1  is {%#x,%#x}  but should be {0x0,0x3}",
				data[4], data[5])
		}
		if data[6] != 'a' || data[7] != '/' || data[8] != 'b' {
			t.Errorf("Topic 1  is {%c%c%c}  but should be {'a/b}",
				data[6], data[7], data[8])
		}
		if data[9] != 1 {
			t.Errorf("QOS of topic 1  is {%#x}  but should be {0x1}",
				data[9])
		}

		if data[10] != 0 || data[11] != 3 {
			t.Errorf("Length of topic 2  is {%#x,%#x}  but should be {0x0,0x3}",
				data[10], data[11])
		}
		if data[12] != 'c' || data[13] != '/' || data[14] != 'd' {
			t.Errorf("Topic 2  is {%c%c%c}  but should be {'a/b}",
				data[12], data[13], data[14])
		}
		if data[15] != 2 {
			t.Errorf("QOS of topic 2  is {%#x}  but should be {0x2}",
				data[15])
		}
	}
}

func TestMarshalUnsubscribeOneTopic(t *testing.T) {
	packet := NewSubscribePacket()
	topic1 := new(TopicFilter)
	topic1.Filter = "a/b"
	topic1.Qos = Retain

	packet.AddTopic(topic1)
	packet.Id = 0x3a74

	data, _ := packet.Marshal()
	if len(data) != 10 {
		t.Errorf("Data lentgh wrong should be 10 but was %d", len(data))
	} else {
		if data[0] != 0x82 {
			t.Errorf("Packet should start with id 0x82 but starts with %#x",
				data[0])
		}
		if data[1] != 8 {
			t.Errorf("Packet length should be 8 but is %d",
				data[1])
		}
		if data[2] != 0x3a || data[3] != 0x74 {
			t.Errorf("ID Bytes  are {%#x,%#x}  but should be {0x3a,0x74}",
				data[2], data[3])
		}
		if data[4] != 0 || data[5] != 3 {
			t.Errorf("Length of topic 1  is {%#x,%#x}  but should be {0x0,0x3}",
				data[4], data[5])
		}
		if data[6] != 'a' || data[7] != '/' || data[8] != 'b' {
			t.Errorf("Topic 1  is {%c%c%c}  but should be {'a/b}",
				data[6], data[7], data[8])
		}
		if data[9] != 1 {
			t.Errorf("QOS of topic 1  is {%#x}  but should be {0x1}",
				data[9])
		}
	}
}

func TestUnmarshalUnsubscribe(t *testing.T) {
	data := []byte{0xa2, 13, 0x5a, 0x22, 0, 3, 'e', '/', 'f', 1, 0, 2, 'g', 'h', 2}
	packet := new(UnsubscribePacket)
	packet.unmarshal(data)
	if packet.FixedHeader == nil {
		t.Errorf("Fixed header nil")
	} else {
		if packet.FixedHeader.cntrlPacketType != Unsubscribe {
			t.Errorf("Packet not Unubscibe %d", packet.FixedHeader.cntrlPacketType)
		}
	}
	if packet.Id != 0x5a22 {
		t.Errorf("Packet id is %#x but should be 0x5a22", packet.Id)
	}
	if len(packet.Topics) != 2 {
		t.Errorf("Topics count is %d but should be 2", len(packet.Topics))
	} else {
		if packet.Topics[0].Filter != "e/f" {
			t.Errorf("Topics filter 0 is not e/f")
		}
		if packet.Topics[0].Qos != Retain {
			t.Errorf("Topics qos 0 is not Retain")
		}
		if packet.Topics[1].Filter != "gh" {
			t.Errorf("Topics filter 1 is not gh")
		}
		if packet.Topics[1].Qos != Qos2 {
			t.Errorf("Topics qos 1 is not Qos1")
		}

	}
}
