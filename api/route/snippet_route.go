package route

import (
	"time"

	"github.com/temryakov/go-backend-snippet-app/api/controller"
	"github.com/temryakov/go-backend-snippet-app/bootstrap"
	"github.com/temryakov/go-backend-snippet-app/domain"
	"github.com/temryakov/go-backend-snippet-app/repository"
	"github.com/temryakov/go-backend-snippet-app/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSnippetRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	sr := repository.NewSnippetRepository(db, domain.CollectionSnippet)
	sc := &controller.SnippetController{
		SnippetUsecase: usecase.NewSnippetUsecase(sr, timeout),
	}

	group.GET("/snippet/:id", sc.FetchByID)
	group.GET("/snippets", sc.Fetch)
	group.POST("/snippet", sc.Create)
	group.DELETE("/snippet/:id", sc.Delete)
	group.PATCH("/snippet/:id", sc.Update)

}
