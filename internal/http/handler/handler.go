package handler

import (
	"github.com/amirhnajafiz/Notifier/internal/client"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Register(app *fiber.App) Handler {
	app.Post("/api/send", h.Publish)
	app.Put("/api/name", h.SetName)
	app.Get("/api/name", h.WhoAmI)
	app.Delete("/api/name", h.RemoveName)

	return h
}
