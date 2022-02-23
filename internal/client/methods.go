package client

import "fmt"

func (c Client) Publish(topic string, msg string) error {
	if token := c.Connection.Publish(topic, 0, false, msg); token.Wait() && token.Error() != nil {
		return fmt.Errorf("publish error: %w", token.Error())
	}

	return nil
}

func (c Client) Sub(topic string) error {
	if token := c.Connection.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		return fmt.Errorf("subscribe failed %w", token.Error())
	}

	return nil
}
