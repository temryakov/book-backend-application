package controller

import (
	"net/http"
	"strconv"

	"github.com/temryakov/go-backend-snippet-app/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SnippetController struct {
	SnippetUsecase domain.SnippetUsecase
}

func (u *SnippetController) FetchByID(c *gin.Context) {

	snippetId, errType := strconv.ParseUint(c.Param("id"), 0, 16)

	if errType != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	snippet, err := u.SnippetUsecase.FetchByID(c, uint(snippetId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.SnippetNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.FetchByIDSnippetsSerializator{
		Message: "Snippet is successfully found! %)",
		Data: domain.SnippetData{
			ID:    snippet.ID,
			Title: snippet.Title,
			Text:  snippet.Text,
		},
	})
}

func (u *SnippetController) Fetch(c *gin.Context) {

	snippets, err := u.SnippetUsecase.Fetch(c)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.SnippetNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	var arr []domain.SnippetData
	for _, item := range snippets {
		arr = append(arr, domain.SnippetData{
			ID:    item.ID,
			Title: item.Title,
			Text:  item.Text,
		})
	}

	c.JSON(http.StatusOK, domain.FetchSnippetsSerializator{
		Message: "Snippets are successfully found! %)",
		Data:    arr,
	})
}

func (u *SnippetController) Create(c *gin.Context) {

	var snippet domain.Snippet

	err := c.ShouldBind(&snippet)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err = u.SnippetUsecase.Save(c, &snippet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Snippet created successfully! %)",
	})
}

func (u *SnippetController) Update(c *gin.Context) {

	snippetId, errType := strconv.ParseUint(c.Param("id"), 0, 16)

	if errType != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	snippet, err := u.SnippetUsecase.FetchByID(c, uint(snippetId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.SnippetNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	errValidation := c.ShouldBind(&snippet)
	if errValidation != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err = u.SnippetUsecase.Save(c, &snippet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Snippet updated successfully! 8-)",
	})
}

func (u *SnippetController) Delete(c *gin.Context) {

	snippetId, errType := strconv.ParseUint(c.Param("id"), 0, 16)

	if errType != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err := u.SnippetUsecase.Delete(c, uint(snippetId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.SnippetNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Snippet successfully deleted! :^)",
	})
}
