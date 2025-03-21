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
	autenticated.POST("/events", createEvent)
	autenticated.PUT("/events/:id", updateEvent)
	autenticated.DELETE("/events/:id", deleteEvent)
	autenticated.POST("/events/:id/register", registerForEvent)
	autenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
