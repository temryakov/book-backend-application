package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("/api")

	// Public APIs
	NewSignupRouter(config, timeout, db, publicRouter)
	NewLoginRouter(config, timeout, db, publicRouter)
}
