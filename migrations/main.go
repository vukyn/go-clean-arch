package main

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Connect to database
	c := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.PostgreSQL.Host, c.PostgreSQL.User, c.PostgreSQL.Password, c.PostgreSQL.DBName, c.PostgreSQL.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
		&models.Todo{},
	)
}
