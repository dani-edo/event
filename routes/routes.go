package routes

import (
	"edo.com/event/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
