package handler

import (
	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
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
	all := h.Client.Cache.Pull()

	h.Client.Cache.Mock()

	return c.JSON(all)
}
