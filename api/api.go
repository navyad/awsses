package api

import (
	"log"
	"net/http"
	"time"

	"awsses/database"
	"awsses/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var warmingPeriod = 2 * 7 * 24 * time.Hour

func getEmailAccount(db *gorm.DB, accountID string) (*models.EmailAccount, error) {
	var account models.EmailAccount
	if err := db.First(&account, "id = ?", accountID).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func updateEmailAccount(db *gorm.DB, account *models.EmailAccount, isBounced bool) error {
	account.DailySendCount++
	account.TotalEmails++

	if isBounced {
		account.Bounce++
	}
	
	if err := db.Save(&account).Error; err != nil {
		return err
	}
	
	return nil
}

func getDailyLimit(account *models.EmailAccount) int {
	accountAge := time.Since(account.CreatedAt)
	log.Println("getDailyLimit", accountAge, warmingPeriod)

	if accountAge < warmingPeriod {
		limit := int(float64(account.DailySendLimit) * (float64(accountAge) / float64(warmingPeriod)))
		log.Println("warmingPeriod", limit)
		if limit < 1 {
			limit = 1
		}
		return limit

	}
	return account.DailySendLimit
}

func CheckwarmingPeriod(account *models.EmailAccount) bool {
	currentLimit := getDailyLimit(account)
	return account.DailySendCount >= currentLimit
}


// isValidEmail validates a single email address using net/mail.

func SendEmail(c *gin.Context) {
	var emailReq EmailRequest

	errorCode, errorMessage := ErrorsCheck()
	if errorCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    errorCode,
				"message": errorMessage,
			},
		})
		return
	}

	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok, inValidEmail := ValidateEmail(&emailReq) 
	if !ok{
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Email: " + inValidEmail})
		return
	}

	ok, errorMessage = ValidateRecipientsLength(&emailReq) 
	if !ok{
		c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}


	account, err := getEmailAccount(database.DB, emailReq.Source)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if CheckwarmingPeriod(account) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": gin.H{
				"code":    "AccountWarmingUp",
				"message": "Your account is still warming up. Please reduce your email sending volume.",
			},
		})
		return
	}

	// mocking the email bounce 
	isBounced := true

	updateEmailAccount(database.DB, account, isBounced)

	c.JSON(http.StatusOK, gin.H{
		"MessageId": RandomMessageID(),
	})
}


func GetEmailStats(c *gin.Context) {
	accountID := c.Query("accountId")
	if accountID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "accountId query parameter is required"})
		return
	}

	var account models.EmailAccount
	if err := database.DB.First(&account, "id = ?", accountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accountId":         account.ID,
		"dailySendLimit":    account.DailySendLimit,
		"dailySendCount":    account.DailySendCount,
		"totalEmailsSent":       account.TotalEmails,
	})
}
