package handler

import (
	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
	"time"
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
