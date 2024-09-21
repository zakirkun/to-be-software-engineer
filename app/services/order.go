package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
	"imzakir.dev/e-commerce/utils"
)

type orderServices struct{}

// HandleSentEmail implements contracts.OrderServices.
func (o orderServices) HandleSentEmail(data []byte) error {

	// define struct body
	body := make(map[string]interface{})

	err := json.Unmarshal(data, &body)
	if err != nil {
		return err
	}

	to, _ := body["to"].(string)
	htmlBody, _ := body["body"].(string)

	if err := utils.SendEmail(to, htmlBody); err != nil {
		return err
	}

	return nil
}

// CreateTransaction implements contracts.OrderServices.
func (o orderServices) CreateTransaction(request types.RequestCreateTransaction) (*types.ResponseGetTransaction, error) {
	// get user
	where := make(map[string]interface{})
	where["username"] = request.Username
	custRepo := repository.NewCustomerRepository()
	getCust, err := custRepo.GetWhere(where)
	if err != nil {
		return nil, err
	}

	if getCust.Id == 0 {
		return nil, errors.New("customer not found")
	}

	data := models.Transaction{
		ProductID:  uint(request.ProductID),
		CustomerID: uint(getCust.Id),
		Amount:     request.Amount,
		Qty:        request.Qty,
	}

	repo := repository.NewOrderRepository()
	if err := repo.Create(data); err != nil {
		return nil, err
	}

	var sendEmail = func() {
		payload := make(map[string]interface{})
		payload["to"] = getCust.Email
		payload["body"] = fmt.Sprintf("You order %d, was created", data.ID)

		if err := rabbitmq.RMQ.Publish("email_services", payload); err != nil {
			log.Printf("MESSAGES_BROKER_ERROR: %v", err)
		}
	}

	go sendEmail()

	return &types.ResponseGetTransaction{
		Transaction: &data,
	}, nil
}

// GetTransaction implements contracts.OrderServices.
func (o orderServices) GetTransaction(id uint) (*types.ResponseGetTransaction, error) {
	getTrx, err := repository.NewOrderRepository().GetByID(id)
	if err != nil {
		return nil, err
	}

	return &types.ResponseGetTransaction{
		Transaction: getTrx,
	}, nil
}

func NewOrderServices() contracts.OrderServices {
	return orderServices{}
}
