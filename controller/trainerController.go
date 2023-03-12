package controller

import (
	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
)

func CreateTrainer(c *fiber.Ctx) error {
	body := new(models.Trainer)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	database.DB.Create(&body)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success create new user!",
		"data":    body,
	})
}

func GetAllTrainer(c *fiber.Ctx) error {
	var result []models.Trainer

	hasil := database.DB.Find(&result)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success!",
		"data":    result,
		"total":   hasil.RowsAffected,
	})
}

func GetTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.Trainer

	result := database.DB.Where("id = ?", id).First(&data)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success!",
		"data":    data,
	})
}

func UpdateTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.Trainer)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	trainer := models.Trainer{
		Name:        body.Name,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		ImgURL:      body.ImgURL,
	}

	database.DB.Model(models.Trainer{}).Where("id = ?", id).Updates(&trainer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success!",
		"data":    trainer,
	})
}

func DeleteTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Delete(&models.Trainer{}, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Trainer Data",
	})
}
