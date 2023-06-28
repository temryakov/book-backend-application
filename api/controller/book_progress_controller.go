package controller

import (
	"net/http"
	"strconv"

	"github.com/temryakov/go-backend-book-app/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookProgressController struct {
	BookProgressUsecase domain.BookProgressUsecase
}

func (u *BookProgressController) FetchByID(c *gin.Context) {

	bookProgressId, errType := strconv.ParseUint(c.Param("id"), 0, 16)

	if errType != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	bookProgress, err := u.BookProgressUsecase.FetchByID(c, uint(bookProgressId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Book progress is not found. =(",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.BookProgressSerializator{
		Message:           "Book progress is successfully found! %)",
		CompletedChapters: bookProgress.CompletedChapters,
		IsBookCompleted:   bookProgress.IsBookCompleted,
	})
}
