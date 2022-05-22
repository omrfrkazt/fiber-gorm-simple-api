package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("host=localhost user=postgres dbname=fiber_gorm password=postgres sslmode=disable"), &gorm.Config{})
}
