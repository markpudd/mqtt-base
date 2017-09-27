package mqttbase

import (
	"errors"
	"fmt"
)

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
		packet, err := Unmarshal(dr.data)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		_ = dr.processor.Process(packet)
		dr.data = dr.data[:0]
		dr.length = MaxPacketSize
		dr.pos = 0
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
	case Connect:
		packet = NewConnectPacket()
		packet.unmarshal(data)
		break
	case Connack:
		packet = NewConnackPacket()
		packet.unmarshal(data)
		break
	case Publish:
		packet = NewPublishPacket()
		packet.unmarshal(data)
		break
	case Puback:
		packet = NewPubackPacket()
		packet.unmarshal(data)
		break
	case Pubrec:
		packet = NewPubrecPacket()
		packet.unmarshal(data)
		break
	case Pubrel & 0xf0:
		packet = NewPubrelPacket()
		packet.unmarshal(data)
		break
	case Pubcomp:
		packet = NewPubcompPacket()
		packet.unmarshal(data)
		break
	case Subscribe & 0xf0:
		packet = NewSubscribePacket()
		packet.unmarshal(data)
		break
	case Suback:
		packet = NewSubackPacket()
		packet.unmarshal(data)
		break
	case Pingreq:
		packet = NewPingReqPacket()
		packet.unmarshal(data)
		break
	case Pingresp:
		packet = NewPingRespPacket()
		packet.unmarshal(data)
		break
	case Unsubscribe & 0xf0:
		packet = NewUnsubscribePacket()
		packet.unmarshal(data)
		break
	case UnSuback:
		packet = NewUnSubackPacket()
		packet.unmarshal(data)
		break
	default:
		fmt.Printf("Unsuported %#x", data[0])
		return nil, errors.New("Unsupported packet type")
	}
	return &packet, nil

}
