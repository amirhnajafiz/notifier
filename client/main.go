package main

import (
	"log"

	"cmd/internal/cache"
	"cmd/internal/client"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func load() client.Config {
	var instance client.Config

	k := koanf.New(".")

	// Load configuration
	if err := k.Load(file.Provider("config.json"), yaml.Parser()); err != nil {
		log.Printf("error loading config.json: %s\n", err)
	}

	// Unmarshalling configurations
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s\n", err)
	}

	return instance
}

func main() {
	c := client.Client{
		Cache:        &cache.Cache{},
		Cfg:          load(),
		IsSubscriber: true,
	}.Register()

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	select {}
}
