package main

import "github.com/gin-gonic/gin"

type Snippet struct {
	id          int    `json: "id"`
	title       string `json: "title"`
	description string `json: "description"`
	author      string `json: "author"`
}

func main() {
	router := gin.Default()
	// router.GET("/", getList)
	// router.GET("/snippets/:id", getSnippet)
	// router.POST("/snippet", postSnippet)

	router.Run("localhost:8080")
}
