package config

import (
	"github.com/amirhnajafiz/notifier/internal/client"
)

func Default() Config {
	return Config{
		Client: client.Config{
			Host:     "broker.emqx.io",
			Port:     1883,
			ClientID: "go_mqtt_client",
			Username: "emqx",
			Password: "public",
			Debug:    false,
			Topic:    "chat/room",
		},
	}
}
