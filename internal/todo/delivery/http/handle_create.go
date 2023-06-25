package http

import (
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/internal/models"
	"boilerplate-clean-arch/pkg/httpErrors"
	"boilerplate-clean-arch/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Create godoc
//	@Summary		Create todo
//	@Description	Create todo handler
//	@Tags			Todo
//	@Accept			json
//	@Produce		json
//	@Param			Content	body		string	true	"Content"
//	@Success		201		{object}	models.Todo
//	@Router			/todo [post]
func (t todoHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		todo := &models.Todo{}
		if err := utils.ReadRequest(c, todo); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
		}

		createdTodo, err := t.todoUC.Create(ctx, todo)
		if err != nil {
			if strings.Contains(err.Error(), constants.ERROR_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpErrors.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
			}
		}

		return c.JSON(http.StatusCreated, createdTodo)
	}
}
