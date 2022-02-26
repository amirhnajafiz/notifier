package handler

import (
	"time"

	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	var req request.Request
	_ = c.BodyParser(&req)

	err := h.Client.Publish(req.Msg)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "OK",
		"length": len(req.Msg),
		"type":   "text",
		"time":   time.Now().Format(time.RFC822),
	})
}

func (h Handler) GetAllInCache(c *fiber.Ctx) error {
	return c.JSON(h.Client.Cache.Pull())
}

func (h Handler) Clear(c *fiber.Ctx) error {
	length := len(h.Client.Cache.Pull())

	h.Client.Cache.Mock()

	return c.JSON(fiber.Map{
		"total":  length,
		"status": "clear",
	})
}
