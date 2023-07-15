package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
	"github.com/temryakov/go-backend-book-app/user-service/tokenutil"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Config       *bootstrap.Config
}

func (u *LoginController) Login(c *gin.Context) {
	var request *domain.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	user, err := u.LoginUsecase.FetchByEmail(c, request.Email)
	if err == domain.ErrUserNotFound {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "User not found",
		})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials %/"})
		return
	}
	accessToken, err := tokenutil.CreateAccessToken(user, u.Config.SecretKey, u.Config.ExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
	}
	refreshToken, err := tokenutil.CreateRefreshToken(user, u.Config.SecretKey, u.Config.ExpiryHours)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
