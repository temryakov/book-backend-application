package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (u *SignupController) Create(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	if err := u.SignupUsecase.Create(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "You have sign up successfully! %)",
	})

}
