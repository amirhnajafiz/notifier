package client

import (
	"fmt"
	"log"
	"time"

	"cmd/internal/cache"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) Publish(msg string) error {
	if token := c.Connection.Publish(c.Cfg.Topic, 0, false, msg); token.Wait() && token.Error() != nil {
		return fmt.Errorf("publish error: %w", token.Error())
	}

	return nil
}

func (c *Client) Sub(mqtt.Client) {
	if token := c.Connection.Subscribe(c.Cfg.Topic, 0, c.MessageHandler); token.Wait() && token.Error() != nil {
		log.Fatal(fmt.Errorf("failed subscrib %w", token.Error()))
	}
}

func (c *Client) MessageHandler(_ mqtt.Client, message mqtt.Message) {
	c.Cache.Put(cache.Message{
		Topic:   message.Topic(),
		Content: string(message.Payload()),
		Date:    time.Now(),
	})
}

func (c *Client) Connect() mqtt.Token {
	return c.Connection.Connect()
}

func (c *Client) Disconnect() {
	c.Connection.Disconnect(250)
}
