package handler

import (
	"encoding/json"
	"time"

	"github.com/amirhnajafiz/notifier/internal/http/request"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	var req request.Request

	if err := c.BodyParser(&req); err != nil {
		return c.SendString(err.Error())
	}

	req.Date = time.Now().Format(time.RFC1123)

	data, er := json.Marshal(req)
	if er != nil {
		return er
	}

	err := h.Client.Publish(req.Topic, string(data))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":    "200|OK",
		"length":    len(req.Msg),
		"send from": req.ID,
		"time":      req.Date,
	})
}
