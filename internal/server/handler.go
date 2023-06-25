package server

import (
	_ "boilerplate-clean-arch/docs"
	authHttp "boilerplate-clean-arch/internal/auth/delivery/http"
	authRepository "boilerplate-clean-arch/internal/auth/repository"
	authUseCase "boilerplate-clean-arch/internal/auth/usecase"
	todoHttp "boilerplate-clean-arch/internal/todo/delivery/http"
	todoRepository "boilerplate-clean-arch/internal/todo/repository"
	todoUseCase "boilerplate-clean-arch/internal/todo/usecase"
	sessionRepository "boilerplate-clean-arch/internal/session/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	authRepo := authRepository.NewAuthRepository(s.db)
	tRepo := todoRepository.NewTodoRepository(s.db)
	sRepo := sessionRepository.NewSessionRepository(s.cfg, s.redisClient)

	// Init useCases
	authUC := authUseCase.NewAuthUseCase(s.cfg, authRepo, sRepo)
	todoUC := todoUseCase.NewTodoUseCase(tRepo)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC)
	todoHandlers := todoHttp.NewTodoHandlers(s.cfg, todoUC)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.Logger())

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todo")

	authHttp.MapAuthRoutes(authGroup, authHandlers)
	todoHttp.MapTodoRoutes(todoGroup, todoHandlers)

	return nil
}
