package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/review-service/bootstrap"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api/review")
	NewReviewRouter(config, db, timeout, publicRouter)
}
