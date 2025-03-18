package routes

import (
	"net/http"

	"edo.com/event/models"
	"edo.com/event/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) // bind the request body to the event struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not login", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not login", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate the user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
