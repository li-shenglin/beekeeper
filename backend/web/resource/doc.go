package resource

import (
	"backend/persist/repository"

	"github.com/gin-gonic/gin"
)

type DocHandler struct {
	DocRepository repository.Doc
}

func NewDocHandler() *DocHandler {
	return &DocHandler{
		DocRepository: repository.NewDoc(),
	}
}

func (handler *DocHandler) POST(context *gin.Context) (interface{}, error) {
	return "", nil
}

func (handler *DocHandler) PUT(context *gin.Context) (interface{}, error) {
	return "", nil
}

func (handler *DocHandler) GET(context *gin.Context) (interface{}, error) {
	return SessionPost{}, nil
}

func (handler *DocHandler) DELETE(context *gin.Context) (interface{}, error) {
	return nil, nil
}

type DocsHandler struct {
	DocRepository repository.Doc
}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{
		DocRepository: repository.NewDoc(),
	}
}

func (handler *DocsHandler) GET(context *gin.Context) (interface{}, error) {
	return SessionPost{}, nil
}
