package main

import (
	"os"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/utils"

	"github.com/gofiber/fiber/v2"
)

func init() {
	utils.LoadEnv()
	database.ConnectDB()
	database.Migrate()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jagona Gym Project")
	})

	api := app.Group("/api")
	api.Get("/user", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"data":    nil,
		})
	})

	app.Listen(":" + os.Getenv("PORT"))
}
