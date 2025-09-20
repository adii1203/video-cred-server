package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error: loading env")
	}
	_ = InitDB()

	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,            // 16 KB for headers
		BodyLimit:      20 * 1024 * 1024, // 20 MB
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
		})
	})

	app.Listen(":8000")
}
