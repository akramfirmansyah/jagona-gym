package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterScheduleRoutes(router fiber.Router) {
	router.Post("/", controller.CreateSchedule)
	router.Get("/", controller.GetAllSchedule)
	router.Put("/:id", controller.UpdateSchedule)
	router.Delete("/:id", controller.DeleteSchedule)
}
