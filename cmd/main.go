package main

import (
	"snippetapp/api/route"
	"snippetapp/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Snippet struct {
	Title string
	Text  string
	gorm.Model
}

func main() {
	app := bootstrap.App()
	config := app.Config
	db := app.DB
	timeout := time.Duration(config.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(config, db, timeout, gin)
	gin.Run(config.ServerAddress)
}

// func getSnippetByID(db *gorm.DB, id string) (Snippet, error) {
// 	var snippet Snippet

// 	err := db.First(&snippet, id).Error

// 	if err != nil {
// 		return snippet, errors.New("Snippet not found")
// 	}

// 	return snippet, nil
// }

// func getSnippetController(c *gin.Context) {
// 	snippet, err := getSnippetByID(db, c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Snippet not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, snippet)
// }
