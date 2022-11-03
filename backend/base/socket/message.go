package socket

import (
	"backend/common"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type Param interface {
	Bind(v any) error
	String() string
}

type Message struct {
	Len   int32
	SeqID []byte
	Data  []byte
	Type  byte
}

func NewMessage(data []byte) *Message {
	msg := &Message{
		SeqID: common.UUID(),
		Data:  data,
		Len:   int32(len(data)),
		Type:  0,
	}
	return msg
}

func HeadMessage(header []byte) *Message {
	msg := &Message{}
	common.PanicNotNull(msg.SetHeader(header))
	return msg
}

func NewReturnMessage(seqID, data []byte) *Message {
	return &Message{
		SeqID: seqID,
		Data:  data,
		Type:  1,
		Len:   int32(len(data)),
	}
}

func (message *Message) GetData() (*Parameter, error) {
	parameter := &Parameter{}
	err := json.Unmarshal(message.Data, parameter)
	return parameter, err
}

func (message *Message) GetHeader() []byte {
	bt := make([]byte, 37)
	lenByte := message.intToBytes(message.Len)
	copy(bt[:], lenByte)
	copy(bt[4:], message.SeqID)
	bt[36] = message.Type
	return bt
}

func (message *Message) SetHeader(buf []byte) error {
	if len(buf) < 37 {
		return fmt.Errorf("error header")
	}
	message.Len = int32(binary.LittleEndian.Uint32(buf[0:4]))
	message.SeqID = buf[4:36]
	message.Type = buf[36]
	return nil
}

func (message *Message) intToBytes(intNum int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	common.PanicNotNull(binary.Write(buf, binary.LittleEndian, uint32(intNum)))
	return buf.Bytes()
}

type Parameter struct {
	Opt  int32
	Data []byte
	Err  error
}

func (parameter *Parameter) Bind(v any) error {
	return json.Unmarshal(parameter.Data, v)
}

func (parameter *Parameter) String() string {
	return fmt.Sprintf("opt: %d, err: %s, data: %s", parameter.Opt, parameter.Err, string(parameter.Data))
}

func (parameter *Parameter) SetErr(err error) *Parameter {
	parameter.Data = nil
	parameter.Err = err
	return parameter
}

func (parameter *Parameter) SetData(d any) *Parameter {
	marshal, err := json.Marshal(d)
	if err != nil {
		return parameter.SetErr(err)
	}
	parameter.Data = marshal
	parameter.Err = nil
	return parameter
}
