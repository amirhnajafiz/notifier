package client

import mqtt "github.com/eclipse/paho.mqtt.golang"

func Register(opts *mqtt.ClientOptions) mqtt.Client {
	return mqtt.NewClient(opts)
}
