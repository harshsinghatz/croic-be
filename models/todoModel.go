package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Status      string
	DueDate     string
	UserID      uint
}
