package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null;size:125"`
	Name     string `gorm:"not null;size:125"`
	Password string `gorm:"not null;size:125"`
	Posts    []Post `gorm:"foreignKey:AuthorID"`
}
