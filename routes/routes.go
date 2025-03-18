package routes

import (
	"edo.com/event/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	autenticated := server.Group("/")
	autenticated.Use(middlewares.Authenticate)
	autenticated.POST("/events", middlewares.Authenticate, createEvent)
	autenticated.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	autenticated.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
