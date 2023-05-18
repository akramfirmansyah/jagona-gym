package models

import (
	"gorm.io/gorm"
)

type TrainerDetail struct {
	gorm.Model
	TrainerID   uint
	NIK         string   `json:"nik"`
	Birthday    string   `json:"birthday"`
	Address     string   `gorm:"size:255" json:"address"`
	Gender      string   `gorm:"type:enum('male','female')" json:"gender"`
	Experience  string   `gorm:"size:255" json:"experience"`
	Achievement string   `gorm:"size:255" json:"achievement"`
	Members     []Member `gorm:"foreignKey:TrainerID" json:"members"`
}
