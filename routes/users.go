package routes

import (
	"net/http"

	"github.com/DMohanty99/eventApp/models"
	"github.com/DMohanty99/eventApp/utils"
	"github.com/gin-gonic/gin"
)

func HandleSignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request data", "actualerr": err})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "object created successfully"})

}

func HandleSignIn(c *gin.Context) {
	var u models.User
	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = u.ValidateCred()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(u.Email, u.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate JWT token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully", "token": token})
}
