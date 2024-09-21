package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"golang.org/x/crypto/bcrypt"
)

func StructToJson(src any) string {
	jsonBytes, err := json.Marshal(src)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	return string(jsonBytes)
}

func JsonToSruct(src []byte, to any) bool {
	err := json.Unmarshal(src, &to)
	if err != nil {
		log.Println("Error:", err)
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SendEmail(to, htmlBody string) error {

	// sending email
	smtp := mail.NewSMTPClient()

	// setup smtp
	smtp.Host = "smtp.mailersend.net"
	smtp.Port = 587
	smtp.Username = "MS_EBHkul@trial-pq3enl6m237g2vwr.mlsender.net"
	smtp.Password = "jz5XmuB79JnXu7KN"
	smtp.Encryption = mail.EncryptionSTARTTLS

	// Variable to keep alive connection
	smtp.KeepAlive = true

	// Timeout for connect to SMTP Server
	smtp.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	smtp.SendTimeout = 10 * time.Second

	// test connection
	smtpClient, err := smtp.Connect()
	if err != nil {
		log.Fatal("Failed Connect SMTP")
		log.Fatal(err)
	}

	// New email simple html with inline and CC
	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("From %v <%v>", "TOKO ONLINE", "MS_EBHkul@trial-pq3enl6m237g2vwr.mlsender.net")).
		AddTo(to).
		SetSubject("ORDER SUCCESS")

	email.SetBody(mail.TextHTML, htmlBody)

	// always check error after send
	if email.Error != nil {
		log.Fatal("Failed Sending Email")
		log.Fatal(email.Error)
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email Sent to " + to)
	}

	fmt.Println("-> Processed!")

	return nil
}
