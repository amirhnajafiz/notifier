package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	Cfg Config
}

func (m MQTT) Register(cfg Config) *mqtt.ClientOptions {
	broker := cfg.Host
	port := cfg.Port

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetDefaultPublishHandler(m.messagePubHandler())
	opts.OnConnect = m.connectHandler()
	opts.OnConnectionLost = m.connectLostHandler()

	return opts
}
