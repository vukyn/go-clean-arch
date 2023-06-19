package http

import (
	"boilerplate-clean-arch/models"
	"boilerplate-clean-arch/pkg/httpErrors"
	"boilerplate-clean-arch/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Register godoc
//	@Summary		Register new user
//	@Description	register new user, returns user and token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		201	{object}	models.User
//	@Router			/auth/register [post]
func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			log.Error(err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := h.authUC.SignUp(ctx, user)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		return c.JSON(http.StatusCreated, createdUser)
	}
}
