package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name      string    `gorm:"index;size:255" json:"name"`
	NIK       uint      `json:"nik"`
	Birthday  time.Time `gorm:"type:date" json:"birthday"`
	Email     string    `gorm:"Index;size:255" json:"email"`
	Contact   string    `gorm:"size:50" json:"contact"`
	Instagram string    `gorm:"size:255" json:"instagram"`
	Address   string    `gorm:"size:255" json:"address"`
	Gender    string    `gorm:"type:enum('male','female')" json:"gender"`
	Weight    uint16    `gorm:"default:0" json:"weight"`
	Package   string    `gorm:"size:50" json:"package"`
	Status    string    `gorm:"size:50" json:"status"`
	TrainerID uint      `json:"trainer_id"`
	Trainer   Trainer   `gorm:"foreignKey:TrainerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trainer"`
}

type MemberBody struct {
	Name      string `json:"name"`
	NIK       uint   `json:"nik"`
	Birthday  string `json:"birthday"`
	Email     string `json:"email"`
	Contact   string `json:"contact"`
	Instagram string `json:"instagram"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Weight    uint16 `json:"weight"`
	Package   string `json:"package"`
	Status    string `json:"status"`
	TrainerID uint   `json:"trainer_id"`
}
