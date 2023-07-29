package route

import (
	"time"

	"github.com/temryakov/go-backend-book-app/book-service/api/controller"
	"github.com/temryakov/go-backend-book-app/book-service/bootstrap"
	"github.com/temryakov/go-backend-book-app/book-service/repository"
	"github.com/temryakov/go-backend-book-app/book-service/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewBookRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	sr := repository.NewBookRepository(db)
	sc := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(sr, timeout),
	}

	group.GET("/:id", sc.FetchByID)
	group.GET("/all", sc.Fetch)
	group.POST("/", sc.Create)
	group.DELETE("/:id", sc.Delete)
	group.PATCH("/:id", sc.Update)

}
