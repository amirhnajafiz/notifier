package handler

import "github.com/gofiber/fiber/v2"

func (h Handler) Publish(c *fiber.Ctx) error {
	h.Client.Publish(c.Params("topic"), c.Params("message"))

	return c.SendString("message published")
}

func (h Handler) Subscribe(c *fiber.Ctx) error {
	return c.SendString("message received")
}
