package initializers

import "teal-gopher/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
