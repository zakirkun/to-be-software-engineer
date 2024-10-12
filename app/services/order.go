package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/models"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
	"imzakir.dev/e-commerce/utils"
)

type orderServices struct{}

// HandleLogging implements contracts.OrderServices.
func (o orderServices) HandleLogging(data []byte) error {

	// Init logstash
	logging := logstash.LogPayload{}

	// define struct body
	body := make(map[string]interface{})

	err := json.Unmarshal(data, &body)
	if err != nil {
		return err
	}

	if err := logging.SetData(body).SetIndex("email_services").SetAppName("email_delivery_services").WriteCaller("INFO", "Email delivery services recive payload"); err != nil {
		return err
	}

	return nil
}

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
	// init repo
	repo := repository.NewOrderRepository()

	// check cache first
	_id := strconv.Itoa(int(id))
	getKey, err := cache.CACHE.Get(context.Background(), fmt.Sprintf("order:%v", _id)).Result()
	if err == redis.Nil {
		getTrx, err := repo.GetByID(id)
		if err != nil {
			return nil, err
		}

		if getTrx.ID == 0 {
			return nil, errors.New("record not found")
		}

		toJson := utils.StructToJson(&getTrx)
		cache.CACHE.Set(context.Background(), fmt.Sprintf("order:%v", _id), toJson, time.Duration(time.Minute*30))
		return &types.ResponseGetTransaction{
			Transaction: getTrx,
		}, nil
	}

	var parse models.Transaction
	if ok := utils.JsonToSruct([]byte(getKey), &parse); !ok {
		return nil, errors.ErrUnsupported
	}

	return &types.ResponseGetTransaction{
		Transaction: &parse,
	}, nil

}

func NewOrderServices() contracts.OrderServices {
	return orderServices{}
}
