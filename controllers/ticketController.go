package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/initializers"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/models"
)

func TicketCreate(c *gin.Context) {
	// Get data from request
	var body struct {
		UserId uint
		Status bool
	}

	c.Bind(&body)

	// Create ticket
	ticket := models.Ticket{
		UserID: body.UserId,
		Status: body.Status,
	}

	// Check if ticket was created
	result := initializers.DB.Create(&ticket)

	// Return response
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Ticket not created",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Ticket created",
			"ticket":  ticket,
		})
	}
}

func TicketIndex(c *gin.Context) {
	// Get all tickets
	var tickets []models.Ticket

	initializers.DB.Preload("User").Find(&tickets)

	// Return response
	c.JSON(200, gin.H{
		"tickets": tickets,
	})
}

func TicketShow(c *gin.Context) {
	// Get ticket id from request
	id := c.Param("id")

	// Get ticket
	var ticket models.Ticket

	initializers.DB.Preload("User").First(&ticket, id)

	// Return response
	c.JSON(200, gin.H{
		"ticket": ticket,
	})
}
