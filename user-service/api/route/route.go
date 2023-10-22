package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/api/middleware"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {

	// Public APIs
	publicRouter := gin.Group("/api")
	NewSignupRouter(config, timeout, db, publicRouter)
	NewLoginRouter(config, timeout, db, publicRouter)
	NewUserRouter(config, timeout, db, publicRouter)

	// Protected APIs
	protectedRouter := gin.Group("/api")
	protectedRouter.Use(middleware.AuthMiddleware(config.SecretKey))
	NewProfileRouter(config, timeout, db, protectedRouter)
}
