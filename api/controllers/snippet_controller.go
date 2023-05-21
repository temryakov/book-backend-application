package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's list"})
}

func Fetch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's snippet"})
}

// func Post(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "Your "})
// }
