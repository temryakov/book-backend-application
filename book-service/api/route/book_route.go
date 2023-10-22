package route

import (
	"time"

	"book-service/api/controller"
	"book-service/bootstrap"
	"book-service/repository"
	"book-service/usecase"

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
