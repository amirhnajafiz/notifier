package main

import (
	"cmd/internal/cache"
	"cmd/internal/client"
	"fmt"
	"time"
)

var count = 10

func main() {
	c := client.Client{
		Cache: cache.Cache{},
		Cfg:   GetConfig(),
	}.Register()

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for i := 0; i < count; i++ {
		fmt.Println(c.Cache.Pull())
		c.Cache.Mock()

		time.Sleep(2 * time.Second)
	}
}
