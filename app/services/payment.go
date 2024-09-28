package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"imzakir.dev/e-commerce/app/domains/contracts"
	"imzakir.dev/e-commerce/app/domains/types"
	"imzakir.dev/e-commerce/app/repository"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
)

type paymentServices struct{}

// HandleLogging implements contracts.PaymentServices.
func (p paymentServices) HandleLogging(data []byte) error {

	// Init logstash
	logging := logstash.LogPayload{}

	var body types.PaymentParameter

	err := json.Unmarshal(data, &body)
	if err != nil {
		return err
	}

	if err := logging.SetData(body).SetIndex("payment_services").SetAppName("payment_delivery_services").WriteCaller("INFO", "Payment delivery services recive payload"); err != nil {
		return err
	}

	return nil
}

var s snap.Client

func initPayment() {
	mode := config.GetString("payment.mode")
	if mode == "sandbox" {
		s.Env = midtrans.Sandbox
	} else {
		s.Env = midtrans.Production
	}

	s.New(config.GetString("payment.sb_server_key"), s.Env)
}

// HandlePaymentCallback implements contracts.PaymentServices.
func (p paymentServices) HandlePayment(data []byte) error {
	var body types.PaymentParameter

	err := json.Unmarshal(data, &body)
	if err != nil {
		return err
	}

	// get Cust
	where := make(map[string]interface{})
	where["username"] = body.Username

	custRepo := repository.NewCustomerRepository()
	getCust, err := custRepo.GetWhere(where)
	if err != nil {
		return err
	}

	if getCust.ID == 0 {
		return errors.New("customer not found")
	}

	// Get Product
	newWhere := make(map[string]interface{})
	newWhere["id"] = body.ProductId
	prodRepo := repository.NewProductRepository()
	getProd, err := prodRepo.FindBy(newWhere)
	if err != nil {
		return err
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ORDER-" + strconv.Itoa(body.TrxId),
			GrossAmt: int64(body.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: getCust.FullName,
			Email: getCust.Email,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		CreditCard: &snap.CreditCardDetails{
			Secure: false,
		},
	}

	var items []midtrans.ItemDetails

	if len(*getProd) > 0 {
		for _, product := range *getProd {
			id := strconv.Itoa(int(product.ID))
			items = append(items, midtrans.ItemDetails{
				ID:    id,
				Name:  product.ProductName,
				Qty:   1,
				Price: int64(body.Amount),
			})
		}
	}

	snapReq.Items = &items

	resp, midErr := s.CreateTransaction(snapReq)
	if midErr != nil {
		return err
	}

	paymentLink := resp.RedirectURL

	sentEmailParam := make(map[string]interface{})
	sentEmailParam["To"] = getCust.Email
	sentEmailParam["Subject"] = fmt.Sprintf("INVOICE PAYMENT: #%d", body.TrxId)
	sentEmailParam["Body"] = fmt.Sprintf("Silahkan melakukan pembayaran melalui tautan berikut: %v", paymentLink)

	log.Info(sentEmailParam)

	if err := rabbitmq.RMQ.Publish("email_services", sentEmailParam); err != nil {
		log.Printf("EMAIL_SERVICES_MESSAGES_BROKER_ERROR: %v", err)
	}

	log.Printf("Payment Generated! %v", getCust.Username)

	return nil
}

func NewPaymentServices() contracts.PaymentServices {
	initPayment()
	return paymentServices{}
}
