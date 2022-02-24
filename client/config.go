package main

import "cmd/internal/client"

func GetConfig() client.Config {
	return client.Config{
		Host:     "broker.emqx.io",
		Port:     1883,
		ClientID: "go_mqtt_client",
		Username: "emqx",
		Password: "public",
		Debug:    false,
		Topic:    "chat/room",
	}
}
