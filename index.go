package main

import (
	"github.com/DMohanty99/eventApp/db"
	"github.com/DMohanty99/eventApp/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
