package routes

import (
	middlewares "github.com/DMohanty99/eventApp/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", HandleSignUp)
	server.POST("/signin", HandleSignIn)
	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)
	authenticatedRoute.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
