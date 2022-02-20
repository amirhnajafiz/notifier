package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Register() {
	app := fiber.New()

	err := app.Listen(":3030")
	if err != nil {
		log.Fatal(err)
	}
}
