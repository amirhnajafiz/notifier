package client

import (
	"time"

	"cmd/internal/broker"
	"cmd/internal/cache"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	Broker     broker.MQTT
	Connection mqtt.Client
	Cache      cache.Cache
}

func (c Client) Register(opts *mqtt.ClientOptions) Client {
	c.Connection = mqtt.NewClient(opts)

	return c
}

func (c *Client) MessageHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		c.Cache.Put(cache.Message{
			Topic:   message.Topic(),
			Content: string(message.Payload()),
			Date:    time.Now(),
		})
	}
}
