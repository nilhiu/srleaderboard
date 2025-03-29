package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = os.Getenv("JWT_SECRET")

func newJWT(sub string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   sub,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(48 * time.Hour)),
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return tok.SignedString([]byte(jwtSecretKey))
}

// ValidateJWT validates if the given token is a valid JWT token.
func ValidateJWT(token string) (*jwt.Token, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	date, err := tok.Claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if date.Compare(time.Now()) <= 0 {
		return nil, errors.New("token is expired")
	}

	return tok, nil
}
