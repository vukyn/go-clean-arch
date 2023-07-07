package auth

import "github.com/labstack/echo/v4"

// Auth HTTP Handlers interface
type Handlers interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	// Logout() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Delete() echo.HandlerFunc
	// GetUserByID() echo.HandlerFunc
	// FindByName() echo.HandlerFunc
	// GetUsers() echo.HandlerFunc
	// GetMe() echo.HandlerFunc
	// UploadAvatar() echo.HandlerFunc
	// GetCSRFToken() echo.HandlerFunc
}
