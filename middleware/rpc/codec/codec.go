package codec

import (
	"bytes"
	"encoding/gob"
	"github.com/vmihailenco/msgpack"
)

type SerializerType byte

const (
	MessagePack SerializerType = iota
	GOB
)

var codec = map[SerializerType]Codec{
	MessagePack: &MessagePackCodec{},
	GOB:         &GobCodec{},
}

type Codec interface {
	Encode(value interface{}) ([]byte, error)
	Decode(data []byte, value interface{}) error
}

type MessagePackCodec struct {
}

func (m *MessagePackCodec) Encode(value interface{}) ([]byte, error) {
	return msgpack.Marshal(value)
}

func (m *MessagePackCodec) Decode(data []byte, value interface{}) error {
	return msgpack.Unmarshal(data, value)
}

type GobCodec struct {
}

func (g *GobCodec) Encode(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(value)
	return buf.Bytes(), err
}

func (g *GobCodec) Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	err := gob.NewDecoder(buf).Decode(value)
	return err
}
