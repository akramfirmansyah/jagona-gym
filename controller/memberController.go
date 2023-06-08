package controller

import (
	"errors"
	"strconv"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type memberRequest struct {
	Name      string `form:"name"`
	NIK       string `form:"nik"`
	Birthday  string `form:"birthday"`
	JoinDate  string `form:"joindate"`
	Email     string `form:"email"`
	Contact   string `form:"contact"`
	Instagram string `form:"instagram"`
	Address   string `form:"address"`
	Gender    string `form:"gender"`
	Weight    uint16 `form:"weight"`
	Package   string `form:"package"`
	Status    string `form:"status"`
	TrainerID uint   `form:"trainer_id"`
}

// Member godoc
//
//	@Summary		Create Member
//	@Description	Creating new Member data
//	@Tags			Member
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			name		formData	string			true	"Name member"
//	@Param			nik			formData	string			true	"NIK member"
//	@Param			birthday	formData	string			true	"Birthday member"
//	@Param			joindate	formData	string			true	"Join date member"
//	@Param			email		formData	string			true	"Email member"
//	@Param			contact		formData	string			true	"Contact member"
//	@Param			instagram	formData	string			false	"Instagram member"
//	@Param			address		formData	string			true	"Address member"
//	@Param			gender		formData	string			true	"Gender member"	Enums(Male, Female)
//	@Param			weight		formData	integer			false	"Weight member"
//	@Param			package		formData	string			true	"Package member"	Enums(bronze, silver, gold, platinum)
//	@Param			status		formData	string			true	"Status member"		Enums(active, nonactive)
//	@Param			trainer_id	formData	string			true	"Trainer ID"
//
//	@Success		200			{object}	models.Member	"Success create member"
//	@Failure		400			{string}	string			"Bad Request"
//	@Failure		500			{string}	string			"Internal Server Error"
//	@Router			/api/member [post]
func CreateMember(c *fiber.Ctx) error {
	body := new(memberRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	member := models.Member{
		Name:         body.Name,
		Birthday:     body.Birthday,
		JoinDate:     body.JoinDate,
		Contact:      body.Contact,
		Email:        body.Email,
		Package:      body.Package,
		Status:       body.Status,
		TrainerRefer: body.TrainerID,
		NIK:          body.NIK,
		Address:      body.Address,
		Gender:       body.Gender,
		Weight:       body.Weight,
		Instagram:    body.Instagram,
	}

	if err := database.DB.Create(&member).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create member",
		})
	}

	return c.JSON(member)
}

// GetAllMember godoc
//
//	@Summary	Get all members
//	@Tags		Member
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		models.Member
//	@Failure	500	{string}	string	"Internal Server Error"
//	@Router		/api/member [get]
func GetAllMember(c *fiber.Ctx) error {
	var member []models.Member

	var limit int = 10
	var offset int = 0

	if c.Query("limit") != "" {
		temp, _ := strconv.Atoi(c.Query("limit"))
		limit = temp
	}

	if c.Query("page") != "" {
		temp, _ := strconv.Atoi(c.Query("page"))
		offset = limit * (temp - 1)
	}

	result := database.DB.Offset(offset).Limit(limit).Preload("Trainer").Find(&member)

	if err := result.Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get member",
		})
	}

	return c.JSON(fiber.Map{
		"data":  member,
		"count": result.RowsAffected,
	})
}

// GetMember godoc
//
//	@Summary		Get a single Member
//	@Description	Get a single Member data by ID
//	@Tags			Member
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Member ID"
//	@Success		200	{object}	models.Member
//	@Failure		404	{string}	string	"Member not found"
//	@Failure		500	{string}	string	"Failed to get trainer"
//	@Router			/api/member/{id} [get]
func GetMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member
	if err := database.DB.Preload("Trainer").Where("id = ?", id).First(&member).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Member not Found!",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get member",
		})
	}

	return c.JSON(member)
}

// UpdateMember godoc
//
//	@Summary		Update Member
//	@Description	Update a specific Member data
//	@Tags			Member
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			id			path		integer	true	"ID Member"
//	@Param			name		formData	string	true	"Name member"
//	@Param			nik			formData	string	true	"NIK member"
//	@Param			birthday	formData	string	true	"Birthday member"
//	@Param			joindate	formData	string	true	"Join date member"
//	@Param			email		formData	string	true	"Email member"
//	@Param			contact		formData	string	true	"Contact member"
//	@Param			instagram	formData	string	false	"Instagram member"
//	@Param			address		formData	string	true	"Address member"
//	@Param			gender		formData	string	true	"Gender member"	Enums(Male, Female)
//	@Param			weight		formData	integer	false	"Weight member"
//	@Param			package		formData	string	true	"Package member"	Enums(bronze, silver, gold, platinum)
//	@Param			status		formData	string	true	"Status member"		Enums(active, nonactive)
//	@Param			trainer_id	formData	string	true	"Trainer ID"
//
//	@Success		200			{object}	models.Member
//	@Failure		400			{string}	string	"Bad Request"
//	@Failure		404			{string}	string	"Member not found"
//	@Failure		500			{string}	string	"Internal Server Error"
//	@Router			/api/member/{id} [put]
func UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(memberRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var member models.Member
	result := database.DB.Where("id = ?", id).Preload("Trainer").First(&member)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Member not found!",
		})
	}

	database.DB.Model(&models.Member{}).Where("id = ?", id).Updates(models.Member{
		Name:      body.Name,
		Birthday:  body.Birthday,
		JoinDate:  body.JoinDate,
		Email:     body.Email,
		Contact:   body.Contact,
		Package:   body.Package,
		Status:    body.Status,
		NIK:       body.NIK,
		Address:   body.Address,
		Gender:    body.Gender,
		Weight:    body.Weight,
		Instagram: body.Instagram,
		// TrainerRefer: body.TrainerID,
	})

	database.DB.Where("id = ?", id).First(&member)

	return c.JSON(member)
}

// DeleteMember godoc
//
//	@Summary		Delete Member
//	@Description	Delete Member data by id
//	@Tags			Member
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int		true	"Member ID"
//	@Success		200	object	string	"Success Delete Member Data"
//	@Failure		404	object	string	"Data not Found!"
//	@Router			/api/member/{id} [delete]
func DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member

	result := database.DB.First(&member, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Delete(&member, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Member Data",
	})
}
