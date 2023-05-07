package models

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name         string       `gorm:"index;size:255" json:"name"`
	Birthday     time.Time    `gorm:"type:date" json:"birthday"`
	JoinDate     time.Time    `gorm:"type:date" json:"join_date"`
	Email        string       `gorm:"Index;size:255" json:"email"`
	Contact      string       `gorm:"size:50" json:"contact"`
	Package      string       `gorm:"size:50" json:"package"`
	Status       string       `gorm:"size:50" json:"status"`
	TrainerID    uint         `json:"trainer_id"`
	Trainer      Trainer      `gorm:"foreignKey:TrainerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"trainer"`
	MemberDetail MemberDetail `json:"detail_member"`
}
