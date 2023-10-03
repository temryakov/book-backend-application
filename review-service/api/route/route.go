package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/review-service/api/middleware"
	"github.com/review-service/bootstrap"
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
