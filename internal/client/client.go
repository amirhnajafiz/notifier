package client

import (
	"cmd/internal/broker"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	Broker     broker.MQTT
	Connection mqtt.Client
}

func (c Client) Register(opts *mqtt.ClientOptions) Client {
	c.Connection = mqtt.NewClient(opts)

	return c
}
