package socket

import (
	"backend/common"
	"fmt"
	"net"
	"time"
)

type Client struct {
	port          int32
	ip            string
	processorList []*Processor
}

func (client *Client) Run() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ip, client.port))
	if err != nil {
		return
	}
	processor := NewProcessor(conn, true, &ChainHandler{})
	client.processorList = append(client.processorList, processor)
	processor.Start()
	data := []byte{1, 3, 4}
	processor.Send(&Message{
		Len:   int32(len(data)),
		SeqID: common.UUID(),
		Data:  data,
		Type:  0,
	})
	client.checkProcessor()
}

func (client *Client) checkProcessor() {
CHECK:
	time.After(time.Second * 10)
	processors := make([]*Processor, 0)
	for i := range client.processorList {
		if client.processorList[i].Status != Closed {
			processors = append(processors, client.processorList[i])
		}
	}
	client.processorList = processors
	goto CHECK
}

func NewClient(ip string, port int32) *Client {
	return &Client{
		ip:   ip,
		port: port,
	}
}
