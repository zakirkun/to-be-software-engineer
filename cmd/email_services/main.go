package main

import (
	"flag"
	"log"
	"os"

	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	bundle := NewBudle(SetRabbitMq())
	bundle.EventEmail()
}

func setConfig() {
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}
}

func SetRabbitMq() rabbitmq.RabbitMQ {
	return rabbitmq.RabbitMQ{
		Address: config.GetString("message_broker.rabbimq_url"),
	}
}

type iBundleListener struct {
	rabbitmq rabbitmq.RabbitMQ
}

// EventEmail implements IBudleInterface.
func (i iBundleListener) EventEmail() {
	_, err := i.rabbitmq.Open()
	if err != nil {
		log.Fatalf("Failed Connect RabbitMq: %v", err)
		os.Exit(1)
	}

	// Enable Listener
	i.rabbitmq.Listener("email_services", services.NewOrderServices().HandleSentEmail)
}

type IBudleInterface interface {
	EventEmail()
}

func NewBudle(rabbitmq rabbitmq.RabbitMQ) IBudleInterface {
	return iBundleListener{rabbitmq: rabbitmq}
}
