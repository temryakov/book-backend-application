package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (u *UserController) Fetch(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadRequest)
		return
	}

	user, err := u.UserUsecase.FetchByID(c, uint(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}
	c.JSON(
		http.StatusOK,
		domain.Profile{
			Name:  user.Name,
			Email: user.Email,
		},
	)
}
