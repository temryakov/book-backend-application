package route

import (
	"time"

	"review-service/api/controller"
	"review-service/bootstrap"
	"review-service/repository"
	"review-service/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewPublicReviewRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	rr := repository.NewReviewRepository(db)
	rc := &controller.ReviewController{
		ReviewUsecase: usecase.NewReviewUsecase(rr, timeout),
	}

	group.GET("/:id", rc.FetchReview)
	// group.GET("/all", rc.FetchAllReview)
}

func NewPrivateReviewRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

	rr := repository.NewReviewRepository(db)
	rc := &controller.ReviewController{
		ReviewUsecase: usecase.NewReviewUsecase(rr, timeout),
	}

	group.POST("/", rc.CreateReview)
	// group.DELETE("/:id", rc.DeleteReview)
	// group.PATCH("/:id", rc.UpdateReview)
}
