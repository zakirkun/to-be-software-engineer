package services

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

var s snap.Client
var c coreapi.Client

func init() {
	// mode := config.GetString("payment.sb_server_mode")
	mode := "sandbox"
	// serverKey := config.GetString("payment.sb_server_key")
	serverKey := "SB-Mid-server-FxnHTffLCcsFPQjMcJTcJrTb"
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
