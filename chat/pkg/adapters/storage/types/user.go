package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string `gorm:"column:first_name"`
	LastName     string `gorm:"column:last_name"`
	Email        string `gorm:"column:email;unique;not null"`
	PasswordHash string `gorm:"column:password_hash;not null"`
}
