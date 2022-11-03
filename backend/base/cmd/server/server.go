package main

import (
	"backend/base/log"
	"backend/base/socket"
	"backend/common"
	"backend/master"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Config("", logrus.DebugLevel)
	server := socket.NewServer(9111)
	server.Router(0, &master.HelloHandler{})
	common.PanicNotNull(server.Run())
}
