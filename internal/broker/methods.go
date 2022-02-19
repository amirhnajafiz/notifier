package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (m MQTT) messagePubHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}
}

func (m MQTT) connectHandler() mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		fmt.Println("Connected")
	}
}

func (m MQTT) connectLostHandler() mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}
}
