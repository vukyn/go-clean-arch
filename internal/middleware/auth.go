package middleware

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/auth"
	"boilerplate-clean-arch/internal/constants"
	"boilerplate-clean-arch/pkg/httpResponse"
	"errors"
	"net/http"
	"strings"

	"boilerplate-clean-arch/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// JWT way of auth using Authorization header
func (mw *MiddlewareManager) AuthJWTMiddleware(cfg *config.Config, authUC auth.UseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearerHeader := c.Request().Header.Get("Authorization")

			log.Infof("auth middleware bearerHeader %s", bearerHeader)

			if bearerHeader != "" {
				headerParts := strings.Split(bearerHeader, " ")
				if len(headerParts) != 2 {
					log.Errorf("auth middleware: %s", len(headerParts) != 2)
					return c.JSON(http.StatusOK, httpResponse.NewUnauthorizedError(nil))
				}

				tokenString := headerParts[1]

				if err := mw.validateJWTToken(c, cfg, tokenString, authUC); err != nil {
					log.Error("middleware validateJWTToken: %s", err.Error())
					return c.JSON(http.StatusUnauthorized, httpResponse.NewUnauthorizedError(nil))
				}

				return next(c)
			}
			return next(c)
		}
	}
}

func (mw *MiddlewareManager) validateJWTToken(c echo.Context, cfg *config.Config, tokenString string, authUC auth.UseCase) error {
	if tokenString == "" {
		return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Errorf("unexpected signin method %v", token.Header["alg"])
			return nil, utils.NewError(constants.STATUS_CODE_INTERNAL_SERVER, constants.STATUS_MESSAGE_INTERNAL_SERVER_ERROR)
		}
		secret := []byte(cfg.Server.JwtSecretKey)
		return secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(string)
		if !ok {
			return errors.New(constants.STATUS_MESSAGE_INVALID_JWT_TOKEN)
		}

		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return err
		}

		u, err := authUC.GetByID(c.Request().Context(), userUUID)
		if err != nil {
			return err
		}

		c.Set("user", u)
	}
	return nil
}
