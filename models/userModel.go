package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string `gorm:"unique"`
	Password    string
	DateOfBirth string
	Activities  []Activity
	Todos       []Todo
	Pomodoros   []Pomodoro
}
