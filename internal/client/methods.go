package client

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) Publish(msg string) error {
	if token := c.Connection.Publish(c.Cfg.Topic, 0, false, msg); token.Wait() && token.Error() != nil {
		return fmt.Errorf("publish error: %w", token.Error())
	}

	return nil
}

func (c *Client) Sub(mqtt.Client) {
	if c.IsSubscriber {
		if token := c.Connection.Subscribe(c.Cfg.Topic, 0, c.MessageHandler); token.Wait() && token.Error() != nil {
			log.Fatal(fmt.Errorf("failed subscrib %w", token.Error()))
		}
	}
}

func (c *Client) MessageHandler(_ mqtt.Client, message mqtt.Message) {
	fmt.Printf("%s #%d: %s\n", time.Now().Format(time.Kitchen), message.MessageID(), message.Payload())
}

func (c *Client) Connect() mqtt.Token {
	return c.Connection.Connect()
}

func (c *Client) Disconnect() {
	c.Connection.Disconnect(250)
}
