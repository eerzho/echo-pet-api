package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text   string `gorm:"size:500;not null"`
	PostID uint   `gorm:"not null"`
	Post   Post   `gorm:"foreignKey:PostID"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
}
