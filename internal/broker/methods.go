package broker

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (m MQTT) connectHandler() mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		return
	}
}

func (m MQTT) connectLostHandler() mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		return
	}
}
