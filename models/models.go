package models

import "time"

type EmailAccount struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	DailySendCount int       `json:"daily_send_count"`
	DailySendLimit int       `json:"daily_send_limit"`
	TotalEmails    int       `json:"total_emails"`
}


