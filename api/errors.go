package api

import (
    "github.com/gin-gonic/gin"
)


type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func isSendingPaused() bool {
	return false
}

func isConfigurationSetDoesNotExist() bool {
	return true
}

func isConfigurationSetSendingPaused() bool {
	return false
}

func isMessageRejected() bool {
	return false
}


func ErrorsCheck() (bool, string, string) {
	if isSendingPaused() {
		return "AccountSendingPaused", "Email sending is disabled for your account. Please contact support."
	}
	if !isConfigurationSetDoesNotExist() {
		return "ConfigurationSetDoesNotExist", "The specified configuration set does not exist."
	}

	if !isConfigurationSetSendingPaused() {
		return "ConfigurationSetDoesNotExist", "The specified configuration set does not exist."
	}

	if isMessageRejected() {
		return "MessageRejected", "The email message was rejected due to policy restrictions."
	}

	return "", ""
}