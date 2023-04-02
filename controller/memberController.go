package controller

import (
	"errors"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Member godoc
//
//	@Summary		Create Member
//	@Description	Creating new member data
//	@Tags			Member
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.MemberBody	true	"Data member"
//	@Success		200		{array}		models.Member		"Success create member"
//	@Failure		400		{string}	string				"ok"
//	@Failure		404		{string}	string				"ok"
//	@Failure		500		{string}	string				"ok"
//	@Router			/api/member [post]
func CreateMember(c *fiber.Ctx) error {
	body := new(models.MemberBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	member := models.Member{
		MemberName:    body.MemberName,
		MemberNIK:     body.MemberNIK,
		MemberContact: body.MemberContact,
		MemberEmail:   body.MemberEmail,
		MemberAddress: body.MemberAddress,
		MemberGender:  body.MemberGender,
		MemberWight:   body.MemberWight,
		MemberPackage: body.MemberPackage,
		TrainerID:     body.TrainerID,
	}

	result := database.DB.Create(&member)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func GetAllMember(c *fiber.Ctx) error {
	var data []models.Member

	total := database.DB.Find(&data)
	if errors.Is(total.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func GetMember(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.Member

	err := database.DB.Preload("Trainer").Where("id = ?", id).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.Member)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	member := models.Member{
		MemberAddress: body.MemberAddress,
		MemberContact: body.MemberContact,
		MemberEmail:   body.MemberEmail,
		MemberGender:  body.MemberGender,
		MemberName:    body.MemberName,
		MemberNIK:     body.MemberNIK,
		MemberPackage: body.MemberPackage,
		MemberWight:   body.MemberWight,
		//TrainerID:     body.TrainerID,
	}

	var data models.Member
	err := database.DB.Where("id = ?", id).Preload("Trainer").First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Model(&models.Member{}).Where("id = ?", id).Updates(&member)

	//database.DB.Model(&models.Trainer{}).Where("id = ?", id).Preload("Members").First(&data)

	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member

	err := database.DB.Where("id = ?", id).First(&member).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Where("id = ?", id).Delete(&member)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Member Data",
	})
}
