package controller

import (
	"errors"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/akramfirmansyah/jagona-gym/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Equipemnt godoc
//
//	@Summary		Create Equipment
//	@Description	Creating new Equipment data
//	@Tags			Equipment
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			name	formData	string				true	"Nama Alat"
//	@Param			count	formData	integer				true	"Jumlah Alat"
//	@Param			status	formData	string				true	"Status Alat: ready atau broken"
//	@Param			image	formData	file				true	"Foto Alat"
//
//	@Success		200		{object}	models.Equipment	"Success create equipment"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/api/equipment [post]
func CreateEquipment(c *fiber.Ctx) error {
	body := new(models.Equipment)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	path, err := utils.SaveFileEquipment(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	equipment := models.Equipment{
		Name:     body.Name,
		Count:    body.Count,
		Status:   body.Status,
		ImageURL: path,
	}

	if err := database.DB.Create(&equipment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Create Data Equipment",
			"data":    nil,
		})
	}

	return c.JSON(equipment)
}

// GetAllEquipment godoc
//
//	@Summary	Get all equipments
//	@Tags		Equipment
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		models.Equipment
//	@Failure	500	{string}	string	"Internal Server Error"
//	@Router		/api/equipment [get]
func GetAllEquipment(c *fiber.Ctx) error {
	var equipments []models.Equipment

	if err := database.DB.Find(&equipments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.JSON(equipments)
}

// GetEquipment godoc
//
//	@Summary		Get a single Equipment
//	@Description	Get a single Equipment data by ID
//	@Tags			Equipment
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Equipment ID"
//	@Success		200	{object}	models.Equipment
//	@Failure		404	{string}	string	"Equipment not found"
//	@Failure		500	{string}	string	"Failed to get equipment"
//	@Router			/api/equipment/{id} [get]
func GetEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	var equipment models.Equipment
	if err := database.DB.First(&equipment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Equipment not Found!",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get Equipment",
		})
	}

	return c.JSON(equipment)
}

// UpdateEquipment godoc
//
//	@Summary		Update Equipment
//	@Description	Update Equipment data by ID
//	@Tags			Equipment
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			id		path		integer				true	"Equipment ID"
//	@Param			name	formData	string				true	"Nama Alat"
//	@Param			count	formData	integer				true	"Jumlah Alat"
//	@Param			status	formData	string				true	"Status Alat: ready atau broken"
//	@Param			image	formData	file				false	"Foto Alat"
//
//	@Success		200		{object}	models.Equipment	"Success create equipment"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/api/equipment/{id} [put]
func UpdateEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.Equipment)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	var equipment models.Equipment
	if err := database.DB.First(&equipment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Equipment not Found!",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update Equipment",
		})
	}

	file, err := c.FormFile("image")
	if file != nil && err == nil {
		path, err := utils.SaveFileEquipment(c, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save Image to Server!",
			})
		}

		equipment.ImageURL = path
	}

	equipment.Name = body.Name
	equipment.Count = body.Count
	equipment.Status = body.Status

	if err := database.DB.Model(&equipment).Updates(&equipment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to save data to Database!",
		})
	}

	return c.JSON(equipment)
}

// DeleteEquipment godoc
//
//	@Summary		Delete Equipment
//	@Description	Delete Equipment data by id
//	@Tags			Equipment
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int		true	"Equipment ID"
//	@Success		200	object	string	"Success Delete Equipment Data"
//	@Failure		404	object	string	"Data not Found!"
//	@Failure		500	object	string	"Internal Server Error"
//	@Router			/api/equipment/{id} [delete]
func DeleteEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	var equipment models.Equipment

	result := database.DB.First(&equipment, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Delete(&equipment, id)

	return c.JSON(fiber.Map{
		"message": "Success delete equipment data",
	})
}
