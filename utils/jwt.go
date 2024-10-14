package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type JWTClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(email, username, name string) (string, error) {
	_ = godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	claims := JWTClaims{
		Email:    email,
		Username: username,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
