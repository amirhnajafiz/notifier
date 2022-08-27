package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amirhnajafiz/notifier/internal/http/request"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) Publish(topic string, msg string) error {
	if token := c.Connection.Publish(topic, 0, false, msg); token.Wait() && token.Error() != nil {
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
	var data request.Request

	err := json.Unmarshal(message.Payload(), &data)
	if err != nil {
		_ = fmt.Errorf("problem in message unmarshaling %w", err)
	}

	fmt.Printf("%s #%s: %s\n", data.Date, data.ID, data.Msg)
}

func (c *Client) Connect() mqtt.Token {
	return c.Connection.Connect()
}

func (c *Client) Disconnect() {
	c.Connection.Disconnect(250)
}
