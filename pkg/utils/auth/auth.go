package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(email string) (string, error) {
	claims := Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenStr string) (*Claims, error) {

	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
