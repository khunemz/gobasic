package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", welcome)

	log.Fatal(app.Listen(":3000"))
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}
