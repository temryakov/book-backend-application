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

	review, err := u.ReviewUsecase.FetchReview(c, uint(reviewId))

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

func (u *ReviewController) DeleteReview(c *gin.Context) {

	reviewId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err := u.ReviewUsecase.DeleteReview(c, uint(reviewId))

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

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Successfully deleted. :=)",
	})
}
