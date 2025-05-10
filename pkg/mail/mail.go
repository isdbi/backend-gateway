package mail

import (
	"net/smtp"
	"os"
)

func SendEmail(to []string, msg []byte) error {
	username := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpAuth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", smtpAuth, username, to, msg)

	if err != nil {
		return err
	}
	return nil
}
