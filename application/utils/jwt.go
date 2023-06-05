package utils

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT Claims struct
type Claims struct {
	Id             string
	Email          string
	jwt.StandardClaims
}

// Generate new JWT Token
func GenerateJWTToken(user *models.User, config *config.Config) (string, error) {
	claims := Claims{
		Id:    user.UserId.String(),
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Auth.JWTSecret))
}
