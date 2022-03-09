package server

import "github.com/gofiber/fiber/v2"

func getConfig() fiber.Config {
	return fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0", "1.1.1.1/30"}, // IP address or IP address range
		ProxyHeader:             fiber.HeaderXForwardedFor,
	}
}
