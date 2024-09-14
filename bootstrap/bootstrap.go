package bootstrap

import (
	"context"
	"log"
	"os"

	"imzakir.dev/e-commerce/pkg/cache"
	"imzakir.dev/e-commerce/pkg/database"
	"imzakir.dev/e-commerce/pkg/server"
)

type Infrastructure interface {
	Database()
	Cache()
	WebServer()
}

type infrastructureContext struct {
	database database.DBModel
	cache    cache.Cache
	server   server.ServerContext
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

func NewInfrastructure(database database.DBModel,
	cache cache.Cache,
	server server.ServerContext,
) Infrastructure {
	return infrastructureContext{
		database: database,
		cache:    cache,
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
