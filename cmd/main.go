package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Snippet struct {
	Title string
	Text  string
	gorm.Model
}

// Loading enviroment variables
var dbHost = viperEnvVariable("DBHOST")
var dbPort = viperEnvVariable("DBPORT")
var dbUser = viperEnvVariable("DBUSER")
var dbName = viperEnvVariable("DBNAME")

var dsn = fmt.Sprintf("host=%s user=%s password=abc123 dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbName, dbPort)
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func getSnippetByID(id string) (Snippet, error) {

	var snippet Snippet

	err := db.First(&snippet, id).Error

	if err != nil {
		return snippet, errors.New("Snippet not found")
	}

	return snippet, nil
}

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {

	serverAddress := viperEnvVariable("SERVER_ADDRESS")

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
	router.Run(serverAddress)
}
