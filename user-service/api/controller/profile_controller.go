package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (p *ProfileController) Fetch(c *gin.Context) {

	userId := c.GetUint("x-user-id")

	user, err := p.ProfileUsecase.FetchByID(c, userId)

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
