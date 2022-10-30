package web

import (
	"backend/base/rest"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{}
}

type SessionPost struct {
	ID   *string
	Pass *string
}

func (handler *SessionHandler) POST(context *gin.Context) (interface{}, *rest.HttpError) {
	param := SessionPost{}
	err := context.ShouldBindJSON(&param)
	if err != nil {
		return nil, rest.Err400(err.Error())
	}
	return param, nil
}

func (handler *SessionHandler) GET(context *gin.Context) (interface{}, *rest.HttpError) {
	return SessionPost{}, nil
}

func (handler *SessionHandler) DELETE(context *gin.Context) (interface{}, *rest.HttpError) {
	return nil, nil
}
