package initializers

import "croic/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{}, &models.Todo{}, &models.Pomodoro{}, &models.Activity{})
}
