package main

import (
	"boilerplate-clean-arch/config"
	todo "boilerplate-clean-arch/internal/todo/entity"
	auth "boilerplate-clean-arch/internal/auth/entity"
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
	fmt.Println("Connect to database successfully")

	fmt.Println("Run migrate ...")

	// Migrate the schema
	db.AutoMigrate(&auth.User{})
	db.AutoMigrate(&todo.Todo{})

	fmt.Println("Migrate successfully")
}
