package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/domain"
	review_proto "github.com/temryakov/go-backend-book-app/user-service/proto"
	"google.golang.org/protobuf/proto"
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
	if c.GetHeader("Accept") == "application/x-protobuf" {

		protoUser := &review_proto.GetUserResponse{
			Name: user.Name,
		}

		data, err := proto.Marshal(protoUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize book data"})
			return
		}

		c.Data(http.StatusOK, "application/x-protobuf", data)
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
