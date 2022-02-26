package cmd

import (
	"log"

	"cmd/internal/cache"
	"cmd/internal/client"
	"cmd/internal/cmd/server"
	"cmd/internal/config"
	"cmd/internal/http/handler"
)

func Execute() {
	cfg := config.New()

	clt := client.Client{
		Cfg:          cfg.Client,
		IsSubscriber: false,
		Rdb:          cache.New(),
	}.Register()

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
