package main

import (
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/initializers"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Ticket{})
}
