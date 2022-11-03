package socket

type Handler interface {
	Accept(param Param) (any, error)
}
