package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(router fiber.Router) {
	router.Post("/", controller.CreateUser)
	router.Get("/", controller.GetAllUser)
	router.Get("/:id", controller.GetUser)
	router.Put("/:id", controller.UpdateUser)
	router.Delete("/:id", controller.DeleteUser)
}
