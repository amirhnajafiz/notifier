package handler

import (
	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
	"time"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	var req request.Request
	_ = c.BodyParser(&req)

	err := h.Client.Publish(req.Topic, req.Msg)

	if err != nil {
		return err
	}

	return c.SendString("message published")
}

func (h Handler) Subscribe(c *fiber.Ctx) error {
	topic, msg, err := h.Client.Sub(c.Query("topic"))

	if err != nil {
		return err
	}

	return c.JSON(&fiber.Map{
		"message": msg,
		"topic":   topic,
		"date":    time.Now(),
	})
}
