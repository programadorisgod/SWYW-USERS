package main

import (
	"log"
	"swyw-users/src/config"
	userController "swyw-users/src/controllers/users"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
		AllowCredentials: false,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("I'm healthy")
	})

	app.Post("/api/register", userController.CreateUser)
	app.Post("/api/login", userController.AuthenticateUser)
	app.Get("/api/users", userController.GetUserByField)

	log.Fatal(app.Listen("0.0.0.0:5002"))
}
