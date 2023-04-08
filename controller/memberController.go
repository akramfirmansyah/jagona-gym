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
//	@Description	Creating new Member data
//	@Tags			Member
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.MemberBody	true	"Data member"
//	@Success		200		{object}	models.Member		"Success create member"
//	@Failure		400		{string}	string				"Bad Request"
//	@Failure		500		{string}	string				"Internal Server Error"
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

	if err := database.DB.Preload("Trainer").Find(&member).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get member",
		})
	}

	return c.Status(fiber.StatusOK).JSON(member)
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
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Member ID"
//	@Param			body	body		models.MemberBody	true	"Update Member data"
//	@Success		200		{object}	models.Member
//	@Failure		400		{string}	string	"Bad Request"
//	@Failure		404		{string}	string	"Member not found"
//	@Failure		500		{string}	string	"Internal Server Error"
//	@Router			/api/member/{id} [put]
func UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")

	body := new(models.MemberBody)
	if err := c.BodyParser(body); err != nil {
		return c.JSON(fiber.Map{
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

	
	member.MemberName = body.MemberName
	member.MemberNIK = body.MemberNIK
	member.MemberContact = body.MemberContact
	member.MemberEmail = body.MemberEmail
	member.MemberAddress = body.MemberAddress
	member.MemberGender = body.MemberGender
	member.MemberWight = body.MemberWight
	member.MemberPackage = body.MemberPackage
	member.TrainerID = body.TrainerID

	if err := database.DB.Model(&member).Updates(&member).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update member",
		})
	}

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
//	@Failure		500	object	string	"Internal Server Error"
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
