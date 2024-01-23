package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/joncalhoun/lenslocked/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

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
	err = es.ForgotPassword("sun.motiani@gmail.com", "https://yayipi.in/")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
