package http

import (
	"github.com/labstack/echo/v4"

	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
	"boilerplate-clean-arch/internal/middleware"
	"boilerplate-clean-arch/internal/todo"
)

// Map todo routes
func MapTodoRoutes(todoGroup *echo.Group, h todo.Handlers, cfg *config.Config, authUC auth.UseCase, mw *middleware.MiddlewareManager) {
	todoGroup.POST("", h.Create(), mw.AuthJWTMiddleware(cfg, authUC))
	// newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.GET("/:news_id", h.GetByID())
	// newsGroup.GET("/search", h.SearchByTitle())
	// newsGroup.GET("", h.GetNews())
}
