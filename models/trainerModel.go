package models

import "gorm.io/gorm"

type Trainer struct {
	gorm.Model
	TrainerName     string   `gorm:"index;size:255" json:"trainer_name"`
	TrainerNIK      uint     `json:"trainer_nik"`
	TrainerBirthday string   `json:"trainer_birthday"`
	TrainerEmail    string   `gorm:"uniqueIndex;size:255" json:"trainer_email"`
	TrainerContact  string   `json:"trainer_contact"`
	TrainerAddress  string   `json:"trainer_address"`
	TrainerGender   string   `json:"trainer_gender"`
	Members         []Member `gorm:"foreignKey:TrainerID" json:"members"`
}

type TrainerBody struct {
	TrainerName     string `json:"trainer_name"`
	TrainerNIK      uint   `json:"trainer_nik"`
	TrainerBirthday string `json:"trainer_birthday"`
	TrainerEmail    string `json:"trainer_email"`
	TrainerContact  string `json:"trainer_contact"`
	TrainerAddress  string `json:"trainer_address"`
	TrainerGender   string `json:"trainer_gender"`
}
