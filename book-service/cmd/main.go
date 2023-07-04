package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temryakov/go-backend-book-app/book-service/api/route"
	"github.com/temryakov/go-backend-book-app/book-service/bootstrap"
)

func main() {
	app := bootstrap.App()

	/* Enviroment variables initializing */
	config := app.Config

	/* Database connecting initializing */
	db := app.DB

	timeout := time.Duration(config.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(config, db, timeout, gin)

	gin.Run(config.ServerAddress)
}
