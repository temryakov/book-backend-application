package route

import (
	"snippetapp/api/controller"
	"snippetapp/bootstrap"
	"snippetapp/domain"
	"snippetapp/repository"
	"snippetapp/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewSnippetRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	sr := repository.NewSnippetRepository(db, domain.CollectionSnippet)
	sc := &controller.SnippetController{
		SnippetUsecase: usecase.NewSnippetUsecase(sr, timeout),
	}

	group.GET("/snippet/:id", sc.FetchByID)

	// group.GET("/snippet", sc.Fetch)
	// group.POST("/snippet")
	// group.PATCH("/snippet/:id")
	// group.DELETE("/snippet/:id")

}
