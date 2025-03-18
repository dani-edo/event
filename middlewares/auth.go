package middlewares

import (
	"net/http"

	"edo.com/event/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"}) // AbortWithStatusJSON stops other requests from continuing
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	context.Set("userId", userId) // Set is a method that sets a value in the context, the data available to the next handler
	context.Next()
}
