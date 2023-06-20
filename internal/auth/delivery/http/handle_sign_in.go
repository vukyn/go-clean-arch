package http

import (
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/pkg/httpErrors"
	"boilerplate-clean-arch/pkg/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	// _ "boilerplate-clean-arch/docs"
)

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Sign in and return token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Email		body		string	true	"Email"
//	@Param			Password	body		string	true	"Password"
//	@Success		200			{object}	models.User
//	@Router			/auth/sign-in [post]
func (h *authHandlers) SignIn() echo.HandlerFunc {
	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c echo.Context) error {
		ctx := utils.GetRequestCtx(c)
		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			log.Error(err)
			return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
		}

		userWithToken, err := h.authUC.SignIn(ctx, login.Email, login.Password)
		if err != nil {
			if strings.Contains(err.Error(), constants.ERROR_CODE_BAD_REQUEST) {
				return c.JSON(http.StatusOK, httpErrors.NewBadRequestError(utils.GetErrorMessage(err)))
			} else {
				return c.JSON(http.StatusOK, httpErrors.NewInternalServerError(err))
			}
		}
		return c.JSON(http.StatusOK, userWithToken)
	}
}
