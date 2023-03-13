package models

import "gorm.io/gorm"

// To Do
// Trainer Description (Experience / Certificate)

type Trainer struct {
	gorm.Model
	TrainerName     string `gorm:"index" json:"trainer_name"`
	TrainerNIK      uint   `json:"trainer_nik"`
	TrainerBirthday string `json:"trainer_birthday"`
	TrainerEmail    string `json:"trainer_email"`
	TrainerContact  string `gorm:"default:-" json:"trainer_contact"`
	TrainerAddress  string `json:"trainer_address"`
	TrainerGender   string `json:"trainer_gender"`
}
