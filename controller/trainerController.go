package controller

import (
	"errors"
	"fmt"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/akramfirmansyah/jagona-gym/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type trainerRequest struct {
	Name           string `form:"name"`
	Email          string `form:"email"`
	Contact        string `form:"contact"`
	Instagram      string `form:"instagram"`
	Image          string `form:"image"`
	Description    string `form:"description"`
	Specialization string `form:"specialization"`
	NIK            uint   `form:"nik"`
	Birthday       string `form:"birthday"`
	Address        string `form:"address"`
	Gender         string `form:"gender"`
	Experience     string `form:"experience"`
	Achievement    string `form:"achievement"`
}

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
//	@Param			birthday		formData	string			true	"Tanggal Lahir trainer. example: 2006-01-02"
//	@Param			email			formData	string			true	"Email trainer"
//	@Param			contact			formData	string			true	"Kontak trainer"
//	@Param			instagram		formData	string			false	"Instagram trainer"
//	@Param			address			formData	string			true	"Alamat trainer"
//	@Param			gender			formData	string			true	"Gender trainer"	Enums(Male, Female)
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
	body := new(trainerRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	path, err := utils.SaveFileTrainer(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println(body.Birthday)
	// Parse time
	birthday, err := utils.ParseTime(body.Birthday)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newTrainer := models.Trainer{
		Name:           body.Name,
		Email:          body.Email,
		Contact:        body.Contact,
		Instagram:      body.Instagram,
		Image:          path,
		Description:    body.Description,
		Specialization: body.Specialization,
		TrainerDetail: models.TrainerDetail{
			NIK:         body.NIK,
			Birthday:    birthday,
			Address:     body.Address,
			Gender:      body.Gender,
			Experience:  body.Experience,
			Achievement: body.Achievement,
		},
	}

	database.DB.Create(&newTrainer)

	if err := database.DB.Save(&newTrainer).Error; err != nil {
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

	if err := database.DB.Find(&trainers).Error; err != nil {
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
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/api/trainer/{id} [get]
func GetTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	var trainer models.Trainer
	if err := database.DB.Preload("TrainerDetail.Members").Preload("TrainerDetail").First(&trainer, id).Error; err != nil {
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
//	@Param			name			formData	string	true	"Name trainer"
//	@Param			nik				formData	integer	true	"NIK trainer"
//	@Param			birthday		formData	string	true	"Tanggal Lahir trainer. example: 2006-01-02"
//	@Param			email			formData	string	true	"Email trainer"
//	@Param			contact			formData	string	true	"Kontak trainer"
//	@Param			instagram		formData	string	true	"Instagram trainer"
//	@Param			address			formData	string	true	"Alamat trainer"
//	@Param			gender			formData	string	true	"Gender trainer"	Enums(Male, Female)
//	@Param			description		formData	string	true	"Deskripsi singkat trainer"
//	@Param			experience		formData	string	true	"Pengalaman trainer"
//	@Param			specialization	formData	string	true	"Spesialisasi trainer"
//	@Param			achievement		formData	string	true	"Pencapaian/Sertifikasi/Lisensi trainer"
//	@Param			image			formData	file	false	"Image trainer"
//
//	@Success		200				{object}	models.Trainer
//	@Failure		400				{string}	string	"Bad Request"
//	@Failure		404				{string}	string	"Trainer not found"
//	@Failure		500				{string}	string	"Internal Server Error"
//
//	@Router			/api/trainer/{id} [put]
func UpdateTrainer(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(trainerRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var trainer models.Trainer
	result := database.DB.Where("id = ?", id).Preload("TrainerDetail.Members").Preload("TrainerDetail").First(&trainer)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Trainer not found!",
		})
	}

	// Parse time
	birthday, err := utils.ParseTime(body.Birthday)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	file, err := c.FormFile("image")
	if file != nil && err == nil {
		path, err := utils.SaveFileTrainer(c, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to save image to Server!",
			})
		}

		database.DB.Model(&models.Trainer{}).Where("id = ?", id).Updates(models.Trainer{
			Image: path,
		})
	}

	database.DB.Model(&models.Trainer{}).Where("id = ?", id).Updates(models.Trainer{
		Name:           body.Name,
		Email:          body.Email,
		Contact:        body.Contact,
		Instagram:      body.Instagram,
		Description:    body.Description,
		Specialization: body.Specialization,
	})

	database.DB.Model(&models.TrainerDetail{}).Where("trainer_id = ?", id).Updates(models.TrainerDetail{
		NIK:         body.NIK,
		Birthday:    birthday,
		Address:     body.Address,
		Gender:      body.Gender,
		Experience:  body.Experience,
		Achievement: body.Achievement,
	})

	database.DB.Where("id = ?", id).Preload("TrainerDetail.Members").Preload("TrainerDetail").First(&trainer)

	return c.JSON(trainer)
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

	return c.JSON(fiber.Map{
		"message": "Success Delete Trainer Data",
	})
}
