package controllers

import (
	"book-api-go/database"
	"book-api-go/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// CreateBook Create
func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(book)

	err := db.Create(&book).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Book not created",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// GetBook Get One Book
func GetBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")
	var book models.Book

	err := db.First(&book, "book_id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Book not found",
				"message": fmt.Sprintf("Book with id %s not found", id),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Book not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// GetBooks Get All Books
func GetBooks(ctx *gin.Context) {
	db := database.GetDB()
	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Books not found",
			"message": err.Error(),
		})
		return
	}

	if len(books) <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Books not found",
			"message": "No books found",
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

// UpdateBook Update Book
func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&book).Where("book_id = ?", id).Updates(models.Book{Author: book.Author, Title: book.Title}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Book not updated",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// DeleteBook Delete Book
func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")

	err := db.Where("book_id = ?", id).Delete(&models.Book{}).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   "Book not deleted",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Book deleted successfully")
}
