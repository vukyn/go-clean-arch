package application

import (
	"boilerplate-clean-arch/config"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	authRepo "boilerplate-clean-arch/application/domains/auth/repository"
	authUseCase "boilerplate-clean-arch/application/domains/auth/usecase"
)

type Server struct {
	cfg   *config.Config
	db    *gorm.DB
	dbCfg *sql.DB
}

func (s *Server) run() {

	//Auth Service
	// authService := authService.InitauthService(s.redisClient, s.cfg, cache)
	authRepo := authRepo.NewAuthRepo(s.db)
	authUsecase := authUseCase.NewAuthUseCase(authRepo)

	fmt.Printf("starting gRPC server at port %v ...", 1234)
}
