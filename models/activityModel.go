package models

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	About  string `gorm:"unique"`
	UserID uint
}
