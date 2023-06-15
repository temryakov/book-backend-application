package main

import (
	"errors"
	"fmt"
	"net/http"

	"snippetapp/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Snippet struct {
	Title string
	Text  string
	gorm.Model
}

// Loading enviroment variables

var dsn = fmt.Sprintf("host=%s user=%s password=abc123 dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DBHost, DBUser, DBName, DBPort)
var db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func getSnippetByID(id string) (Snippet, error) {
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
	app := bootstrap.App()

	env := app.Env

	router := gin.Default()

	router.GET("/api/snippet/:id", getSnippetController)
	router.Run(env.ServerAddress)
}
