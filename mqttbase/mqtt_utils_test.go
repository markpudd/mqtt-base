package mqttbase

import (
	"testing"
)

func TestEncodeLength(t *testing.T) {
	r := EncodeLength(64)
	if len(r) != 1 || r[0] != 0x40 {
		t.Errorf("Unable to encode 64")
	}
	r = EncodeLength(321)
	if len(r) != 2 {
		t.Errorf("Unable to encode 321 - array %d when should be 2", len(r))
	} else if r[0] != 193 || r[1] != 2 {
		t.Errorf("Unable to encode 321 array is %d,%d but should be 193,2", r[0], r[1])
	}
}

func TestUnencodeLength(t *testing.T) {
	r, len := UnencodeLength([]byte{64})
	if r != 64 {
		t.Errorf("Unable to unencode 64 got %d", r)
	}
	if !len {
		t.Errorf("Data still in buffer but number found")
	}
	r, len = UnencodeLength([]byte{193, 2})
	if r != 321 {
		t.Errorf("Unable to unencode 321 got %d", r)
	}
	if !len {
		t.Errorf("Data still in buffer but number found")
	}
}

func TestEncodeASCIIString(t *testing.T) {
	data, err := EncodeString("TEST")
	if err != nil {
		t.Errorf("Bad string %s", err)
	}
	if len(data) != 6 {
		t.Errorf("String data is %d but should be 6", len(data))
	}
	if data[0] != 0 || data[1] != 4 {
		t.Errorf("String length is %d,%d but should be 0,4", data[0], data[1])
	}
	if data[2] != 'T' ||
		data[3] != 'E' ||
		data[4] != 'S' ||
		data[5] != 'T' {
		t.Errorf("Bytes 3-6 should are {%d,%d,%d,%d}  but should be {T,E,S,T}",
			data[2], data[3], data[4], data[5])
	}
}

func TestUnencodeASCIIString(t *testing.T) {
	data := []byte{0, 4, 'T', 'E', 'S', 'T'}
	str, err := UnencodeString(data)
	if err != nil {
		t.Errorf("Bad string %s", err)
	}
	if str != "TEST" {
		t.Errorf("String is %s but should be TEST", str)
	}
}

func TestUnencodeBadString(t *testing.T) {
	data := []byte{0, 5, 'T', 'E', 'S', 'T'}
	str, err := UnencodeString(data)
	if err == nil {
		t.Errorf("Bad string but no error %s", str)
	}
}
