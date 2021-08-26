package controllers

import (
	"api-books/database"
	"api-books/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, id).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book with this ID",
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind params, please, verify your payload",
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book with your payload",
		})
		return
	}

	c.JSON(200, book)

}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books",
		})
		return
	}

	c.JSON(200, books)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind param, please, verify your payload",
		})
		return
	}

	db := database.GetDatabase()

	var books models.Book
	err = db.Where("id = ?", id).Delete(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books",
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))

	var book models.Book

	err = c.ShouldBindJSON(&book)
	book.ID = uint(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind params, please, verify your payload",
		})
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book with your payload",
		})
		return
	}

	c.JSON(200, book)
}
