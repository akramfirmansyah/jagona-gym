package controller

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/akramfirmansyah/jagona-gym/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateTrainer godoc
//
//	@Summary		Create Trainer
//	@Description	Creating new Trainer data
//	@Tags			Trainer
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			name			formData	string			true	"Name trainer"
//	@Param			nik				formData	integer			true	"NIK trainer"
//	@Param			birthday		formData	string			true	"Tanggal Lahir trainer"
//	@Param			email			formData	string			true	"Email trainer"
//	@Param			contact			formData	string			true	"Kontak trainer"
//	@Param			instagram		formData	string			false	"Instagram trainer"
//	@Param			address			formData	string			true	"Alamat trainer"
//	@Param			gender			formData	string			true	"Gender trainer"
//	@Param			description		formData	string			false	"Deskripsi singkat trainer"
//	@Param			experience		formData	string			false	"Pengalaman trainer"
//	@Param			specialization	formData	string			true	"Spesialisasi trainer"
//	@Param			achievement		formData	string			false	"Pencapaian/Sertifikasi/Lisensi trainer"
//	@Param			image			formData	file			true	"Image trainer"
//
//	@Success		200				{object}	models.Trainer	"Success create trainer"
//	@Failure		400				{string}	string			"Bad Request"
//	@Failure		500				{string}	string			"Internal Server Error"
//	@Router			/api/trainer [post]
func CreateTrainer(c *fiber.Ctx) error {
	body := new(models.TrainerBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	path, err := utils.SaveFileTrainer(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Parse time
	birthday, err := utils.ParseTime(body.Birthday)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newTrainer := models.Trainer{
		Name:           body.Name,
		NIK:            body.NIK,
		Birthday:       birthday,
		Email:          body.Email,
		Contact:        body.Contact,
		Instagram:      body.Instagram,
		Address:        body.Address,
		Gender:         body.Gender,
		Image:          path,
		Description:    body.Description,
		Experience:     body.Experience,
		Specialization: body.Specialization,
		Achievement:    body.Achievement,
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
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			id				path		int		true	"Trainer ID"
//	@Param			name			formData	string	false	"Name trainer"
//	@Param			nik				formData	integer	false	"NIK trainer"
//	@Param			birthday		formData	string	false	"Tanggal Lahir trainer"
//	@Param			email			formData	string	false	"Email trainer"
//	@Param			contact			formData	string	false	"Kontak trainer"
//	@Param			instagram		formData	string	false	"Instagram trainer"
//	@Param			address			formData	string	false	"Alamat trainer"
//	@Param			gender			formData	string	false	"Gender trainer"
//	@Param			description		formData	string	false	"Deskripsi singkat trainer"
//	@Param			experience		formData	string	false	"Pengalaman trainer"
//	@Param			specialization	formData	string	false	"Spesialisasi trainer"
//	@Param			achievement		formData	string	false	"Pencapaian/Sertifikasi/Lisensi trainer"
//	@Param			image			formData	file	false	"Image trainer"
//
//	@Success		200				{object}	models.Trainer
//	@Failure		400				{string}	string	"Bad Request"
//	@Failure		404				{string}	string	"Trainer not found"
//	@Failure		500				{string}	string	"Internal Server Error"
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
			"message": "Trainer not found!",
		})
	}

	file, err := c.FormFile("image")
	if file != nil && err == nil {
		// Generate unique filename
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

		// Save file to disk
		if err := c.SaveFile(file, fmt.Sprintf("public/trainer/%s", filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		trainer.Image = fmt.Sprintf("%s/images/trainer/%s", os.Getenv("BASE_URL"), filename)
	}

	// Parse time
	birthday, err := utils.ParseTime(body.Birthday)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	trainer.Name = body.Name
	trainer.NIK = body.NIK
	trainer.Birthday = birthday
	trainer.Email = body.Email
	trainer.Contact = body.Contact
	trainer.Instagram = body.Instagram
	trainer.Address = body.Address
	trainer.Gender = body.Gender
	trainer.Description = body.Description
	trainer.Experience = body.Experience
	trainer.Specialization = body.Specialization
	trainer.Achievement = body.Achievement

	fmt.Println(body.Image)

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
