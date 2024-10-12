package bootstrap

import (
	"context"
	"log"
	"os"

	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/pkg/database"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
	"imzakir.dev/e-commerce/pkg/server"
)

type Infrastructure interface {
	Logstash()
	Database()
	Cache()
	MessagesBroker()
	WebServer()
}

type infrastructureContext struct {
	logstash logstash.LogstashModel
	database database.DBModel
	cache    cache.Cache
	rabbitmq rabbitmq.RabbitMQ
	server   server.ServerContext
}

// Logstash implements Infrastructure.
func (i infrastructureContext) Logstash() {
	_, err := i.logstash.Open()
	if err != nil {
		log.Fatalf("Failed Connect Logstash: %v", err)
		os.Exit(1)
	}

	logstash.LOGSTASH = &i.logstash
}

// MessagesBroker implements Infrastructure.
func (i infrastructureContext) MessagesBroker() {
	_, err := i.rabbitmq.Open()
	if err != nil {
		log.Fatalf("Failed Connect RabbitMq: %v", err)
		os.Exit(1)
	}

	rabbitmq.RMQ = &i.rabbitmq
}

// Cache implements Infrastructure.
func (i infrastructureContext) Cache() {
	conn := i.cache.Open()

	// Check Connection
	if err := conn.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	// Assign connection
	cache.CACHE = conn
}

func NewInfrastructure(
	logstash logstash.LogstashModel,
	database database.DBModel,
	cache cache.Cache,
	rabbitmq rabbitmq.RabbitMQ,
	server server.ServerContext,
) Infrastructure {
	return infrastructureContext{
		logstash: logstash,
		database: database,
		cache:    cache,
		rabbitmq: rabbitmq,
		server:   server,
	}
}

func (i infrastructureContext) Database() {
	_, err := i.database.OpenDB()
	if err != nil {
		os.Exit(1)
	}

	database.DB = &i.database

}

func (i infrastructureContext) WebServer() {
	i.server.Run()
}
