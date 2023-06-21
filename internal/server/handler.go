package server

import (
	_ "boilerplate-clean-arch/docs"
	authHttp "boilerplate-clean-arch/internal/auth/delivery/http"
	authRepository "boilerplate-clean-arch/internal/auth/repository"
	authUseCase "boilerplate-clean-arch/internal/auth/usecase"
	sessionRepository "boilerplate-clean-arch/internal/session/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	aRepo := authRepository.NewAuthRepository(s.db)
	sRepo := sessionRepository.NewSessionRepository(s.cfg, s.redisClient)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, aRepo, sRepo)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers)

	return nil
}
