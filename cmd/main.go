package main

import (
	"log"

	"github.com/adii1203/video-cred/internals/handlers"
	"github.com/adii1203/video-cred/internals/service"
	"github.com/adii1203/video-cred/internals/storage"
	"github.com/adii1203/video-cred/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error: loading env")
	}
	db := pkg.InitDB()
	logger := pkg.NewLogger()

	repo := storage.New(db)
	user_service := service.NewUserService(repo)
	user_handler := handlers.NewUserHandler(user_service, logger)

	app := fiber.New(fiber.Config{
		ReadBufferSize: 16384,            // 16 KB for headers
		BodyLimit:      20 * 1024 * 1024, // 20 MB
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello",
		})
	})

	app.Post("/clerk", user_handler.ClerkHandler())

	app.Listen(":8000")
}
