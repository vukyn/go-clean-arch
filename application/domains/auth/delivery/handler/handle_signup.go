package handler

import (
	http_helper "boilerplate-clean-arch/application/utils/httpHelper"
	"boilerplate-clean-arch/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := http_helper.GetRequestCtx(c)
		user := &models.User{}
		if err := http_helper.ReadRequest(c, user); err != nil {
			log.Error(err)
			return c.JSON(http_helper.ErrorResponse(err))
		}

		createdUser, err := h.authUC.SignUp(ctx, user)
		if err != nil {
			return c.JSON(http_helper.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, createdUser)
	}
}
