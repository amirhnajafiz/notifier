package handler

import (
	"cmd/internal/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Register(app *fiber.App) Handler {
	app.Post("/api/send", h.Publish)
	app.Use("/clone", websocket.New(h.Subscribe))

	return h
}
