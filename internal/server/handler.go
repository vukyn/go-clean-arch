package server

import (
	authRepository "boilerplate-clean-arch/internal/auth/repository"
	authUseCase "boilerplate-clean-arch/internal/auth/usecase"
	authHttp "boilerplate-clean-arch/internal/auth/delivery/http"
	"github.com/labstack/echo/v4"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	aRepo := authRepository.NewAuthRepository(s.db)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(aRepo)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC)

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers)
	
	return nil
}