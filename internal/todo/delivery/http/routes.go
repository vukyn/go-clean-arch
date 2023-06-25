package http

import (
	"github.com/labstack/echo/v4"

	"boilerplate-clean-arch/internal/todo"
)

// Map todo routes
func MapTodoRoutes(newsGroup *echo.Group, h todo.Handlers) {
	newsGroup.POST("/", h.Create())
	// newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.GET("/:news_id", h.GetByID())
	// newsGroup.GET("/search", h.SearchByTitle())
	// newsGroup.GET("", h.GetNews())
}
