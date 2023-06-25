package route

import (
	"time"

	"github.com/temryakov/go-backend-book-app/api/controller"
	"github.com/temryakov/go-backend-book-app/bootstrap"
	"github.com/temryakov/go-backend-book-app/domain"
	"github.com/temryakov/go-backend-book-app/repository"
	"github.com/temryakov/go-backend-book-app/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewBookRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	sr := repository.NewBookRepository(db, domain.CollectionBook)
	sc := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(sr, timeout),
	}

	group.GET("/book/:id", sc.FetchByID)
	group.GET("/books", sc.Fetch)
	group.POST("/book", sc.Create)
	group.DELETE("/book/:id", sc.Delete)
	group.PATCH("/book/:id", sc.Update)

}
