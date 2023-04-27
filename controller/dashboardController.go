package controller

import (
	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
)

// TODO: tambahkan fitur untuk menghandle permintaan ke dashboard
// - Statistik member baru 6 bulan terakhir
// - Statistik member aktif 6 bulan 6 terakhir
// - Display Member aktif bulan ini
// - Display jumlah member yang hadir /hari

func GetDashboard(c *fiber.Ctx) error {

	var memberActive int64
	if err := database.DB.Model(&models.Member{}).Where("status = ?", "active").Count(&memberActive).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var trainerCount int64
	if err := database.DB.Model(&models.Trainer{}).Count(&trainerCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var equipmentCount int64
	if err := database.DB.Model(&models.Equipment{}).Count(&equipmentCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"statistics_new_member":    "default",
		"statistics_active_member": "default",
		"present_member":           "default",
		"active_member":            memberActive,
		"number_of_trainer":        trainerCount,
		"number_of_equip":          equipmentCount,
	})
}
