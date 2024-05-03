package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"olayml.xyz/gocut/database/models"
	"olayml.xyz/gocut/utils"
)

func FetchUrl(ctx *gin.Context) {
	var url models.URL
	if err := models.DbClient.Where("short_code = ?", ctx.Param("short_code")).First(&url).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Resource not found!"})
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, url.LongUrl)
}

func CreateShortUrl(ctx *gin.Context) {
	var RequestBody utils.UrlRequestBody
	if err := ctx.ShouldBindJSON(&RequestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Short URL
	ShortUrl, ShortCode := utils.GenerateShortUrl(RequestBody.LongUrl, 5)

	url := models.URL{
		LongUrl:   RequestBody.LongUrl,
		ShortUrl:  ShortUrl,
		ShortCode: ShortCode,
	}

	if err := models.DbClient.Create(&url).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		log.Printf("%v\n", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": url})
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Service is up..."})
}
