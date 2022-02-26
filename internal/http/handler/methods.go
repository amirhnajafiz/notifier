package handler

import (
	"encoding/json"
	"os"
	"time"

	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	var req request.Request
	_ = c.BodyParser(&req)

	req.ID, _ = os.Hostname()
	req.Date = time.Now().Format(time.RFC1123)

	data, er := json.Marshal(req)
	if er != nil {
		panic(er)
	}

	err := h.Client.Publish(string(data))

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
