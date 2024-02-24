package middlewares

import (
	"net/http"

	"github.com/DMohanty99/eventApp/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
