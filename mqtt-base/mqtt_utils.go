package mqtt

import (
	"errors"
)

// EncodeString -  encodes a string for wire
func EncodeString(str string) ([]byte, error) {
	length := len(str)
	data := make([]byte, 2, length+2)
	data[0] = byte(length & 0xFF00 >> 16)
	data[1] = byte(length & 0xFF)
	data = append(data, str...)
	return data, nil
}

// UnencodeString -  unencodes a string from wire
func UnencodeString(data []byte) (string, error) {
	length := int(data[0])<<16 | int(data[1])
	if length > len(data)-2 {
		return "", errors.New("String length wrong")
	}
	return string(data[2 : 2+length]), nil
}

// EncodeLength -  encodes  length
func EncodeLength(i uint32) []byte {
	ret := make([]byte, 0, 4)
	for i > 0 && len(ret) <= 4 {
		if i < 128 {
			ret = append(ret, byte(i&127))
		} else {
			ret = append(ret, byte(byte(i&0x7f|0x80)))
		}
		i = i >> 7
	}
	if len(ret) > 4 {
		return []byte{}
	}
	return ret
}

// UnencodeLength -  unencodes  length
func UnencodeLength(b []byte) (uint32, bool) {
	var ret uint32
	var shift uint32
	i := 0
	for i <= len(b) && b[i]&0x80 > 0 {
		ret = ret + (uint32(b[i]&0x7f) << shift)
		shift = shift + 7
		i++
	}
	return ret + (uint32(b[i]&0x7f) << shift), i <= len(b)
}
