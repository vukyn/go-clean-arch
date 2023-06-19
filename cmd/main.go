package main

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/server"
	"fmt"
	"github.com/labstack/gommon/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host petstore.swagger.io
// @BasePath api/v1
func main() {
	log.Info("Starting api server")

	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.PostgreSQL.Host, cfg.PostgreSQL.User, cfg.PostgreSQL.Password, cfg.PostgreSQL.DBName, cfg.PostgreSQL.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	}

	s := server.NewServer(cfg, db)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
