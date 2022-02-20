package config

import (
	"cmd/internal/broker"
	"cmd/internal/cmd/server"
)

func Default() Config {
	return Config{
		Broker: broker.Config{
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
