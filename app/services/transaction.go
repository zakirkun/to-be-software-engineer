package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"teukufuad/e-commerce/app/domains/contracts"
	"teukufuad/e-commerce/app/domains/models"
	"teukufuad/e-commerce/app/domains/types"
	"teukufuad/e-commerce/app/repository"
	"teukufuad/e-commerce/pkg/rabbitmq"
	"teukufuad/e-commerce/utils"
)

type TransactionService struct {
}

func (t TransactionService) Order(request types.RequestOrder) (*types.ResponseOrder, error) {

	custRepo := repository.NewCustomerRepository()
	getCust, err := custRepo.GetByEmail("coba2@admin.com")
	if err != nil {
		return nil, err
	}

	if getCust.Id == 0 {
		return nil, errors.New("customer not found")
	}

	data := models.Transaction{
		ProductId:  request.ProductId,
		CustomerId: getCust.Id,
		Amount:     request.Amount,
		Qty:        request.Qty,
	}

	repo := repository.NewTransactionRepository()
	//auth := jwt.NewJWTImpl(config.GetString("jwt.signature_key"), config.GetInt("jwt.day_expired"))
	//calims := auth.ParseToken()
	response, err := repo.Order(data)

	if err != nil {
		return nil, err
	}

	var sendEmail = func() {
		payload := make(map[string]interface{})
		payload["to"] = "fuad.pcd@gmail.com"
		//payload["to"] = getCust.Email
		payload["body"] = fmt.Sprintf("You order %d, was created", response.Id)

		if err := rabbitmq.RMQ.Publish("email_services", payload); err != nil {
			log.Printf("MESSAGES_BROKER_ERROR: %v", err)
		}
	}

	go sendEmail()

	return &types.ResponseOrder{
		Transaction: response,
	}, nil
}

// HandleSentEmail implements contracts.OrderServices.
func (t TransactionService) HandleSentEmail(data []byte) error {

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

func NewTransactionService() contracts.TransactionServices {
	return TransactionService{}
}
