package util

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTPEmail(to string, otp string) error {

	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf(
		"Subject: DoFocus OTP Verification\r\n\r\nYour OTP is: %s",
		otp,
	))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		message,
	)

	return err
}