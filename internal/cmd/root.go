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

	mqtt := broker.MQTT{
		Cfg: cfg.Broker,
	}

	clt := client.Client{
		Broker: mqtt,
	}.Register(mqtt.Register())

	if token := clt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	app := server.New()
	handler.Handler{
		Client: clt,
	}.Register(app)

	err := app.Listen(cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}

	clt.Disconnect()
}
