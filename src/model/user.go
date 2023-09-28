package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
