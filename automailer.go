// Written by Martijn
// Automatic mailing system

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/go-mail/mail"
)

// Configuration struct for reading from config file
type Configuration struct {
    SMTPServer   string `json:"smtp_server"`
    SMTPPort     int    `json:"smtp_port"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    Recipient    string `json:"recipient"`
    Subject      string `json:"subject"`
    Body         string `json:"body"`
    Attachment   string `json:"attachment"`
}

// Read config from file
func readConfig() Configuration {
    configFile, err := os.Open("config.json")
    if err != nil {
        log.Fatalf("Failed to open config file: %v", err)
    }
    defer configFile.Close()

    configBytes, err := ioutil.ReadAll(configFile)
    if err != nil {
        log.Fatalf("Failed to read config file: %v", err)
    }

    var config Configuration
    err = json.Unmarshal(configBytes, &config)
    if err != nil {
        log.Fatalf("Failed to parse config file: %v", err)
    }

    return config
}

func main() {
    // Read configuration from file
    config := readConfig()

    // Read attachment file
    attachmentFile, err := os.Open(config.Attachment)
    if err != nil {
        log.Fatalf("Failed to open attachment file: %v", err)
    }
    defer attachmentFile.Close()

    attachmentBytes, err := ioutil.ReadAll(attachmentFile)
    if err != nil {
        log.Fatalf("Failed to read attachment file: %v", err)
    }

    attachment := mail.Attachment{
        Name:        filepath.Base(config.Attachment),
        Content:     attachmentBytes,
        ContentType: "application/octet-stream",
    }

    // Create email message
    m := mail.NewMessage()
    m.SetHeader("From", config.Username)
    m.SetHeader("To", config.Recipient)
    m.SetHeader("Subject", config.Subject)
    m.SetBody("text/plain", config.Body)
    m.Attach(attachment)

    // Connect to SMTP server and send email
    auth := smtp.PlainAuth("", config.Username, config.Password, config.SMTPServer)
    err = mail.SendMail(config.SMTPServer+":"+string(config.SMTPPort), auth, m)
    if err != nil {
        log.Fatalf("Failed to send email: %v", err)
    }

    log.Println("Email sent successfully")
}
