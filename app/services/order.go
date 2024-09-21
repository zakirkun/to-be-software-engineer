package services

import (
	"encoding/json"
	"errors"

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

	to, _ := body["To"].(string)
	htmlBody, _ := body["Body"].(string)
	subject, _ := body["Subject"].(string)

	if err := utils.SendEmail(to, htmlBody, subject); err != nil {
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

	if getCust.ID == 0 {
		return nil, errors.New("customer not found")
	}

	data := models.Transaction{
		ProductID:  uint(request.ProductID),
		CustomerID: getCust.ID,
		Amount:     request.Amount,
		Qty:        request.Qty,
	}

	repo := repository.NewOrderRepository()
	if err := repo.Create(&data); err != nil {
		return nil, err
	}

	var sendPaymentServices = func() {
		req := types.PaymentParameter{
			Username:  getCust.Username,
			ProductId: request.ProductID,
			TrxId:     int(data.ID),
			Amount:    int(request.Amount),
		}

		if err := rabbitmq.RMQ.Publish("payment_services", req); err != nil {
			log.Printf("PAYMENT_SERVICES_MESSAGES_BROKER_ERROR: %v", err)
		}
	}

	go sendPaymentServices()

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
