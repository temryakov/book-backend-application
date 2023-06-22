package controller

import (
	"net/http"
	"snippetapp/domain"
	"strconv"

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

	snippet, err := u.SnippetUsecase.FetchByID(c, uint16(snippetId))

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.SnippetNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	res := domain.SuccessCreation(&snippet)

	c.JSON(http.StatusOK, res)
}
