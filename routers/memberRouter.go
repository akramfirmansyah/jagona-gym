package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberRoutes(router fiber.Router) {
	// Member
	router.Post("/", controller.CreateMember)
	router.Get("/", controller.GetAllMember)
	router.Get("/:id", controller.GetMember)
	router.Put("/:id", controller.UpdateMember)
	router.Delete("/:id", controller.DeleteMember)
}