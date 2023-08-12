package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imMatheus/pictionary/server/words"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Get("/words", func(c *fiber.Ctx) error {
		randomWords := words.GetRandomWords(20)

		return c.JSON(randomWords)
	})

	app.Listen(":3000")
}
