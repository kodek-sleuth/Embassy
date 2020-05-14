package helpers

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(toEmail string, code string) {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", os.Getenv("SEND_MAIL_EMAIL"), os.Getenv("SEND_MAIL_PASSWORD"),
		"smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{toEmail}
	msg := []byte(fmt.Sprintf("To: %s\r\n", toEmail) +
		"Subject: UHCA CODE?\r\n" +
		"\r\n" +
		fmt.Sprintf("Please safe guard your code %s\r\n", code))
	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("SEND_MAIL_EMAIL"), to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
