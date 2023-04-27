package routers

import (
	"github.com/akramfirmansyah/jagona-gym/controller"
	"github.com/gofiber/fiber/v2"
)

// TODO: Create Dashboard Route
func RegisterDashboardRoutes(router fiber.Router) {
	router.Get("/", controller.GetDashboard)
}
