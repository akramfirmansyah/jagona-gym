package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Name     string   `gorm:"index;size:255" json:"name"`
	NIK      uint     `json:"nik"`
	Birthday string   `json:"birthday"`
	Email    string   `gorm:"uniqueIndex;size:255" json:"email"`
	Contact  string   `json:"contact"`
	Address  string   `json:"address"`
	Gender   string   `json:"gender"`
	Image    string   `json:"img_url"`
	Members  []Member `gorm:"foreignKey:TrainerID" json:"members"`
}

type TrainerBody struct {
	Name     string `form:"name"`
	NIK      uint   `form:"nik"`
	Birthday string `form:"birthday"`
	Email    string `form:"email"`
	Contact  string `form:"contact"`
	Address  string `form:"address"`
	Gender   string `form:"gender"`
	Image    string `form:"image"`
}
