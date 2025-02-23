package api

import (
	"log"
	"net/mail"
	"strings"
)

func isValidEmail(email string) bool {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	if dot := strings.LastIndex(domain, "."); dot < 0 || dot == len(domain)-1 {
		return false
	}
	return true
}


func ValidateEmail(req *EmailRequest) (bool, string) {
	if !isValidEmail(req.Source) {
		return false, req.Source
	}

	if len(req.Destination.ToAddresses) == 0 {
		return false, ""
	}
	for _, addr := range req.Destination.ToAddresses {
		if !isValidEmail(addr) {
			return false, addr
		} else {
			log.Println(addr, isValidEmail(addr))
		}
	}

	for _, addr := range req.Destination.CcAddresses {
		if addr != "" && !isValidEmail(addr) {
			return false, addr
		}
	}

	for _, addr := range req.Destination.BccAddresses {
		if addr != "" && !isValidEmail(addr) {
			return false, addr
		}
	}
	return true, ""
}
