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
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	book.BookId = len(BookDatas) + 1
	BookDatas = append(BookDatas, book)

	ctx.JSON(http.StatusOK, "Created")
}

// GetBook Get One Book
func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var book models.Book

	// convert id string to int
	bookId, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for i, book := range BookDatas {
		if book.BookId == bookId {
			condition = true
			book = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
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
	condition := false
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// convert id string to int
	bookId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if book.BookId == bookId {
			condition = true
			BookDatas[i] = book
			BookDatas[i].BookId = bookId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
	}

	ctx.JSON(http.StatusOK, "Updated")
}

// DeleteBook Delete Book
func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false

	for i, book := range BookDatas {
		if book.BookId == book.BookId {
			condition = true
			BookDatas = append(BookDatas[:i], BookDatas[i+1:]...)
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Book not found",
			"message": fmt.Sprintf("Book with id %s not found", id),
		})
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
