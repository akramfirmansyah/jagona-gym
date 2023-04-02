package controller

import (
	"errors"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTrainer godoc
//
//	@Summary		Create Trainer
//	@Description	Creating new Trainer data
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.TrainerBody	true	"Data trainer"
//	@Success		200		{object}	models.Trainer		"Success create trainer"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/api/trainer [post]
func CreateTrainer(c *fiber.Ctx) error {
	body := new(models.TrainerBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	newTrainer := models.Trainer{
		TrainerName:     body.TrainerName,
		TrainerNIK:      body.TrainerNIK,
		TrainerBirthday: body.TrainerBirthday,
		TrainerEmail:    body.TrainerEmail,
		TrainerContact:  body.TrainerContact,
		TrainerAddress:  body.TrainerAddress,
		TrainerGender:   body.TrainerGender,
	}

	if err := database.DB.Create(&newTrainer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create trainer",
		})
	}

	return c.JSON(newTrainer)
}

// GetAllTrainer godoc
//
//	@Summary	Get all trainers
//	@Tags		Trainer
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		models.Trainer
//	@Failure	500	{string}	string	"Internal Server Error"
//	@Router		/api/trainer [get]
func GetAllTrainer(c *fiber.Ctx) error {
	var trainers []models.Trainer

	if err := database.DB.Preload("Members").Find(&trainers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get trainer",
		})
	}

	return c.JSON(trainers)
}

// GetTrainer godoc
//
//	@Summary		Get a single Trainer
//	@Description	Get a single Trainer data by ID
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Trainer ID"
//	@Success		200	{object}	models.Trainer
//	@Failure		404	{string}	string	"Trainer not found"
//	@Failure		500	{string}	string	"Failed to get trainer"
//	@Router			/api/trainer/{id} [get]
func GetTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	var trainer models.Trainer
	if err := database.DB.Preload("Members").First(&trainer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Trainer not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get trainer",
		})
	}

	return c.JSON(trainer)
}

// UpdateTrainer godoc
//
//	@Summary		Update Trainer
//	@Description	Update a specific Trainer data
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Trainer ID"
//	@Param			body	body		models.TrainerBody	true	"Update Trainer data"
//	@Success		200		{object}	models.Trainer
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		404		{string}	string	"Trainer not found"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/trainer/{id} [put]
func UpdateTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.TrainerBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var trainer models.Trainer
	result := database.DB.Where("id = ?", id).Preload("Members").First(&trainer)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Trainer not found",
		})
	}

	trainer.TrainerName = body.TrainerName
	trainer.TrainerNIK = body.TrainerNIK
	trainer.TrainerBirthday = body.TrainerBirthday
	trainer.TrainerEmail = body.TrainerEmail
	trainer.TrainerContact = body.TrainerContact
	trainer.TrainerAddress = body.TrainerAddress
	trainer.TrainerGender = body.TrainerGender

	if err := database.DB.Model(&trainer).Updates(&trainer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update trainer",
		})
	}

	return c.Status(fiber.StatusOK).JSON(trainer)
}

// DeleteTrainer godoc
//
//	@Summary		Delete Trainer
//	@Description	Delete Trainer data by id
//	@Tags			Trainer
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int		true	"Trainer ID"
//	@Success		200	object	string	"Success Delete Trainer Data"
//	@Failure		404	object	string	"Data not Found!"
//	@Failure		500	object	string	"Internal Server Error"
//	@Router			/api/trainer/{id} [delete]
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
