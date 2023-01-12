package resource

import (
	"backend/base/rest"
	"backend/persist/repository"
	"backend/web/util"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	userRepository repository.User
}

func NewSessionHandler() *SessionHandler {
	return &SessionHandler{
		userRepository: repository.NewUser(),
	}
}

type SessionPost struct {
	ID   string `json:"id"`
	Pass string `json:"pass"`
}

func (handler *SessionHandler) POST(context *gin.Context) (interface{}, error) {
	param := SessionPost{}
	err := context.ShouldBindJSON(&param)
	if err != nil {
		return nil, rest.Err400(err.Error())
	}
	user := handler.userRepository.FindByNameAndPass(param.ID, param.Pass)
	if user == nil {
		return nil, rest.Err400("ID/Pass error")
	}
	token, err := util.GenerateToken(user.ID, user.Name)
	if err != nil {
		return nil, rest.Err500(err.Error())
	}
	context.SetCookie(util.JwtKey, token, 3600, "/", "", false, false)
	return "", nil
}

func (handler *SessionHandler) GET(context *gin.Context) (interface{}, error) {
	cookie, err := context.Cookie(util.JwtKey)
	if err != nil {
		return nil, err
	}
	token, err := util.ParseToken(cookie)
	return token, err
}

func (handler *SessionHandler) DELETE(context *gin.Context) (interface{}, error) {
	context.SetCookie(util.JwtKey, "", 3600, "/", "", false, true)
	return nil, nil
}
