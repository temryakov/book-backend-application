package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

func CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	exp := &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(expiry))}
	claims := &domain.JwtCustomClaims{
		Name: user.Name,
		ID:   string(rune(user.Model.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: string(rune(user.Model.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(expiry))},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}
