package routes

import (
	"fmt"
	"net/http"

	"github.com/DMohanty99/eventApp/models"
	"github.com/gin-gonic/gin"
)

func HandleSignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data", "actualerr": err})
		fmt.Println(user.Email)
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "object created successfully"})

}
