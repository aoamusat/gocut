package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type UrlRequestBody struct {
	LongUrl string `json:"long_url"`
}

func GenerateRandomString(length uint) string {
	if length == 0 {
		length = 6
	}
	var randomStr string = ""
	var alphabets string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	AlphaList := strings.Split(alphabets, "")
	var i uint
	for i = 0; i < length; i++ {
		randomStr += AlphaList[rand.Intn(len(AlphaList)-1)]
	}

	return randomStr
}

func GenerateShortUrl(LongUrl string, length uint) (string, string) {
	BaseUrl := os.Getenv("BASE_URL")
	ShortCode := GenerateRandomString(length)
	return fmt.Sprintf("%s/%s", BaseUrl, ShortCode), ShortCode
}
