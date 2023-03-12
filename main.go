package main

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"os"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/json-iterator/go"
)

func init() {
	utils.LoadEnv()
	database.ConnectDB()
	database.Migrate()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Jagona Gym API v0.0.1",
		JSONDecoder: jsoniter.Unmarshal,
		JSONEncoder: jsoniter.Marshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jagona Gym Project")
	})

	api := app.Group("/api")

	// Trainer
	api.Post("/trainer", controller.CreateTrainer)
	api.Get("/trainer", controller.GetAllTrainer)
	api.Get("/trainer/:id", controller.GetTrainer)
	api.Put("/trainer/:id", controller.UpdateTrainer)
	api.Delete("/trainer/:id", controller.DeleteTrainer)

	_ = app.Listen(":" + os.Getenv("PORT"))
}
