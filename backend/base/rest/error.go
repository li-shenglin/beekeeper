package rest

import (
	"fmt"
	h "net/http"
)

type HttpError struct {
	code int
	msg  string
}

func (err *HttpError) Error() string {
	return err.msg
}

func (err *HttpError) String() string {
	return fmt.Sprintf("{code: %d, msg: \"%s\"}", err.code, err.msg)
}

func Err400(msg string) *HttpError {
	return &HttpError{
		code: h.StatusBadRequest,
		msg:  msg,
	}
}

func Err401() *HttpError {
	return &HttpError{
		code: h.StatusUnauthorized,
		msg:  "Unauthorized",
	}
}

func Err403() *HttpError {
	return &HttpError{
		code: h.StatusForbidden,
		msg:  "Forbidden",
	}
}

func Err500(msg string) *HttpError {
	return &HttpError{
		code: h.StatusInternalServerError,
		msg:  msg,
	}
}
