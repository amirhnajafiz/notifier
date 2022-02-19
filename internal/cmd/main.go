package cmd

import (
	"cmd/internal/config"
	"fmt"
)

func Execute() {
	cfg := config.New()
	fmt.Println(cfg.Broker)
}
