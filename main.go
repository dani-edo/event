package main

import (
	"edo.com/event/db"
	"edo.com/event/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
