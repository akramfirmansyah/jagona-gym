package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name         string  `gorm:"index;size:255" json:"name"`
	Birthday     string  `gorm:"type:date" json:"birthday"`
	JoinDate     string  `gorm:"type:date" json:"join_date"`
	Email        string  `gorm:"Index;size:255" json:"email"`
	Contact      string  `gorm:"size:50" json:"contact"`
	Package      string  `gorm:"size:50" json:"package"`
	Status       string  `gorm:"size:50" json:"status"`
	NIK          string  `json:"nik"`
	Address      string  `gorm:"size:255" json:"address"`
	Gender       string  `gorm:"type:enum('male','female')" json:"gender"`
	Weight       uint16  `gorm:"default:0" json:"weight"`
	Instagram    string  `gorm:"size:255" json:"instagram"`
	TrainerRefer uint    `json:"trainer_id"`
	Trainer      Trainer `gorm:"foreignKey:TrainerRefer;references:id;" json:"trainer"`
}
