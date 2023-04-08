package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterTrainerRoutes(router fiber.Router) {
	// Trainer
	router.Post("/", controller.CreateTrainer)
	router.Get("/", controller.GetAllTrainer)
	router.Get("/:id", controller.GetTrainer)
	router.Put("/:id", controller.UpdateTrainer)
	router.Delete("/:id", controller.DeleteTrainer)
}
