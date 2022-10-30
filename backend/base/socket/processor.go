package socket

import (
	"fmt"
	"net"
	"time"
)

type Processor struct {
	Conn        net.Conn
	IsServer    bool
	ConnectTime time.Time
	Status      ProcessorStatus
	handler     Handler
	queue       chan *Message
}

func NewProcessor(conn net.Conn, server bool, handler Handler) *Processor {
	return &Processor{
		Conn:        conn,
		IsServer:    server,
		ConnectTime: time.Now(),
		Status:      Init,
		handler:     handler,
		queue:       make(chan *Message, 8),
	}
}

func (processor *Processor) Send(message *Message) {
	_, err := processor.Conn.Write(message.GetHeader())
	if err != nil {
		processor.close()
		return
	}
	_, err = processor.Conn.Write(message.Data)
	if err != nil {
		processor.close()
		return
	}
}

func (processor *Processor) Start() {
	go processor.receiving()
	log.Infof("Processor[%s] start", processor.Conn.RemoteAddr().String())
}

func (processor *Processor) receiving() {
	for {
		header, err := processor.read(37)
		if err != nil {
			processor.close()
			return
		}
		message := HeadMessage(header)

		data, err := processor.read(int(message.Len))
		if err != nil {
			processor.close()
			return
		}
		message.Data = data
		processor.handler.Accept(message, processor)
	}
}

func (processor *Processor) read(len int) ([]byte, error) {
	buf := make([]byte, len)
	n, err := processor.Conn.Read(buf)
	if err != nil {
		return nil, err
	}

	if n != len { // message error
		return nil, fmt.Errorf("message error: length is %d not %d", n, len)
	}
	return buf, nil
}

func (processor *Processor) close() {
	log.Infof("Processor[%s] close", processor.Conn.RemoteAddr().String())
	processor.Status = Closing
	_ = processor.Conn.Close()
	processor.Status = Closed
}
