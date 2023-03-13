package controller

import (
	"errors"
	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateMember(c *fiber.Ctx) error {
	body := new(models.Member)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	var existing models.Member
	if err := database.DB.Where("member_id = ?", body.MemberID).Preload("Trainer").First(&existing).Error; err != nil {
		database.DB.Preload("Trainer").Create(&body)
	} else {
		return c.JSON(fiber.Map{
			"message": "Data already exist",
			"data":    existing,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created new member!",
		"data":    body,
	})
}

func GetAllMember(c *fiber.Ctx) error {
	var data []models.Member

	total := database.DB.Preload("Trainer").Find(&data)
	if errors.Is(total.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":  data,
		"total": total.RowsAffected,
	})
}

func GetMember(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.Member

	err := database.DB.Preload("Trainer").Where("member_id = ?", id).First(&data).Error
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
		MemberID:      body.MemberID,
		MemberAddress: body.MemberAddress,
		MemberContact: body.MemberContact,
		MemberEmail:   body.MemberEmail,
		MemberGender:  body.MemberGender,
		MemberName:    body.MemberName,
		MemberNIK:     body.MemberNIK,
		MemberPackage: body.MemberPackage,
		MemberWight:   body.MemberWight,
		TrainerID:     body.TrainerID,
	}

	var data models.Member
	err := database.DB.Where("member_id = ?", id).Preload("Trainer").First(&data).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Model(&models.Member{}).Where("member_id = ?", id).Updates(&member)

	database.DB.Model(&models.Trainer{}).Where("id = ?", id).Preload("Members").First(&data)

	return c.Status(fiber.StatusOK).JSON(data)
}

func DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")

	var member models.Member

	err := database.DB.Where("member_id = ?", id).First(&member).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found!",
		})
	}

	database.DB.Where("member_id = ?", id).Delete(&member)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete Member Data",
	})
}
