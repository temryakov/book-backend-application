package main

import (
	"errors"
	"fmt"
	"net/http"

	"snippetapp/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Snippet struct {
	Title string
	Text  string
	gorm.Model
}

var env = config.Get()

func getSnippetByID(id string) (Snippet, error) {
	var dsn = fmt.Sprintf("host=%s user=%s password=abc123 dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", env.DBHost, env.DBUser, env.DBName, env.DBPort)
	var db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var snippet Snippet

	if dbErr != nil {
		return snippet, errors.New("wtf error within db")
	}
	err := db.First(&snippet, id).Error

	if err != nil {
		return snippet, errors.New("Snippet not found")
	}

	return snippet, nil
}

func getSnippetController(c *gin.Context) {
	snippet, err := getSnippetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Snippet not found"})
		return
	}
	c.JSON(http.StatusOK, snippet)
}

func main() {

	gin := gin.Default()

	gin.GET("/api/snippet/:id", getSnippetController)
	gin.Run(env.ServerAddress)
}
