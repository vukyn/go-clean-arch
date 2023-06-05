package main

import (
	"boilerplate-clean-arch/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=go_demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Register new user
	// user := &models.User{
	// 	UserId:    uuid.New(),
	// 	FirstName: "Michael",
	// 	LastName:  "Smith",
	// 	Email:     "09_michael@email.com",
	// 	Password: "123456",
	// }
	// register(db, user)

	// Get user
	fmt.Printf("User: %v", getUser(db, "0ecda023-975a-48d0-aa1c-72243e3b71e6"))
}

func register(db *gorm.DB, user *models.User) {
	// Hash password
	err := user.HashPassword()
	if err != nil {
		panic("error hashing password")
	}
	// Create
	db.Create(user)
}

func getUser(db *gorm.DB, UserId string) *models.User {
	user := &models.User{}
	db.First(user, "user_id = ?", UserId)
	return user
}
