package main

import (
	"cmd/internal/cache"
	"cmd/internal/client"
)

var count = 10

func getConfig() client.Config {
	return client.Config{
		Host:     "broker.emqx.io",
		Port:     1883,
		ClientID: "go_mqtt_client2",
		Username: "emqx2",
		Password: "public",
		Debug:    false,
		Topic:    "chat/message",
	}
}

func main() {
	c := client.Client{
		Cache:        &cache.Cache{},
		Cfg:          getConfig(),
		IsSubscriber: true,
	}.Register()

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
