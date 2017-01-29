package mqttbase

import "errors"

// MaxPacketSize - maximum packet size
const MaxPacketSize int = 65000

type DataReader struct {
	data      []byte
	length    int
	pos       int
	processor Processor
}

// Processor = processes packet
type Processor interface {
	Process(p *Packet) bool
}

// NewDataReader = new data reader
func NewDataReader(p Processor) *DataReader {
	dr := new(DataReader)
	dr.data = make([]byte, 0, MaxPacketSize)
	dr.length = MaxPacketSize
	dr.processor = p
	return dr
}

// RecieveByte - recieves packet
func (dr *DataReader) RecieveByte(b byte) {

	dr.data = append(dr.data, b)
	if dr.pos > 0 && dr.length == MaxPacketSize {
		l, success := UnencodeLength(dr.data[1:])
		if success {
			dr.length = int(l) + 2
		}
	}
	dr.pos++
	if dr.pos == dr.length {
		packet, _ := Unmarshal(dr.data)
		_ = dr.processor.Process(packet)
		dr.data = dr.data[:0]
		dr.length = MaxPacketSize
	}
}

// Unmarshal data, this assumes we have already got read the fixed header
// buffer and entire buffer
func Unmarshal(data []byte) (*Packet, error) {
	var packet Packet
	// check first four bits of data to get packet type
	if len(data) == 0 {
		return nil, errors.New("Zero length data packet")
	}
	switch data[0] & 0xF0 {
	case 0x10:
		packet = NewConnectPacket()
		packet.unmarshal(data)
		break
	case 0x20:
		packet = NewConnackPacket()
		packet.unmarshal(data)
		break
	case 0x30:
		packet = NewPublishPacket()
		packet.unmarshal(data)
		break
	case 0x40:
		packet = NewPubackPacket()
		packet.unmarshal(data)
		break
	default:
		return nil, errors.New("Unsupported packet type")
	}
	return &packet, nil

}
