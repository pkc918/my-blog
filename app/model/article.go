package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PWD         string `json:"pwd"`
}
