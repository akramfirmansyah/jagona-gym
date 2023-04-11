package models

import (
	"time"

	"gorm.io/gorm"
)

type Trainer struct {
	gorm.Model
	Name           string    `gorm:"index;size:255" json:"name"`
	NIK            uint      `json:"nik"`
	Birthday       time.Time `gorm:"type:date" json:"birthday"`
	Email          string    `gorm:"uniqueIndex;size:255" json:"email"`
	Contact        string    `gorm:"size:50" json:"contact"`
	Instagram      string    `gorm:"size:255" json:"instagram"`
	Address        string    `gorm:"size:255" json:"address"`
	Gender         string    `gorm:"type:enum('male','female')" json:"gender"`
	Image          string    `gorm:"size:255" json:"img_url"`
	Description    string    `gorm:"type:mediumtext" json:"description"`
	Experience     string    `gorm:"size:255" json:"experience"`
	Specialization string    `gorm:"size:255" json:"specialization"`
	Achievement    string    `gorm:"size:255" json:"achievement"`
	Members        []Member  `gorm:"foreignKey:TrainerID" json:"members"`
}

type TrainerBody struct {
	Name           string `form:"name"`
	NIK            uint   `form:"nik"`
	Birthday       string `form:"birthday"`
	Email          string `form:"email"`
	Contact        string `form:"contact"`
	Instagram      string `form:"instagram"`
	Address        string `form:"address"`
	Gender         string `form:"gender"`
	Image          string `form:"image"`
	Description    string `form:"description"`
	Experience     string `form:"experience"`
	Specialization string `form:"specialization"`
	Achievement    string `form:"achievement"`
}
