package controllers

import (
	"book-api-go/database"
	"book-api-go/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var BookDatas []models.Book

// CreateBook Create
func CreateBook(ctx *gin.Context) {
	db := database.GetDB()
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sql := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING *;
	`

	err := db.QueryRow(sql, book.Title, book.Author, book.Description).Scan(&book.BookId, &book.Title, &book.Author, &book.Description)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	defer db.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Book created successfully",
	})
}

// GetBook Get One Book
func GetBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")
	var book models.Book

	sql := `SELECT * FROM books WHERE id = $1;`

	err := db.QueryRow(sql, id).Scan(&book.BookId, &book.Title, &book.Author, &book.Description)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Book not found",
		})
		return
	}
	defer db.Close()

	ctx.JSON(http.StatusOK, book)
}

// GetBooks Get All Books
func GetBooks(ctx *gin.Context) {
	db := database.GetDB()
	var results []models.Book

	sql := `SELECT * FROM books ORDER BY id ASC;`

	rows, err := db.Query(sql)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.Description)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		results = append(results, book)
	}
	defer db.Close()

	if len(results) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "No books found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total": len(results),
		"books": results,
	})
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

	sql := `
	UPDATE books
	SET title = $1, author = $2, description = $3
	WHERE id = $4;
	`

	result, err := db.Exec(sql, book.Title, book.Author, book.Description, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	affected, err := result.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Updated %d book", affected),
	})
}

// DeleteBook Delete Book
func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("id")

	sql := `DELETE FROM books WHERE id = $1;`

	result, err := db.Exec(sql, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	affected, err := result.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "deleted",
		"message": fmt.Sprintf("Deleted %d book", affected),
	})
}
