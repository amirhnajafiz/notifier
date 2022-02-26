package handler

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"cmd/internal/http/request"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) Publish(c *fiber.Ctx) error {
	var req request.Request
	_ = c.BodyParser(&req)

	ctx := context.Background()
	req.ID, _ = os.Hostname()
	req.Date = time.Now().Format(time.RFC1123)

	val, e := h.Client.Rdb.Get(ctx, req.ID).Result()
	if e == nil {
		req.ID = val
	}

	data, er := json.Marshal(req)
	if er != nil {
		return er
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

func (h Handler) SetName(c *fiber.Ctx) error {
	ctx := context.Background()
	value := c.Query("nickname")
	key, _ := os.Hostname()

	if err := h.Client.Rdb.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status":   "200|OK",
		"name":     key,
		"nickname": value,
	})
}

func (h Handler) RemoveName(c *fiber.Ctx) error {
	ctx := context.Background()
	key, _ := os.Hostname()

	if err := h.Client.Rdb.Del(ctx, key); err != nil {
		return err.Err()
	}

	return c.JSON(fiber.Map{
		"status": "200|OK",
	})
}

func (h Handler) WhoAmI(c *fiber.Ctx) error {
	ctx := context.Background()
	key, _ := os.Hostname()

	value, err := h.Client.Rdb.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"nickname": value,
	})
}
