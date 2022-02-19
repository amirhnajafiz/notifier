package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MQTT struct {
	Cfg Config
}

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

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)
	publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
