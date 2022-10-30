package main

import (
	"backend/base/log"
	"backend/base/socket"
	"backend/common"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Config("", logrus.DebugLevel)
	server := socket.NewServer(9111)
	common.PanicNotNull(server.Run())
}
