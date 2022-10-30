package socket

import (
	LOG "backend/base/log"
	"fmt"
	"net"
	"time"
)

var log = LOG.GetLog()

type Server struct {
	port          int32
	processorList []*Processor
}

func (server *Server) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))
	if err != nil {
		return err
	}
	defer listen.Close()
	go server.checkProcessor()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Accept err=", err)
		} else {
			processor := NewProcessor(conn, true, &ChainHandler{})
			server.processorList = append(server.processorList, processor)
			processor.Start()
		}
	}
}

func (server *Server) checkProcessor() {
CHECK:
	time.After(time.Second * 10)
	processors := make([]*Processor, 0)
	for i := range server.processorList {
		if server.processorList[i].Status != Closed {
			processors = append(processors, server.processorList[i])
		}
	}
	server.processorList = processors
	goto CHECK
}

func NewServer(port int32) *Server {
	return &Server{
		port: port,
	}
}
