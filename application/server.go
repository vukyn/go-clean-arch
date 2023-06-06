package application

import (
	"boilerplate-clean-arch/config"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	authRepo "boilerplate-clean-arch/application/domains/auth/repository"
	authUseCase "boilerplate-clean-arch/application/domains/auth/usecase"
	authHandler "boilerplate-clean-arch/application/domains/auth/delivery/handler"
)

type Server struct {
	cfg   *config.Config
	db    *gorm.DB
	dbCfg *sql.DB
}

func (s *Server) Start() {
	s.run()
}

func (s *Server) run() {

	//Auth Service
	authRepo := authRepo.NewAuthRepo(s.db)
	authUC := authUseCase.NewAuthUseCase(authRepo)
	authHandler := authHandler.NewAuthHandlers(s.cfg, authUC)
}