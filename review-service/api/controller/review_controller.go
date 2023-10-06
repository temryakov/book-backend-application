package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/review-service/domain"
	"gorm.io/gorm"
)

type ReviewController struct {
	ReviewUsecase domain.ReviewUsecase
}

func (u *ReviewController) FetchReview(c *gin.Context) {

	reviewId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	query := domain.ReviewQuery{
		ID: uint(reviewId),
	}

	review, err := u.ReviewUsecase.FetchReview(c, &query)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{
			Message: "Review is not found. =(",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.ReviewResponse{
		BookAuthor:   "Book Author",
		BookTitle:    "Book Title",
		ReviewAuthor: "Review Author",
		Rating:       review.Rating,
		ReviewTitle:  review.Title,
		ReviewText:   review.Text,
	})
}

// func (u *ReviewController) DeleteReview(c *gin.Context) {

// 	userId := c.GetUint("x-user-id")

// 	reviewId, err := strconv.ParseUint(c.Param("id"), 0, 16)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
// 		return
// 	}

// 	err = u.ReviewUsecase.DeleteReview(c, uint(reviewId), userId)

// 	if err == gorm.ErrRecordNotFound {
// 		c.JSON(http.StatusNotFound, domain.ErrorResponse{
// 			Message: "Review is not found. =(",
// 		})
// 		return
// 	}

// 	if err == errors.New("user is not have permission to delete") {
// 		c.JSON(http.StatusForbidden, domain.ErrorResponse{
// 			Message: "You don't have permission to delete this review. %(",
// 		})
// 		return
// 	}

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
// 		return
// 	}

// 	c.JSON(http.StatusOK, domain.SuccessfulMessage{
// 		Message: "Successfully deleted. :=)",
// 	})
// }

func (u *ReviewController) CreateReview(c *gin.Context) {

	userId := c.GetUint("x-user-id")

	var request domain.ReviewRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	review := domain.Review{
		BookId: request.BookId,
		UserId: userId,
		Rating: request.Rating,
		Title:  request.Title,
		Text:   request.Text,
	}

	fetchedReview, err := u.ReviewUsecase.FetchReview(c, &domain.ReviewQuery{
		BookId: request.BookId,
		UserId: userId,
	})

	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	if fetchedReview != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Review already exist!",
		})
		return
	}

	if err := u.ReviewUsecase.CreateReview(c, &review); err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Review created successfully! %)",
	})
}
