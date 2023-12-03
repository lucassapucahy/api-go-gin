package controllers

import (
	initializers "GoGinApi/Initializers"
	"GoGinApi/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookCreate(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadRequest, "failed trying to bind new book "+err.Error())
		return
	}

	if result := initializers.DB.Create(&newBook); result.Error != nil {
		log.Println(result.Error.Error())
		c.String(http.StatusBadGateway, "Unexpected error"+result.Error.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	if result := initializers.DB.Find(&books); result.Error != nil {
		log.Println(result.Error.Error())
		c.String(http.StatusBadGateway, "Unexpected error"+result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}

func GetBooksById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := initializers.DB.Find(&book, id); result.Error != nil {
		log.Println(result.Error.Error())
		c.String(http.StatusBadGateway, "Unexpected error"+result.Error.Error())
		return
	}

	if book.ID == 0 {
		c.String(http.StatusBadRequest, "Id not found")
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var dbBook models.Book
	var requestBook models.Book

	if result := initializers.DB.Find(&dbBook, id); result.Error != nil {
		log.Println(result.Error.Error())
		c.String(http.StatusBadGateway, "Unexpected error"+result.Error.Error())
		return
	}

	if dbBook.ID == 0 {
		c.String(http.StatusBadRequest, "Id not found")
		return
	}

	if err := c.BindJSON(&requestBook); err != nil {
		log.Println(err.Error())
		c.String(http.StatusBadGateway, "Unexpected error while binding object"+err.Error())
		return
	}

	dbBook.Author = requestBook.Author
	dbBook.Pages = requestBook.Pages
	dbBook.Name = requestBook.Name
	dbBook.Description = requestBook.Description

	initializers.DB.Save(&dbBook)

	c.JSON(http.StatusOK, dbBook)
}

func Deletebook(c *gin.Context) {
	id := c.Param("id")

	var bookToDelete models.Book

	if result := initializers.DB.Find(&bookToDelete, id); result.Error != nil {
		log.Println(result.Error.Error())
		c.String(http.StatusBadGateway, "Unexpected error"+result.Error.Error())
		return
	}

	if bookToDelete.ID == 0 {
		c.String(http.StatusBadRequest, "book to delete not found")
		return
	}

	initializers.DB.Delete(&bookToDelete)

	c.JSON(http.StatusOK, bookToDelete)
}
