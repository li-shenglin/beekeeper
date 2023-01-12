package resource

import (
	"backend/base/rest"
	"backend/persist/model"
	"backend/persist/repository"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectRepository repository.Project
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		projectRepository: repository.NewProject(),
	}
}

func (handler *ProjectHandler) PUT(context *gin.Context) (interface{}, error) {
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("projectID is required")
	}
	var form model.Project
	err := context.BindJSON(&form)
	if err != nil {
		return nil, err
	}
	form.ID = id
	return "", handler.projectRepository.UpdateProject(&form)
}

func (handler *ProjectHandler) GET(context *gin.Context) (interface{}, error) {
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("projectID is required")
	}
	return handler.projectRepository.GetProject(id)
}

func (handler *ProjectHandler) DELETE(context *gin.Context) (interface{}, error) {
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("projectID is required")
	}
	return handler.projectRepository.DelProject(int64(id)), nil
}

type ProjectsHandler struct {
	projectRepository repository.Project
}

func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{
		projectRepository: repository.NewProject(),
	}
}

func (handler *ProjectsHandler) GET(context *gin.Context) (interface{}, error) {
	return SessionPost{}, nil
}

func (handler *ProjectsHandler) POST(context *gin.Context) (interface{}, error) {
	return "", nil
}
