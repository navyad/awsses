package api

import (
	"crypto/rand"
	"fmt"
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

	destination := req.Destination

	if len(destination.ToAddresses) == 0 && len(destination.CcAddresses) == 0 && len(destination.BccAddresses) == 0 {
		return false, "At least one recipient must be provided in To, CC, or BCC."
	}

	if !isValidEmail(req.Source) {
		return false, req.Source
	}

	if len(destination.ToAddresses) == 0 {
		return false, "ToAddresses required"
	}
	for _, addr := range destination.ToAddresses {
		if !isValidEmail(addr) {
			return false, addr
		}
	}

	for _, addr := range destination.CcAddresses {
		if addr != "" && !isValidEmail(addr) {
			return false, addr
		}
	}

	for _, addr := range destination.BccAddresses {
		if addr != "" && !isValidEmail(addr) {
			return false, addr
		}
	}
	return true, ""
}

func ValidateRecipientsLength(req *EmailRequest) (bool, string) {

	destination := req.Destination

	if len(destination.ToAddresses) > 50 {
		return false, "ToAddresses contains more than 50 emails"
	}
	if len(destination.CcAddresses) > 50 {
		return false, "CcAddresses contains more than 50 emails"
	}
	if len(destination.BccAddresses) > 50 {
		return false, "BccAddresses contains more than 50 emails"
	}
	return true, ""
}

func randomHexDigits(n int) string {
	bytesNeeded := (n + 1) / 2
	b := make([]byte, bytesNeeded)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)
	if len(s) > n {
		s = s[:n]
	}
	return s
}

func RandomMessageID() string {
	part1 := randomHexDigits(4)
	part2 := randomHexDigits(4)
	part3 := randomHexDigits(4)
	part4 := randomHexDigits(12)
	part5 := randomHexDigits(6)
	return fmt.Sprintf("%s-%s-%s-%s-%s", part1, part2, part3, part4, part5)
}
