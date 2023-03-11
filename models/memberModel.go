package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name     string
	Email    string
	Contact  sql.NullString
	Address  string
	URLImage string
}
