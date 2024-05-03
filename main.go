package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func main() {
	env := os.Getenv("GIN_MODE")

	if env == "release" {
		gin.SetMode("release")
	} else {
		gin.SetMode("debug")
	}

	Router := gin.Default()

	Router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world: " + env})
	})

	Router.Run("localhost:5050")
}
