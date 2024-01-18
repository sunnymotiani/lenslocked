package main

import (
	"fmt"

	"github.com/joncalhoun/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "989b7e62f05bc6"
	password = "50c07f343af50a"
)

func main() {
	/*
		email := models.Email{
			From:      "test@lenslocked.com",
			To:        "sun.motiani@gmail.com",
			Subject:   "This is a test email",
			Plaintext: "This is the body of the email",
			HTML:      `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`,
		}
	*/
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.ForgotPassword("sun.motiani@gmail.com", "https://yayipi.in/")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
