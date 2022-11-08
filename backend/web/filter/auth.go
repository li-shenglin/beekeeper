package filter

import (
	"backend/web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.URL.Path == "/api/session" && context.Request.Method == "POST" {
			context.Next()
			return
		}
		token, _ := context.Cookie("beekeeper-auth")
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"isSuccess": false,
				"error":     "Unauthorized",
			})
			context.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"isSuccess": false,
					"error":     "invalid token ",
				})
				context.Abort()
				return
			}
			err = claims.Valid()
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{
					"isSuccess": false,
					"error":     "token err:" + err.Error(),
				})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
