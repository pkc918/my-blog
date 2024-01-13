package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Address     string `json:"address"`
}
