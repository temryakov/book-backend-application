package domain

import "context"

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	FetchByEmail(c context.Context, email string) (*User, error)
	// CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	// CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
