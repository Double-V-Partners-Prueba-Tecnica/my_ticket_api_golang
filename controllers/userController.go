package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/initializers"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/models"
)

func UserCreate(c *gin.Context) {
	// Get data from request
	var body struct {
		Username string
	}

	c.Bind(&body)

	// Create user
	user := models.User{
		Username: body.Username,
	}

	// Check if user was created
	result := initializers.DB.Create(&user)

	// Return response
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "User not created",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "User created",
			"user":    user,
		})
	}
}

func UserIndex(c *gin.Context) {
	// Get all users
	var users []models.User

	initializers.DB.Find(&users)

	// Return response
	c.JSON(200, gin.H{
		"users": users,
	})
}

func UserShow(c *gin.Context) {
	// Get user id from request
	id := c.Param("id")

	// Get user
	var user models.User

	initializers.DB.First(&user, id)

	// Return response
	c.JSON(200, gin.H{
		"user": user,
	})
}
