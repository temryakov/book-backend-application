package controller

import (
	"net/http"
	"snippetapp/domain"

	"github.com/gin-gonic/gin"
)

type SnippetController struct {
	SnippetUsecase domain.SnippetUsecase
}

func (u *SnippetController) FetchByID(c *gin.Context) {
	userID := c.Param("id")
	tasks, err := u.SnippetUsecase.FetchByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
