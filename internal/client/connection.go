package client

import mqtt "github.com/eclipse/paho.mqtt.golang"

func (c Client) Connect() mqtt.Token {
	return c.Connection.Connect()
}

func (c Client) Disconnect() {
	c.Connection.Disconnect(250)
}
