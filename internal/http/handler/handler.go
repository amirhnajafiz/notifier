package handler

import (
	"github.com/amirhnajafiz/notifier/internal/client"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Register(app *fiber.App) Handler {
	app.Post("/api/send", h.Publish)

	return h
}
