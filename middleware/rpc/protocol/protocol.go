package protocol

import (
	"encoding/binary"
	"errors"
	"github.com/BazingaLyn/jarvis/middleware/rpc/codec"
	"github.com/vmihailenco/msgpack"
	"io"
)

var MAGIC = []byte{0xab, 0xba}

type MessageType byte

const (
	MessageTypeRequest MessageType = iota
	MessageTypeResponse
	MessageTypeBeat
)

type CompressType byte

const (
	CompressTypeNone CompressType = iota
)

type Header struct {
	Seq            uint64
	MessageType    MessageType
	CompressType   CompressType
	SerializerType codec.SerializerType
	ServiceName    string
	MethodName     string
	Error          string
	MetaData       map[string]interface{}
}

type Message struct {
	*Header
	Data []byte
}

type Protocol interface {
	NewMessage() *Message
	DecodeMessage(r io.Reader) (*Message, error)
	EncodeMessage(message *Message) []byte
}

type RPCProtocol struct {
}

func (RPCProtocol) NewMessage() *Message {
	return &Message{Header: &Header{}}
}

func (RPCProtocol) DecodeMessage(r io.Reader) (msg *Message, err error) {
	first3bytes := make([]byte, 3)
	_, err = io.ReadFull(r, first3bytes)
	if err != nil {
		return
	}
	if !CheckMagic(first3bytes[:2]) {
		err = errors.New("wrong protocol")
		return
	}

	totalLenBytes := make([]byte, 4)
	_, err = io.ReadFull(r, totalLenBytes)
	if err != nil {
		return
	}
	totalLen := int(binary.BigEndian.Uint32(totalLenBytes))
	if totalLen < 4 {
		err = errors.New("invalid total length")
		return
	}
	data := make([]byte, totalLen)
	_, err = io.ReadFull(r, data)
	headerLen := int(binary.BigEndian.Uint32(data[:4]))
	headerBytes := data[4 : headerLen+4]
	header := &Header{}
	err = msgpack.Unmarshal(headerBytes, header)
	if err != nil {
		return
	}
	msg = new(Message)
	msg.Header = header
	msg.Data = data[headerLen+4:]
	return
}

func CheckMagic(bytes []byte) bool {
	return bytes[0] == MAGIC[0] && bytes[1] == MAGIC[1]
}

//-------------------------------------------------------------------------------------------------
//|2byte|1byte  |4byte       |4byte        | header length |(total length - header length - 4byte)|
//-------------------------------------------------------------------------------------------------
//|magic|version|total length|header length|     header    |                    body              |
//-------------------------------------------------------------------------------------------------
func (RPCProtocol) EncodeMessage(msg *Message) []byte {
	first3bytes := []byte{0xab, 0xba, 0x00}
	headerBytes, _ := msgpack.Marshal(msg.Header)

	totalLen := 4 + len(headerBytes) + len(msg.Data)
	totalLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(totalLenBytes, uint32(totalLen))

	data := make([]byte, totalLen+7)
	start := 0
	copyFullWithOffset(data, first3bytes, &start)
	copyFullWithOffset(data, totalLenBytes, &start)

	headerLenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(headerLenBytes, uint32(len(headerBytes)))
	copyFullWithOffset(data, headerLenBytes, &start)
	copyFullWithOffset(data, headerBytes, &start)
	copyFullWithOffset(data, msg.Data, &start)
	return data

}

func copyFullWithOffset(dst []byte, src []byte, start *int) {
	copy(dst[*start:*start+len(src)], src)
	*start = *start + len(src)
}
