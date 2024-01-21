package middlewares

import (
	"net/http"

	"github.com/JFMajer/rest-api-gin/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	context.Set("userId", userId)
	context.Next()
}
