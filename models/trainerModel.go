package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	Name        string `json:"trainer_name"`
	Email       string `json:"trainer_email"`
	PhoneNumber string `json:"phone_number"`
	ImgURL      string `json:"image_url"`
}
