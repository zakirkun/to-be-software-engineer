package main

import (
	"flag"
	"log"
	"os"
	"time"

	"imzakir.dev/e-commerce/bootstrap"
	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/pkg/config"
	"imzakir.dev/e-commerce/pkg/database"
	"imzakir.dev/e-commerce/pkg/logstash"
	"imzakir.dev/e-commerce/pkg/rabbitmq"
	"imzakir.dev/e-commerce/pkg/server"
	"imzakir.dev/e-commerce/router"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()

}

func main() {
	setConfig()

	infra := bootstrap.NewInfrastructure(SetLogstash(), SetDatabase(), SetCache(), SetRabbitMq(), SetWebServer())
	infra.Logstash()
	infra.Database()
	infra.Cache()
	infra.MessagesBroker()
	infra.WebServer()
}

func setConfig() {
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}
}

func SetCache() cache.Cache {
	return cache.Cache{
		Addr:     config.GetString("cache.cache_addr"),
		Password: config.GetString("cache.cache_password"),
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

func SetWebServer() server.ServerContext {
	return server.ServerContext{
		AppName:      config.GetString("server.app_name"),
		Host:         ":" + config.GetString("server.port"),
		Handler:      router.InitRouters(),
		ReadTimeout:  time.Duration(config.GetInt("server.http_timeout")),
		WriteTimeout: time.Duration(config.GetInt("server.http_timeout")),
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
