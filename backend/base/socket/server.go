package socket

import (
	"fmt"
	"net"
)

type Server struct {
	*Application
	port int32
}

func (server *Server) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.port))
	if err != nil {
		return err
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Accept err=", err)
		} else {
			server.Add(conn)
		}
	}
}

func NewServer(port int32) *Server {
	return &Server{
		Application: GetApplication(),
		port:        port,
	}
}
