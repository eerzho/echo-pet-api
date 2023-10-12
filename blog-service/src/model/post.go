package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Slug     string `gorm:"index;unique;not null;size:255"`
	Title    string `gorm:"not null;size:255"`
	Desc     string `gorm:"size:500"`
	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`
}
