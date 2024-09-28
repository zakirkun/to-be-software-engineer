package main

import (
	"flag"
	"log"
	"os"

	"imzakir.dev/e-commerce/app/services"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/database"
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

	bundle := NewBudle(SetLogstash(), SetDatabase(), SetRabbitMq())
	bundle.Logstash()
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

func SetLogstash() logstash.LogstashModel {
	return logstash.LogstashModel{
		Network: config.GetString("logstash.network"),
		Addr:    config.GetString("logstash.addr"),
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
	logstash logstash.LogstashModel
	database database.DBModel
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
	i.rabbitmq.Listener("payment_services", services.NewPaymentServices().HandlePayment, services.NewPaymentServices().HandleLogging)

}

type IBudleInterface interface {
	Logstash()
	Database()
	EventPayment()
}

func NewBudle(logstash logstash.LogstashModel, database database.DBModel, rabbitmq rabbitmq.RabbitMQ) IBudleInterface {
	return iBundleListener{
		logstash: logstash,
		database: database,
		rabbitmq: rabbitmq,
	}
}
