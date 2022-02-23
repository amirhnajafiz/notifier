package config

import (
	"cmd/internal/client"
	"cmd/internal/cmd/server"
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
		},
		Server: server.Config{
			Port: ":3030",
		},
	}
}
