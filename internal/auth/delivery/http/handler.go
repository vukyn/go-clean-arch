package handler

import (
	"go-clean-arch/config"
	"go-clean-arch/internal/auth/models"
	"go-clean-arch/internal/auth/usecase"
	"go-clean-arch/internal/constants"
	"go-clean-arch/pkg/httpResponse"
	"go-clean-arch/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
}

func NewHandler(cfg *config.Config, usecase usecase.IUseCase) IHandler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
	}
}

// Map auth routes
func (h Handler) MapAuthRoutes(authGroup *echo.Group) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	// authGroup.POST("/logout", h.Logout())
	// authGroup.GET("/find", h.FindByName())
	// authGroup.GET("/all", h.GetUsers())
	// authGroup.GET("/:user_id", h.GetUserByID())
	// authGroup.Use(middleware.AuthJWTMiddleware(authUC, cfg))
	// authGroup.GET("/me", h.GetMe())
	// authGroup.GET("/token", h.GetCSRFToken())
	// authGroup.POST("/:user_id/avatar", h.UploadAvatar(), mw.CSRF)
	// authGroup.PUT("/:user_id", h.Update(), mw.OwnerOrAdminMiddleware(), mw.CSRF)
	// authGroup.DELETE("/:user_id", h.Delete(), mw.CSRF, mw.RoleBasedAuthMiddleware([]string{"admin"}))
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login and return token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Success		200			{object}	models.UserWithToken
//	@Router			/auth/login [post]
func (h Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		params := &models.LoginRequest{}
		if err := utils.ReadBodyRequest(c, params); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		userWithToken, err := h.usecase.Login(ctx, params)
		if err != nil {
			return c.JSON(http.StatusOK, httpResponse.ParseError(err))
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusOK, "Success", userWithToken))
	}
}

// Register godoc
//
//	@Summary		Create new user
//	@Description	Create new user, returns user and token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			FirstName	body		string	true	"First name"
//	@Param			LastName	body		string	true	"Last name"
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Param			Gender		body		string	true	"Gender"
//	@Param			City		body		string	false	"City"
//	@Param			Country		body		string	false	"Country"
//	@Param			Birthday	body		string	false	"Gender"
//	@Success		201			{object}	models.UserResponse
//	@Router			/auth/register [post]
func (h Handler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := &models.SaveRequest{}
		if err := utils.ReadBodyRequest(c, user); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}

		createdUser, err := h.usecase.Register(ctx, user)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}
		return c.JSON(http.StatusCreated, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, createdUser))
	}
}
