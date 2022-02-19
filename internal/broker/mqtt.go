package broker

import (
	"cmd/internal/client"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	Cfg Config
}

func (m MQTT) Register(cfg Config) {
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

	clt := client.Client{}.Register(opts)

	if token := clt.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	clt.Publish()
	clt.Sub()

	clt.Disconnect()
}
