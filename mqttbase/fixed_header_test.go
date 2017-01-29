package mqttbase

import (
	"testing"
)

func TestMarshalHeader(t *testing.T) {
	fh := new(FixedHeader)
	fh.cntrlPacketType = Connect
	fh.flags = Retain | Qos1
	fh.remaingLength = 321

	data := fh.Marshal()

	if len(data) != 3 {
		t.Errorf("Data length %d when should be 3", len(data))
	} else {
		n := EncodeLength(321)
		if data[1] != n[0] {
			t.Errorf("First byte of length %d when should be %d", data[1], n[0])
		}
		if data[2] != n[1] {
			t.Errorf("Second byte of length %d when should be %d", data[2], n[1])
		}
	}
	if data[0] != (Connect | Retain | Qos1) {
		t.Errorf("First byte   %d when should be %d", data[0], (Connect | Retain | Qos1))
	}

}

func TestUnmarshalHeader(t *testing.T) {
	data := []byte{21, 193, 2}
	fh := new(FixedHeader)
	fh.unmarshal(data)

	if fh.cntrlPacketType != Connect {
		t.Errorf("Packet type %d when should be %d", fh.cntrlPacketType, Connect)
	}
	if fh.flags != Retain|Qos1 {
		t.Errorf("Flags %d when should be %d", fh.flags, Retain|Qos1)
	}
	if fh.remaingLength != 321 {
		t.Errorf("Remaing Length %d when should be 321", fh.remaingLength)
	}

}
