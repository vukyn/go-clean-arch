package utils

import (
	"go-clean-arch/internal/auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// JWT Claims struct
type Claims struct {
	Id    int
	Email string
	jwt.StandardClaims
}

// Generate new JWT Token
func GenerateJWTToken(user *models.UserResponse, secret string, ttl int) (string, error) {
	claims := Claims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ComparePasswords(hashedPwd string, plainPwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return err
	}
	return nil
}
