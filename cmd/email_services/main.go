package main

import (
	"flag"
	"log"
	"os"

	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	bundle := NewBudle(SetLogstash(), SetRabbitMq())
	bundle.Logstash()
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

func SetLogstash() logstash.LogstashModel {
	return logstash.LogstashModel{
		Network: config.GetString("logstash.network"),
		Addr:    config.GetString("logstash.addr"),
	}
}

type iBundleListener struct {
	logstash logstash.LogstashModel
	rabbitmq rabbitmq.RabbitMQ
}

// Logstash implements IBudleInterface.
func (i iBundleListener) Logstash() {

	_, err := i.logstash.Open()
	if err != nil {
		log.Fatalf("Failed Connect Logstash: %v", err)
		os.Exit(1)
	}

	logstash.LOGSTASH = &i.logstash
}

// EventEmail implements IBudleInterface.
func (i iBundleListener) EventEmail() {
	_, err := i.rabbitmq.Open()
	if err != nil {
		log.Fatalf("Failed Connect RabbitMq: %v", err)
		os.Exit(1)
	}

	// Enable Listener
	i.rabbitmq.Listener("email_services", services.NewOrderServices().HandleSentEmail, services.NewOrderServices().HandleLogging)
}

type IBudleInterface interface {
	Logstash()
	EventEmail()
}

func NewBudle(logstash logstash.LogstashModel,
	rabbitmq rabbitmq.RabbitMQ) IBudleInterface {
	return iBundleListener{
		logstash: logstash,
		rabbitmq: rabbitmq,
	}
}
