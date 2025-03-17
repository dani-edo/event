package routes

import (
	"net/http"

	"edo.com/event/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) // bind the request body to the event struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not register", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
