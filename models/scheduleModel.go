package models

import "gorm.io/gorm"

type Schedule struct {
	gorm.Model
	Title     string `json:"title"`
	IsAllDay  bool   `json:"is_all_day"`
	StartTime string `gorm:"type:datetime" json:"start_time"`
	EndTime   string `gorm:"type:datetime" json:"end_time"`
}
