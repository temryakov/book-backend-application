package route

import (
	"time"

	"github.com/temryakov/go-backend-snippet-app/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	NewSnippetRouter(config, db, timeout, publicRouter)
}
