package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Config       *bootstrap.Config
}

func (u *LoginController) Login(c *gin.Context) {
	var request *domain.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}
	if _, err := u.LoginUsecase.FetchByEmail(c, request.Email); err == domain.ErrUserNotFound {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "User not found",
		})
		return
	}
}
