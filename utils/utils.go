package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
	"golang.org/x/crypto/bcrypt"
	"imzakir.dev/e-commerce/pkg/config"
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

func SendEmail(to, htmlBody, subject string) error {

	// sending email
	smtp := mail.NewSMTPClient()

	// setup smtp
	smtp.Host = config.GetString("smtp.host")
	smtp.Port = config.GetInt("smtp.port")
	smtp.Username = config.GetString("smtp.username")
	smtp.Password = config.GetString("smtp.password")
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
	email.SetFrom(fmt.Sprintf("From %v <%v>", config.GetString("server.app_name"), config.GetString("smtp.username"))).
		AddTo(to).
		SetSubject(subject)

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

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
