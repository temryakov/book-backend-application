package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Snippet struct {
	ID          string
	Name        string
	Description string
	Author      string
}

func getSnippetListController(c *gin.Context) {

	s1 := Snippet{"1", "It's my snippet", "I've just created first snippet in that app!", "Maxim"}
	s2 := Snippet{"2", "Quality", "Quality is not an act, it is a habit.", "Aristotle"}
	s3 := Snippet{"3", "Safety", "Safety is important", "Anonymous"}

	array := []Snippet{s1, s2, s3}

	c.JSON(http.StatusOK, array)
}

func getSnippetController(c *gin.Context) {

	s1 := Snippet{"1", "It's my snippet", "I've just created first snippet in that app!", "Maxim"}
	s2 := Snippet{"2", "Quality", "Quality is not an act, it is a habit.", "Aristotle"}
	s3 := Snippet{"3", "Safety", "Safety is important", "Anonymous"}

	array := []Snippet{s1, s2, s3}

	id := c.Param("id")
	found := false

	for _, snippet := range array {
		if id == snippet.ID {
			c.JSON(http.StatusOK, snippet)
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	}
}

func helloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, it's my snippet application!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/api", helloController)
	router.GET("/api/snippet", getSnippetListController)
	router.GET("/api/snippet/:id", getSnippetController)
	router.Run("localhost:8080")
}
