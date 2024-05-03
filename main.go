package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"olayml.xyz/gocut/controllers"
	"olayml.xyz/gocut/database/models"
)

var _ = godotenv.Load()

func main() {
	models.ConnectDatabase()

	Router := gin.Default()

	Router.GET("/ping", controllers.Ping)
	Router.GET("/:short_code", controllers.FetchUrl)
	Router.POST("/create", controllers.CreateShortUrl)

	Router.Run("localhost:" + os.Getenv("PORT"))
}
