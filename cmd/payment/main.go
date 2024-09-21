package main

import (
	"flag"
	"log"
	"os"

	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/database"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	bundle := NewBudle(SetDatabase(), SetRabbitMq())
	bundle.Database()
	bundle.EventPayment()
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

func SetDatabase() database.DBModel {
	return database.DBModel{
		ServerMode:   config.GetString("server.mode"),
		Driver:       config.GetString("database.db_driver"),
		Host:         config.GetString("database.db_host"),
		Port:         config.GetString("database.db_port"),
		Name:         config.GetString("database.db_name"),
		Username:     config.GetString("database.db_username"),
		Password:     config.GetString("database.db_password"),
		MaxIdleConn:  config.GetInt("pool.conn_idle"),
		MaxOpenConn:  config.GetInt("pool.conn_max"),
		ConnLifeTime: config.GetInt("pool.conn_lifetime"),
	}
}

type iBundleListener struct {
	database database.DBModel
	rabbitmq rabbitmq.RabbitMQ
}

// Database implements IBudleInterface.
func (i iBundleListener) Database() {
	_, err := i.database.OpenDB()
	if err != nil {
		os.Exit(1)
	}

	database.DB = &i.database
}

// RabbitMq implements IBudleInterface.
func (i iBundleListener) EventPayment() {
	_, err := i.rabbitmq.Open()
	if err != nil {
		log.Fatalf("Failed Connect RabbitMq: %v", err)
		os.Exit(1)
	}

	rabbitmq.RMQ = &i.rabbitmq

	// Enable Listener
	i.rabbitmq.Listener("payment_services", services.NewPaymentServices().HandlePayment)

}

type IBudleInterface interface {
	Database()
	EventPayment()
}

func NewBudle(database database.DBModel, rabbitmq rabbitmq.RabbitMQ) IBudleInterface {
	return iBundleListener{database: database, rabbitmq: rabbitmq}
}
