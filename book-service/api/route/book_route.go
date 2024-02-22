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

	br := repository.NewBookRepository(db)
	bc := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(br, timeout),
	}

	group.GET("/:id", bc.FetchByID)
	group.GET("/all", bc.Fetch)
	group.POST("/", bc.Create)
	group.DELETE("/:id", bc.Delete)
	group.PATCH("/:id", bc.Update)

}
