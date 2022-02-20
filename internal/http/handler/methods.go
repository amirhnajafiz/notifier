package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	err := h.Client.Publish(c.Params("topic"), c.Params("message"))

	if err != nil {
		return err
	}

	return c.SendString("message published")
}

func (h Handler) Subscribe(c *fiber.Ctx) error {
	topic, msg, err := h.Client.Sub(c.Params("topic"))

	if err != nil {
		return err
	}

	return c.JSON(&fiber.Map{
		"message": msg,
		"topic":   topic,
		"date":    time.Now(),
	})
}
