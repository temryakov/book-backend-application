package controller

import (
	"net/http"
	"net/mail"

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
	if _, err := mail.ParseAddress(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid email address",
		})
		return
	}
	err := u.SignupUsecase.Create(c, &user)

	if err == domain.ErrUserAlreadyExists {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User with this email address already exists. =(",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "You have sign up successfully! %)",
	})

}
