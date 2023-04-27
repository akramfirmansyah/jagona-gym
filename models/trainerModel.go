package models

import (
	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	Name           string `gorm:"index;size:255" json:"name"`
	Email          string `gorm:"index;size:255" json:"email"`
	Contact        string `gorm:"size:50" json:"contact"`
	Instagram      string `gorm:"size:255" json:"instagram"`
	Image          string `gorm:"size:255" json:"img_url"`
	Description    string `gorm:"type:mediumtext" json:"description"`
	Specialization string `gorm:"size:255" json:"specialization"`
	TrainerDetail  TrainerDetail
}
