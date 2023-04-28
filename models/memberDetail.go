package models

import "gorm.io/gorm"

type MemberDetail struct {
	gorm.Model
	MemberID  uint
	NIK       uint   `json:"nik"`
	Address   string `gorm:"size:255" json:"address"`
	Gender    string `gorm:"type:enum('male','female')" json:"gender"`
	Weight    uint16 `gorm:"default:0" json:"weight"`
	Instagram string `gorm:"size:255" json:"instagram"`
}
