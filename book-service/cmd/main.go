package main

import (
	"time"

	"book-service/api/route"
	"book-service/bootstrap"

	"github.com/gin-gonic/gin"
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

	gin.Run("localhost:" + config.MicroservicePort)
}
