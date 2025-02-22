package main

import (
	"log"

	"awsses/api"
	"awsses/database"
	"awsses/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	err := database.DB.AutoMigrate(&models.EmailAccount{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	router := gin.Default()
	router.POST("api/v1/email/send", api.SendEmail)

	router.Run("localhost:8000")
}
