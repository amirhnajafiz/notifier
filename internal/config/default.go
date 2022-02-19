package config

import "cmd/internal/broker"

func Default() Config {
	return Config{
		Broker: broker.Config{
			Host:     "broker.emqx.io",
			Port:     1883,
			ClientID: "go_mqtt_client",
			Username: "emqx",
			Password: "public",
		},
	}
}
