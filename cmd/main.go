package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/", getListController.getList)
	// router.GET("/snippets/:id", getSnippet)
	// router.POST("/snippet", postSnippet)

	// log.Println("Run server on", env.serverAddress)

	router.Run("localhost:8080")
}
