package client

import (
	"time"
)

func (c Client) Publish(topic string, msg string) {
	token := c.Connection.Publish(topic, 0, false, msg)
	token.Wait()
	time.Sleep(time.Second)
}

func (c Client) Sub(topic string) {
	token := c.Connection.Subscribe(topic, 1, nil)
	token.Wait()
}
