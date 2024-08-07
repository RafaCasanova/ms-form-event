package mail

import (
	"challenger/adapter/input/model/response"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"text/template"
)

func SendEmailToUser(contact response.ContactResponse) {
	from := os.Getenv("MAIL_AUTH_USER")
	password := os.Getenv("MAIL_AUTH_PASS")
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	t, err := template.ParseFiles("./template/user_email.html")
	if err != nil {
		log.Println("Fail to load template:", err)
		return
	}
	var body strings.Builder
	err = t.Execute(&body, contact)
	if err != nil {
		log.Println(err)
		return
	}
	to := []string{contact.Email}
	subject := os.Getenv("TEXT_MAIL_TITLE_USER")
	message := []byte("From: " + os.Getenv("MAIL_AUTH_USER") + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body.String())

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email enviado com sucesso!")
}
