package rest

import (
	"backend/web/util"

	"github.com/gin-gonic/gin"
)

func ParamID(context *gin.Context) uint {
	return uint(util.ToInt64Or(context.Param("id"), 0))
}
