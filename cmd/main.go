package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
}

var snippets = []Snippet{
	{ID: 1, Title: "First Snippet", Content: "This is the first snippet."},
	{ID: 2, Title: "Second Snippet", Content: "This is the second snippet."},
	{ID: 3, Title: "Third Snippet", Content: "This is the third snippet."},
}

func getSnippetByID(id string) (Snippet, error) {
	for _, snippet := range snippets {
		if strconv.Itoa(snippet.ID) == id {
			return snippet, nil
		}
	}

	return Snippet{}, errors.New("Snippet not found")
}

func main() {
	router := gin.Default()

	router.GET("/api/snippet/:id", func(c *gin.Context) {
		id := c.Param("id")

		snippet, err := getSnippetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Snippet not found"})
			return
		}

		c.JSON(http.StatusOK, snippet)
	})

	router.Run("localhost:8080")
}
