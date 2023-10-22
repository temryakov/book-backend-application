package route

import (
	"time"

	"book-service/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api/book")
	NewBookRouter(config, db, timeout, publicRouter)
}
