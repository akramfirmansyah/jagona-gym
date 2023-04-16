package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterEquipmetRoutes(router fiber.Router) {
	router.Post("/", controller.CreateEquipment)
	router.Get("/", controller.GetAllEquipment)
	router.Get("/:id", controller.GetEquipment)
	router.Put("/:id", controller.UpdateEquipment)
	router.Delete("/:id", controller.DeleteEquipment)
}
