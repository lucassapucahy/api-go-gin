package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Pages       int    `json:"pages" binding:"required"`
	Description string `json:"description" binding:"required"`
}
