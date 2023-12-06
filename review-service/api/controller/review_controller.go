package controller

import (
	"net/http"
	"strconv"

	"review-service/domain"

	"github.com/gin-gonic/gin"
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
		ReviewID: uint(reviewId),
	}

	review, err := u.ReviewUsecase.FetchReview(c, &query)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.ReviewNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, review)
}

func (u *ReviewController) DeleteReview(c *gin.Context) {

	reviewId, err := strconv.ParseUint(c.Param("id"), 0, 16)

	userId := c.GetUint("x-user-id")

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.MessageBadRequest)
		return
	}

	err = u.ReviewUsecase.DeleteReview(c, uint(reviewId), userId)

	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, domain.ReviewNotFound)
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Review successfully deleted",
	})
}

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

	err := u.ReviewUsecase.CreateReview(c, &review)

	if err == domain.ErrReviewIsExist {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Review already exist!",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.MessageInternalServerError)
		return
	}

	c.JSON(http.StatusOK, domain.SuccessfulMessage{
		Message: "Review created successfully! %)",
	})
}
