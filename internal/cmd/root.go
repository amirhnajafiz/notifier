package cmd

import (
	"cmd/internal/broker"
	"cmd/internal/client"
	"cmd/internal/cmd/server"
	"cmd/internal/config"
)

func Execute() {
	cfg := config.New()

	opts := broker.MQTT{
		Cfg: cfg.Broker,
	}.Register()

	clt := client.Client{}.Register(opts)

	if token := clt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	server.Register()
}
