package utils

import (
	"time"

	"github.com/akramfirmansyah/jagona-gym/database"
	"github.com/akramfirmansyah/jagona-gym/models"
)

func GetMonthlyNewMembers() ([]struct {
	Month string
	Total uint
}, error) {
	var result []struct {
		Month string
		Total uint
	}

	// calculate six months ago from now
	database.DB.Model(&models.Member{}).
		Preload("Member").
		Select("MONTHNAME(created_at) AS month, COUNT(*) AS total").
		Where("created_at >= ?", time.Now().AddDate(0, -6, 0)).
		Group("MONTHNAME(created_at), MONTH(created_at)").
		Find(&result)

	return result, nil
}
