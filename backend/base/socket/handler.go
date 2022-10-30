package socket

import "github.com/sirupsen/logrus"

type Handler interface {
	Accept(message *Message, processor *Processor)
}

type ChainHandler struct {
}

func (h *ChainHandler) Accept(message *Message, processor *Processor) {
	log.WithFields(logrus.Fields{
		"SeqID": string(message.SeqID),
		"Len":   message.Len,
	}).Infof("recieved message")
}
