package services

import (
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/jwt"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
)

type customerServices struct{}

func (c customerServices) Login(customer types.RequestLogin) (*types.ResponseLogin, error) {
	repo := repository.NewCustomerRepository()
	dataCustomer, err := repo.GetByUsername(customer.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dataCustomer.Password), []byte(customer.Password)); err != nil {
		return nil, err
	}
	signature := config.GetString("jwt.signature_key")
	expired, _ := strconv.Atoi(config.GetString("jwt.day_expired"))

	jwtClaims := jwt.NewJWTImpl(signature, expired)
	token, err := jwtClaims.GenerateToken(map[string]interface{}{
		"username": dataCustomer.Username,
		"type":     "Bearer",
	})
	if err != nil {
		return nil, err
	}

	return &types.ResponseLogin{
		JwtToken: token,
	}, nil
}

func (c customerServices) AddCustomer(request types.RequestCreateCustomer) (*types.ResponseCreateCustomer, error) {
	repo := repository.NewCustomerRepository()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	request.Password = string(hashedPassword)
	data, err := repo.Insert(models.Customer{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		FullName: request.FullName,
	})
	if err != nil {
		return nil, err
	}

	var sendWelcomeLetter = func() {
		sentEmailParam := make(map[string]interface{})
		sentEmailParam["To"] = data.Email
		sentEmailParam["Subject"] = fmt.Sprintf("Welcome To %v", config.GetString("server.app_name"))
		sentEmailParam["Body"] = "Thanks for registration"
		log.Info(sentEmailParam)
		if err := rabbitmq.RMQ.Publish("email_services", sentEmailParam); err != nil {
			log.Printf("EMAIL_SERVICES_MESSAGES_BROKER_ERROR: %v", err)
		}
		log.Info("Welcome letter success send")
	}
	go sendWelcomeLetter()

	return &types.ResponseCreateCustomer{
		Customer: data,
	}, nil
}

func NewCustomerServices() contracts.CustomerServices {
	return customerServices{}
}
