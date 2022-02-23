package client

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c Client) Publish(topic string, msg string) error {
	if token := c.Connection.Publish(topic, 0, false, msg); token.Wait() && token.Error() != nil {
		return fmt.Errorf("publish error: %w", token.Error())
	}

	return nil
}

func (c Client) Sub(mqtt.Client) {
	if token := c.Connection.Subscribe("snapp/item", 1, c.MessageHandler); token.Wait() && token.Error() != nil {
		log.Fatal(fmt.Errorf("failed subscrib %w", token.Error()))
	}
}

func (c Client) Connect() mqtt.Token {
	return c.Connection.Connect()
}

func (c Client) Disconnect() {
	c.Connection.Disconnect(250)
}
