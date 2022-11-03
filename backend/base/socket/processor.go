package socket

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Processor struct {
	Conn        net.Conn
	ConnectTime time.Time
	instance    *Instance
	lock        *sync.Mutex
}

func NewProcessor(conn net.Conn, instance *Instance) *Processor {
	return &Processor{
		Conn:        conn,
		ConnectTime: time.Now(),
		instance:    instance,
		lock:        &sync.Mutex{},
	}
}

func (processor *Processor) Send(message *Message) {
	processor.lock.Lock()
	defer processor.lock.Unlock()
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
	log.Debugf("Processor[%s] start", processor.Conn.RemoteAddr())
	processor.instance.Status = InstanceOnline
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
		processor.instance.Accept(message, processor)
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
	log.Debugf("Processor[%s] close", processor.Conn.RemoteAddr())
	_ = processor.Conn.Close()
	processor.instance.Status = InstanceUnHealth
	processor.instance.processor = nil
}
