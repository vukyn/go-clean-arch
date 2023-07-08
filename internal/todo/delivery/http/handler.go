package http

import (
	"boilerplate-clean-arch/config"
	userModel "boilerplate-clean-arch/internal/auth/models"
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/middleware"
	"boilerplate-clean-arch/internal/todo/models"
	"boilerplate-clean-arch/internal/todo/usecase"
	"boilerplate-clean-arch/pkg/httpResponse"
	"boilerplate-clean-arch/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	cfg     *config.Config
	usecase usecase.IUseCase
	mw      *middleware.MiddlewareManager
}

func NewHandler(cfg *config.Config, usecase usecase.IUseCase, mw *middleware.MiddlewareManager) Handler {
	return Handler{
		cfg:     cfg,
		usecase: usecase,
		mw:      mw,
	}
}

// Map todo routes
func (h Handler) MapTodoRoutes(todoGroup *echo.Group) {
	todoGroup.POST("", h.Create(), h.mw.AuthJWTMiddleware())
	// newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	// newsGroup.GET("/:news_id", h.GetByID())
	// newsGroup.GET("/search", h.SearchByTitle())
	// newsGroup.GET("", h.GetNews())
}

// Create godoc
//
//	@Summary		Create todo
//	@Description	Create todo handler
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			Content	body		string	true	"Content"
//	@Success		201		{object}	models.TodoResponse
//	@Router			/todo [post]
func (h Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		todo := &models.SaveRequest{}
		if err := utils.ReadRequest(c, todo); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
		}
		user := c.Get("user").(*userModel.UserResponse)
		createdTodo, err := h.usecase.Create(ctx, user.Id, todo)
		if err != nil {
			if strings.Contains(err.Error(), constants.STATUS_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpResponse.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpResponse.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusOK, httpResponse.NewRestResponse(http.StatusCreated, constants.STATUS_MESSAGE_CREATED, createdTodo))
	}
}
