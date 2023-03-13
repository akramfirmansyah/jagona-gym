package controller

import (
	"errors"
	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTrainer(c *fiber.Ctx) error {
	body := new(models.Trainer)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	var existing models.Trainer
	if err := database.DB.Where(&body).Preload("Members").First(&existing).Error; err != nil {
		database.DB.Create(&body)
	} else {
		return c.JSON(fiber.Map{
			"message": "Data already exist",
			"data":    existing,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success create new Trainer!",
		"data":    body,
	})
}

func GetAllTrainer(c *fiber.Ctx) error {
	var data []models.Trainer

	result := database.DB.Preload("Members").Find(&data)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   data,
		"result": result.RowsAffected,
	})
}

func GetTrainer(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.Trainer

	result := database.DB.Where("id = ?", id).Preload("Members").First(&data)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(data)
}

func UpdateTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.Trainer)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	trainer := models.Trainer{
		TrainerName:     body.TrainerName,
		TrainerNIK:      body.TrainerNIK,
		TrainerBirthday: body.TrainerBirthday,
		TrainerEmail:    body.TrainerEmail,
		TrainerContact:  body.TrainerContact,
		TrainerAddress:  body.TrainerAddress,
		TrainerGender:   body.TrainerGender,
	}
	var data models.Trainer
	result := database.DB.Where("id = ?", id).Preload("Members").First(&data)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		database.DB.Model(&models.Trainer{}).Where("id = ?", id).Updates(&trainer)
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}
	database.DB.Model(&models.Trainer{}).Where("id = ?", id).Preload("Members").First(&data)

	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	var trainer models.Trainer

	result := database.DB.First(&trainer, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}
	database.DB.Delete(&trainer, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Trainer Data",
	})
}
