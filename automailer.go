// Written by Martijn
// Automatic mailing system

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Open the config file
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	// Parse the JSON file
	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the authentication information.
	auth := smtp.PlainAuth("", config.Email, config.Password, "smtp.office365.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err = smtp.SendMail("smtp.office365.com:587", auth, config.Email, []string{"500334@student.fontys.nl"}, []byte("Subject: Test email\n\nThis is a test email."))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully!")
}
