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

	res := domain.FetchByIdSuccess(&snippet)

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusOK, snippets)
}

func (u *SnippetController) Create(c *gin.Context) {

	var snippet domain.Snippet

	err := c.ShouldBind(&snippet)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err = u.SnippetUsecase.Create(c, &snippet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Snippet created successfully. %)",
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

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Snippet successfully deleted. :^)",
	})
}
