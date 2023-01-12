package resource

import (
	"backend/base/rest"
	"backend/persist/model"
	"backend/persist/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository repository.User
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRepository: repository.NewUser(),
	}
}

func (handler *UserHandler) GET(context *gin.Context) (interface{}, error) {
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("userID is required")
	}
	return handler.userRepository.GetUser(id)
}

func (handler *UserHandler) DELETE(context *gin.Context) (interface{}, error) {
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("userID is required")
	}
	return nil, handler.userRepository.DeleteUser(id)
}

func (handler *UserHandler) PUT(context *gin.Context) (interface{}, error) {
	var form model.User
	if err := context.ShouldBind(&form); err != nil {
		return nil, err
	}
	id := rest.ParamID(context)
	if id == 0 {
		return nil, rest.Err400("userID is required")
	}
	form.ID = id
	return handler.userRepository.CreateUser(&form)
}

type UsersHandler struct {
	userRepository repository.User
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{
		userRepository: repository.NewUser(),
	}
}

func (handler *UsersHandler) GET(context *gin.Context) (interface{}, error) {
	return SessionPost{}, nil
}

func (handler *UsersHandler) POST(context *gin.Context) (interface{}, error) {
	var form model.User
	if err := context.ShouldBind(&form); err != nil {
		return nil, err
	}
	return handler.userRepository.CreateUser(&form)
}
