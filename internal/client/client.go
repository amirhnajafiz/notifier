package client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type Client struct {
	Connection mqtt.Client
}

func (c Client) Register(opts *mqtt.ClientOptions) Client {
	c.Connection = mqtt.NewClient(opts)

	return c
}

func (c Client) Publish() {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := c.Connection.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func (c Client) Sub() {
	topic := "topic/test"
	token := c.Connection.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
