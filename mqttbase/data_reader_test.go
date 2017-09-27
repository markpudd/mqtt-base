package mqttbase

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
		if packet.FixedHeader == nil {
			t.Errorf("Fixed header nil")
		} else {
			if packet.FixedHeader.cntrlPacketType != Publish {
				t.Errorf("Packet not Publish")
			}
		}
		if packet.Id != 0x3acd {
			t.Errorf("Packet id is %#x but should 0x3acd", packet.Id)
		}
		if packet.TopicName != "a/b" {
			t.Errorf("Packet topic name is %s but should be a/b", packet.TopicName)
		}
		if len(packet.Data) != 9 {
			t.Errorf("Packet data length should be 9 but is %d", len(packet.Data))
		} else {
			expected := []byte("TEST_DATA")
			for i := 0; i < 9; i++ {
				if packet.Data[i] != expected[i] {
					t.Errorf("Expected %#x at %d but got %#x",
						expected[i], i, packet.Data[i])
				}
			}
		}
	}
}

func TestUnmarshalPublishPacket(t *testing.T) {
	buffer := []byte{0x30, 16, 0, 3, 'a', '/', 'b', 0x3a, 0xcd, 'T', 'E',
		'S', 'T', '_', 'D', 'A', 'T', 'A'}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Publish {
		t.Errorf("Packet not publish")
	}
}

func TestUnmarshalConnectPacket(t *testing.T) {
	buffer := []byte{21, 14, 0, 4, 'M', 'Q', 'T', 'T', 4, 0, 1, 2, 0, 4, 't', 'e', 's', 't'}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Connect {
		t.Errorf("Packet not connect")
	}
}

func TestUnmarshalConnackPacket(t *testing.T) {
	buffer := []byte{32, 2, 0, 3}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Connack {
		t.Errorf("Packet not connectack")
	}
}

func TestUnmarshalPubackPacket(t *testing.T) {
	buffer := []byte{0x40, 2, 0x3a, 0xcd}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Puback {
		t.Errorf("Packet not publish")
	}
}

func TestUnmarshalPingreqPacket(t *testing.T) {
	buffer := []byte{0xc0, 0}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Pingreq {
		t.Errorf("Packet not Pingreq")
	}
}

func TestUnmarshalPingrespPacket(t *testing.T) {
	buffer := []byte{0xd0, 0}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Pingresp {
		t.Errorf("Packet not Pingresp")
	}
}

func TestUnmarshalPubrecPacket(t *testing.T) {
	buffer := []byte{0x50, 0x02, 0xa6, 0xf2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Pubrec {
		t.Errorf("Packet not Pingrec")
	}
}

func TestUnmarshalPubrelPacket(t *testing.T) {
	buffer := []byte{0x62, 0x02, 0xa6, 0xf2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Pubrel {
		t.Errorf("Packet not Pubrel")
	}
}

func TestUnmarshalPubcompPacket(t *testing.T) {
	buffer := []byte{0x70, 0x02, 0xa6, 0xf2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Pubcomp {
		t.Errorf("Packet not Pubcomp")
	}
}

func TestUnmarshalSubscribePacket(t *testing.T) {
	buffer := []byte{0x82, 13, 0x5a, 0x22, 0, 3, 'e', '/', 'f', 1, 0, 2, 'g', 'h', 2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Subscribe {
		t.Errorf("Packet not Subscribe")
	}
}

func TestUnmarshalSubackPacket(t *testing.T) {
	buffer := []byte{0x90, 0x06, 0xa6, 0xf2, 0x00, 0x01, 0x02, 0xf0}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Suback {
		t.Errorf("Packet not Suback")
	}
}

func TestUnmarshalUnsubscribePacket(t *testing.T) {
	buffer := []byte{0xa2, 13, 0x5a, 0x22, 0, 3, 'e', '/', 'f', 1, 0, 2, 'g', 'h', 2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != Unsubscribe {
		t.Errorf("Packet not Unsubscribe")
	}
}

func TestUnmarshalUnSubackPacket(t *testing.T) {
	buffer := []byte{0xb0, 0x02, 0xa6, 0xf2}
	packet, err := Unmarshal(buffer)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if (*packet).PacketType() != UnSuback {
		t.Errorf("Packet not Unsuback")
	}
}
