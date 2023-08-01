package route

// func NewBookRouter(cfg *bootstrap.Config, db *gorm.DB, timeout time.Duration, group *gin.RouterGroup) {

// 	rr := repository.NewReviewRepository(db)
// 	rc := &controller.ReviewController{
// 		ReviewUsecase: usecase.NewReviewUsecase(rr, timeout),
// 	}

// 	group.GET("/:id", rc.FetchByID)
// 	group.GET("/all", rc.Fetch)
// 	group.POST("/", rc.Create)
// 	group.DELETE("/:id", rc.Delete)
// 	group.PATCH("/:id", rc.Update)

// }
