package controllers

import (
	"book-api-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var BookDatas []models.Book

// CreateBook Create
func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	book.BookId = len(BookDatas) + 1
	BookDatas = append(BookDatas, book)

	ctx.JSON(http.StatusOK, "Created")
}

// GetBook Get One Book
func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	// convert id string to int
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if bookId > len(BookDatas) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
	} else {
		ctx.JSON(http.StatusOK, BookDatas[bookId-1])
	}
}

// GetBooks Get All Books
func GetBooks(ctx *gin.Context) {
	if len(BookDatas) <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": "Book is empty",
		})
		return
	}

	ctx.JSON(http.StatusOK, BookDatas)
}

// UpdateBook Update Book
func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var bookData models.Book

	if err := ctx.ShouldBindJSON(&bookData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// convert id string to int
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Can't convert id, please check your id",
		})
		return
	}

	if bookId > len(BookDatas) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
	} else {
		BookDatas[bookId-1] = bookData
		BookDatas[bookId-1].BookId = bookId

		ctx.JSON(http.StatusOK, "Updated")
	}
}

// DeleteBook Delete Book
func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	// convert id string to int
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Can't convert id, please check your id",
		})
		return
	}

	if bookId > len(BookDatas) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
		return
	} else {
		BookDatas = append(BookDatas[:bookId-1], BookDatas[bookId:]...)
		ctx.JSON(http.StatusOK, "Deleted")
	}
}
