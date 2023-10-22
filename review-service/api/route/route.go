package route

import (
	"time"

	"review-service/api/middleware"
	"review-service/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {

	endpointRoute := "/api/review"

	publicRouter := gin.Group(endpointRoute)
	NewPublicReviewRouter(config, db, timeout, publicRouter)

	protectedRouter := gin.Group(endpointRoute)
	protectedRouter.Use(middleware.AuthMiddleware(config.SecretKey))
	NewPrivateReviewRouter(config, db, timeout, protectedRouter)
}
