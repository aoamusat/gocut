package models

import "github.com/jinzhu/gorm"

type URL struct {
	gorm.Model
	ShortUrl  string `json:"short_url" gorm:"unique"`
	LongUrl   string `json:"long_url"`
	ShortCode string `json:"short_code" gorm:"unique"`
}
