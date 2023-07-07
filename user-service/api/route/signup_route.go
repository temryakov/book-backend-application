package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/user-service/bootstrap"
	"github.com/temryakov/go-backend-book-app/user-service/repository"
	"gorm.io/gorm"
)

func NewSignupRouter(config *bootstrap.Config, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
