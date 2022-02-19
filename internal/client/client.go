package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	Connection mqtt.Client
}

func (c Client) Register(opts *mqtt.ClientOptions) Client {
	c.Connection = mqtt.NewClient(opts)

	return c
}
