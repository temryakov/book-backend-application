package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/review-service/api/middleware"
	"github.com/review-service/bootstrap"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api/review")
	NewPublicReviewRouter(config, db, timeout, publicRouter)
	protectedRouter := gin.Group("/api/review")
	protectedRouter.Use(middleware.AuthMiddleware(config.SecretKey))
	NewPrivateReviewRouter(config, db, timeout, protectedRouter)
}
