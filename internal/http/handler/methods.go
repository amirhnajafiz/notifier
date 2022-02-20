package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Content struct {
	Topic string `json:"topic" xml:"topic" form:"topic"`
	Msg   string `json:"message" xml:"message" form:"message"`
}

func (h Handler) Publish(c *fiber.Ctx) error {
	var cnt Content
	_ = c.BodyParser(&cnt)

	err := h.Client.Publish(cnt.Topic, cnt.Msg)

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
