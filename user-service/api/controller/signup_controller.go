package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (u *SignupController) Create(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadRequest)
		return
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid email address",
		})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid password",
		})
	}

	user.Password = string(encryptedPassword)

	err = u.SignupUsecase.Create(c, &user)

	if err == domain.ErrUserAlreadyExists {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User with this email address already exists. =(",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "You have sign up successfully! %)",
	})

}
