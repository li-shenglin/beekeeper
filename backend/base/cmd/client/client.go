package main

import (
	"backend/base/log"
	"backend/base/socket"

	"github.com/sirupsen/logrus"
)

func main() {
	log.Config("", logrus.DebugLevel)
	app := socket.NewClient("127.0.0.1", 9111)
	app.Run()
}
