package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name      string  `gorm:"index" json:"name"`
	NIK       uint    `json:"nik"`
	Contact   string  `json:"contact"`
	Email     string  `gorm:"uniqueIndex;size:256" json:"email"`
	Address   string  `json:"address"`
	Gender    string  `json:"gender"`
	Wight     uint16  `gorm:"index;default:0" json:"wight"`
	Package   string  `json:"package"`
	TrainerID uint    `json:"trainer_id"`
	Trainer   Trainer `gorm:"foreignKey:TrainerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trainer"`
}

type MemberBody struct {
	Name      string `json:"name"`
	NIK       uint   `json:"nik"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Wight     uint16 `json:"wight"`
	Package   string `json:"package"`
	TrainerID uint   `json:"trainer_id"`
}
