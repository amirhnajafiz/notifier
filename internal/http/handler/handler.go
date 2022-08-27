package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/amirhnajafiz/notifier/internal/client"
	"github.com/amirhnajafiz/notifier/internal/http/request"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client client.Client
}

func (h Handler) publish(c *fiber.Ctx) error {
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

	return c.SendStatus(http.StatusNoContent)
}

func (h Handler) Register(app *fiber.App) Handler {
	app.Post("/api/send", h.publish)

	return h
}
