package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name     string `gorm:"index" json:"name"`
	Count    uint8  `json:"count"`
	Status   string `gorm:"type:enum('ready','broken')" json:"status"`
	ImageURL string `gorm:"size:255" json:"img_url"`
}
