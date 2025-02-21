package main

import (
    "github.com/gin-gonic/gin"

    "awsses/api"
)



func main() {
    router := gin.Default()
    router.POST("api/v1/email/send", api.SendEmail)

    router.Run("localhost:8000")
}