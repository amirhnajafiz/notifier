package cmd

import (
	"log"

	"github.com/amirhnajafiz/notifier/internal/cache"
	"github.com/amirhnajafiz/notifier/internal/client"
	"github.com/amirhnajafiz/notifier/internal/cmd/server"
	"github.com/amirhnajafiz/notifier/internal/config"
	"github.com/amirhnajafiz/notifier/internal/http/handler"
)

func Execute() {
	cfg := config.New()

	clt := client.Client{
		Cfg:          cfg.Client,
		IsSubscriber: false,
		Rdb:          cache.New(cfg.Cache),
	}.Register()

	if token := clt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	app := server.New()

	handler.Handler{
		Client: clt,
	}.Register(app)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

	clt.Disconnect()
}
