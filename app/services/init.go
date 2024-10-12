package services

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"imzakir.dev/e-commerce/pkg/config"
)

var s snap.Client
var c coreapi.Client

func init() {
	mode := config.GetString("payment.sb_server_mode")
	serverKey := config.GetString("payment.sb_server_key")
	midtrans.ServerKey = serverKey

	if mode == "sandbox" {
		s.Env = midtrans.Sandbox
		midtrans.Environment = midtrans.Sandbox
	} else {
		s.Env = midtrans.Production
		midtrans.Environment = midtrans.Production
	}

	s.New(serverKey, s.Env)
}
