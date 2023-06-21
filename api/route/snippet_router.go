package route

import (
	"snippetapp/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSnippetRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/snippet")
	group.GET("/snippet/:id")
	group.POST("/snippet")
	group.PATCH("/snippet/:id")
	group.DELETE("/snippet/:id")

}
