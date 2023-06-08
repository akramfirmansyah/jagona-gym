package main

import (
	"os"

	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/akramfirmansyah/jagona-gym/middlewares"
	"github.com/akramfirmansyah/jagona-gym/routers"
	"github.com/gofiber/swagger"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/utils"

	_ "github.com/akramfirmansyah/jagona-gym/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	utils.LoadEnv()
	database.ConnectDB()
	database.Migrate()
	database.Seeder()
}

//	@title						Jagona Gym API
//	@version					0.1.0
//	@description				This is a swagger for Jagona Gym API
//	@contact.name				API Support
//	@contact.email				akram.firman@gmail.com
//
//	@securityDefinitions.apikey	ApiKeyAuth
//
//	@host						localhost:8080
//	@BasePath					/
func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Jagona Gym API v0.0.1",
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/images", "./storage")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Jagona Gym Project")
	})

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Post("/login", controller.Login)

	api := app.Group("/api", middlewares.Protected())

	// User
	routers.RegisterUserRoutes(api.Group("/user"))

	// Dashboard
	routers.RegisterDashboardRoutes(api.Group("/dashboard"))

	// Trainer
	routers.RegisterTrainerRoutes(api.Group("/trainer"))

	// Member
	routers.RegisterMemberRoutes(api.Group("/member"))

	// Equipment
	routers.RegisterEquipmetRoutes(api.Group("/equipment"))

	// Schedule
	routers.RegisterScheduleRoutes(api.Group("/schedule"))

	_ = app.Listen(":" + os.Getenv("PORT"))
}
