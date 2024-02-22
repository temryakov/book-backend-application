package route

import (
	"time"

	"book-service/bootstrap"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, producer *kafka.Producer, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api/book")
	NewBookRouter(config, db, producer, timeout, publicRouter)
}
