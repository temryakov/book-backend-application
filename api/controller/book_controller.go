package controller

import (
	"net/http"
	"strconv"

	"github.com/temryakov/go-backend-book-app/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	BookUsecase domain.BookUsecase
}

func (u *BookController) FetchByID(c *gin.Context) {

	bookId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	book, err := u.BookUsecase.FetchByID(c, uint(bookId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Book is not found. =(",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.BookDataSerializator{
		Message: "Book is successfully found! %)",
		Data: domain.BookData{
			ID:             book.ID,
			Title:          book.Title,
			Author:         book.Author,
			ChaptersAmount: book.ChaptersAmount,
		},
	})
}

func (u *BookController) Fetch(c *gin.Context) {

	books, err := u.BookUsecase.Fetch(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}
	var arr []domain.BookData

	for _, item := range *books {
		arr = append(arr, domain.BookData{
			ID:             item.ID,
			Title:          item.Title,
			Author:         item.Author,
			ChaptersAmount: item.ChaptersAmount,
		})
	}

	c.JSON(http.StatusOK, domain.BookDataArraySerializator{
		Message: "Book list =P",
		Data:    arr,
	})
}

func (u *BookController) Create(c *gin.Context) {

	var book domain.Book

	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	if err := u.BookUsecase.Create(c, &book); err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Book created successfully! %)",
	})
}

func (u *BookController) Update(c *gin.Context) {

	bookId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}
	var book *domain.Book

	err = c.ShouldBind(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}
	err = u.BookUsecase.Update(c, book, uint(bookId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.BookNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}
	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Book updated successfully! 8-)",
	})
}

func (u *BookController) Delete(c *gin.Context) {

	bookId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}
	err = u.BookUsecase.Delete(c, uint(bookId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.BookNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}
	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Book successfully deleted! :^)",
	})
}
