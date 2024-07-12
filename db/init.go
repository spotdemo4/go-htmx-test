package db

import (
	"go-htmx-test/models"
)

func AutoMigrate() {
	// Migrate the schema
	DB.AutoMigrate(&models.Item{})
}
