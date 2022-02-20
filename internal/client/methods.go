package client

func (c Client) Publish(topic string, msg string) error {
	if token := c.Connection.Publish(topic, 0, false, msg); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (c Client) Sub(topic string) (string, string, error) {
	if token := c.Connection.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		return "", "", token.Error()
	}

	return c.Broker.Topic, c.Broker.Msg, nil
}
