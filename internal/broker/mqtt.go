package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	Cfg   Config
	Topic string
	Msg   string
}

func (m MQTT) Register() *mqtt.ClientOptions {
	broker := m.Cfg.Host
	port := m.Cfg.Port

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(m.Cfg.ClientID)
	opts.SetUsername(m.Cfg.Username)
	opts.SetPassword(m.Cfg.Password)
	opts.SetDefaultPublishHandler(m.messagePubHandler())
	opts.OnConnect = m.connectHandler()
	opts.OnConnectionLost = m.connectLostHandler()

	return opts
}
