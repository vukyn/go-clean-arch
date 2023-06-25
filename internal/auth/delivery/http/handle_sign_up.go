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

// SignUp godoc
//
//	@Summary		Sign up new user
//	@Description	Sign up new user, returns user and token
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
//	@Success		201			{object}	models.User
//	@Router			/auth/sign-up [post]
func (h *authHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
		}

		createdUser, err := h.authUC.SignUp(ctx, user)
		if err != nil {
			if strings.Contains(err.Error(), constants.ERROR_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpErrors.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
			}
		}
		return c.JSON(http.StatusCreated, createdUser)
	}
}
