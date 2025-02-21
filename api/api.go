package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
)


func SendEmail(c *gin.Context) {
	var emailReq EmailRequest

	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"Message": "email sent successfully",
	})
}