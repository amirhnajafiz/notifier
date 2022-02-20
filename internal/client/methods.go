package client

import (
	"fmt"
	"time"
)

func (c Client) Publish(topic string, msg string) {
	token := c.Connection.Publish(topic, 0, false, msg)
	token.Wait()
	time.Sleep(time.Second)
}

func (c Client) Sub() {
	topic := "topic/test"
	token := c.Connection.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
}
