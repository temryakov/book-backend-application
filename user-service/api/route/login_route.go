package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/api/controller"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"github.com/temryakov/go-backend-book-app/user-service/repository"
	"github.com/temryakov/go-backend-book-app/user-service/usecase"
	"gorm.io/gorm"
)

func NewLoginRouter(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Config:       config,
	}
	group.POST("/login", lc.Login)
}
