package models

import "gorm.io/gorm"

type Pomodoro struct {
	gorm.Model
	FocusTime string `gorm:"unique"`
	BreakTime string
	UserID    uint
}
