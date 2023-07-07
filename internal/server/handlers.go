package server

import (
	_ "boilerplate-clean-arch/docs"
	initAuth "boilerplate-clean-arch/internal/auth/init"
	apiMiddlewares "boilerplate-clean-arch/internal/middleware"
	initTodo "boilerplate-clean-arch/internal/todo/init"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init Auth
	auth := initAuth.NewInit(s.db, s.cfg, s.redisClient)

	// Init middlewares
	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, auth.Repository)

	// Init Todo
	todo := initTodo.NewInit(s.db, s.cfg, mw)

	v1 := e.Group("/api/v1")
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Use(middleware.Logger())

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todo")

	auth.Handler.MapAuthRoutes(authGroup)
	todo.Handler.MapTodoRoutes(todoGroup)

	if s.cfg.Server.Debug {
		log.SetLevel(log.DEBUG)
	}

	return nil
}
