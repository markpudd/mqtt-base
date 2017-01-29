package mqttbase

// FixedHeader - FixedHeader structure
type FixedHeader struct {
	cntrlPacketType byte
	flags           byte
	remaingLength   uint32
}

// Connect - Constant Connect bit
const Connect byte = 0x10

// Connack - Constant Connack bit
const Connack byte = 0x20

// Publish - Constant Publish bit
const Publish byte = 0x30

// Puback - Constant Puback bit
const Puback byte = 0x40

// Pubrec - Constant Pubrec bit
const Pubrec byte = 0x50

// Pubrel - Constant Pubrel bit
const Pubrel byte = 0x62

// Pubcomp - Constant Pubcomp bit
const Pubcomp byte = 0x70

// Subscribe - Constant Subscribe bit
// Note this is slightly odd as it has a bit 2 set.....
const Subscribe byte = 0x82

// Suback - Constant Suback bit
const Suback byte = 0x90

// Unsubscribe - Constant Unsubscribe bit
const Unsubscribe byte = 0xa2

// Pingreq - Constant Pingreq bit
const Pingreq byte = 0xb0

// Pingresp - Constant Pingresp bit
const Pingresp byte = 0xc0

// Disconnect - Constant Disconnect bit
const Disconnect byte = 0xd0

// Dup - Constant Dup bit
const Dup byte = 0x08

// Qos1 - Constant Qos1 bit
const Qos1 byte = 0x04

// Qos2 - Constant Qos2 bit
const Qos2 byte = 0x02

// Retain - Constant Retain bit
const Retain byte = 0x01

func (fh *FixedHeader) Marshal() []byte {
	lendata := EncodeLength(fh.remaingLength)
	data := make([]byte, 1, len(lendata)+1)
	data[0] = fh.cntrlPacketType | fh.flags
	data = append(data, lendata...)
	return data
}

func (fh *FixedHeader) unmarshal(b []byte) {
	if b[0] == 0x82 {
		fh.cntrlPacketType = 0x82
		fh.flags = 0
	} else {
		fh.cntrlPacketType = b[0] & 0xF0
		fh.flags = b[0] & 0x0F
	}
	fh.remaingLength, _ = UnencodeLength(b[1:])
}
