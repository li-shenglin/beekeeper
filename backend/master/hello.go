package master

import (
	"backend/base/socket"
)

type HelloHandler struct {
}

func (h *HelloHandler) Accept(param socket.Param) (any, error) {
	log.Infof("Get：%s", param.String())
	return "Copy", nil
}
