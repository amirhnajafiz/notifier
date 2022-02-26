package client

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Connection   mqtt.Client
	Cfg          Config
	IsSubscriber bool
	Rdb          *redis.Client
}

func (c Client) Register() Client {
	if c.Cfg.Debug {
		mqtt.DEBUG = log.New(os.Stdout, "mqtt", 0)
		mqtt.ERROR = log.New(os.Stdout, "mqtt", 0)
	}

	broker := c.Cfg.Host
	port := c.Cfg.Port

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(c.Cfg.ClientID)
	opts.SetUsername(c.Cfg.Username)
	opts.SetPassword(c.Cfg.Password)
	opts.SetPingTimeout(1 * time.Second)

	opts.OnConnect = c.Sub

	c.Connection = mqtt.NewClient(opts)

	return c
}
