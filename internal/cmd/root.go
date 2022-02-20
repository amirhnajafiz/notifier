package cmd

import (
	"cmd/internal/broker"
	"cmd/internal/client"
	"cmd/internal/cmd/server"
	"cmd/internal/config"
	"cmd/internal/http/handler"
	"log"
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

	app := server.New()
	handler.Handler{
		Client: clt,
	}.Register(app)

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal(err)
	}
}
