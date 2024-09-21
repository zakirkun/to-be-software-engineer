package main

import (
	"flag"
	"log"
	"os"

	"teukufuad/e-commerce/app/services"
	"teukufuad/e-commerce/pkg/config"
	"teukufuad/e-commerce/pkg/rabbitmq"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	i := SetRabbitMq()
	_, err := i.Open()
	if err != nil {
		log.Fatalf("Failed Connect RabbitMq: %v", err)
		os.Exit(1)
	}

	// Enable Listener
	i.Listener("email_services", func(payload []byte) error {
		log.Printf("Recive Messages: %v", string(payload))

		return nil
	}, services.NewTransactionService().HandleSentEmail)
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
