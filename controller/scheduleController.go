package controller

import (
	"errors"
	"strconv"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Create Schedule godoc
//
//	@Summary		Create Schedule
//	@Description	Create new Schedule data
//	@Tags			Schedule
//	@Accept			mpfd
//	@Produce		json
//	@Param			title		formData	string			true	"Title or name Schedule"
//	@Param			isallday	formData	string			false	"Is Schedule for all day?"
//	@Param			start_time	formData	string			true	"Start time schedule"
//	@Param			end_time	formData	string			true	"End time schedule"
//	@Success		200			{object}	models.Schedule	"Success create schedule"
//	@Failure		400			{string}	string			"Bad Request"
//	@Failure		500			{string}	string			"Internal Server Error"
//	@Router			/api/schedule [post]
func CreateSchedule(c *fiber.Ctx) error {
	isallday, err := strconv.ParseBool(c.FormValue("isallday", "false"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create schedule",
		})
	}

	newSchedule := models.Schedule{
		Title:     c.FormValue("title"),
		IsAllDay:  isallday,
		StartTime: c.FormValue("start_time"),
		EndTime:   c.FormValue("end_time"),
	}

	if err := database.DB.Create(&newSchedule).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create schedule",
		})
	}

	return c.JSON(newSchedule)
}

// GetAllSchedule godoc
//
//	@Summary	Get all schedules
//	@Tags		Schedule
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		models.Schedule
//	@Failure	500	{string}	string	"Internal Server Error"
//	@Router		/api/schedule [get]
func GetAllSchedule(c *fiber.Ctx) error {
	var schedule []models.Schedule

	if err := database.DB.Find(&schedule).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get schedule",
		})
	}

	return c.JSON(schedule)
}

// Update Schedule godoc
//
//	@Summary		Update Schedule
//	@Description	Update new Schedule data
//	@Tags			Schedule
//	@Accept			mpfd
//	@Produce		json
//
//	@Param			id			path		int				true	"Schedule ID"
//	@Param			title		formData	string			true	"Title or name Schedule"
//	@Param			isallday	formData	string			false	"Is Schedule for all day?"
//	@Param			start_time	formData	string			true	"Start time schedule"
//	@Param			end_time	formData	string			true	"End time schedule"
//
//	@Success		200			{object}	models.Schedule	"Success update schedule"
//	@Failure		400			{string}	string			"Bad Request"
//	@Failure		404			{string}	string			"Schedule not found"
//	@Failure		500			{string}	string			"Internal Server Error"
//	@Router			/api/schedule/{id} [put]
func UpdateSchedule(c *fiber.Ctx) error {
	id := c.Params("id")

	var schedule models.Schedule
	result := database.DB.Where("id = ?", id).First(&schedule)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Schedule not found!",
		})
	}

	isallday, err := strconv.ParseBool(c.FormValue("isallday", "false"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update schedule",
		})
	}

	database.DB.Where("id = ?", id).Updates(models.Schedule{
		Title:     c.FormValue("title"),
		IsAllDay:  isallday,
		StartTime: c.FormValue("start_time"),
		EndTime:   c.FormValue("end_time"),
	})

	database.DB.Where("id = ?", id).First(&schedule)

	return c.JSON(schedule)
}

// DeleteSchedule godoc
//
//	@Summary		Delete Schedule
//	@Description	Delete Schedule data by id
//	@Tags			Schedule
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path	int		true	"Schedule ID"
//	@Success		200	object	string	"Success Delete schedule"
//	@Failure		404	object	string	"Data not Found!"
//	@Router			/api/schedule/{id} [delete]
func DeleteSchedule(c *fiber.Ctx) error {
	id := c.Params("id")

	var schedule models.Schedule
	result := database.DB.Where("id = ?", id).First(&schedule)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Schedule not found!",
		})
	}

	if err := database.DB.Delete(&schedule, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Can't delete schedule",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success Delete Schedule Data",
	})
}
