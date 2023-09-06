package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/review-service/api/controller"
	"github.com/review-service/bootstrap"
	"github.com/review-service/repository"
	"github.com/review-service/usecase"
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

	// group.POST("/:id", rc.CreateReview)
	group.DELETE("/:id", rc.DeleteReview)
	// group.PATCH("/:id", rc.UpdateReview)
}
