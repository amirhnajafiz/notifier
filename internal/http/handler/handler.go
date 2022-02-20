package handler

import "github.com/gofiber/fiber/v2"

func Publish(c *fiber.Ctx) error {
	return c.SendString("message published")
}

func Subscribe(c *fiber.Ctx) error {
	return c.SendString("message received")
}
