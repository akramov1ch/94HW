package handlers

import (
	"94HW/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetBooks(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []models.Book
		if err := db.Select(&books, "SELECT * FROM books"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
			return
		}
		c.JSON(http.StatusOK, books)
	}
}

func CreateBook(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO books (title, author, published_date, isbn) 
		          VALUES (:title, :author, :published_date, :isbn) 
		          RETURNING id`
		if err := db.QueryRowx(query, book).Scan(&book.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}

		c.JSON(http.StatusCreated, book)
	}
}

func GetBookByID(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		if err := db.Get(&book, "SELECT * FROM books WHERE id=$1", id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}

func UpdateBook(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		id := c.Param("id")
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE books SET title=:title, author=:author, 
		          published_date=:published_date, isbn=:isbn WHERE id=:id`
		bookID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		book.ID = bookID
		if _, err := db.NamedExec(query, book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

func DeleteBook(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := db.Exec("DELETE FROM books WHERE id=$1", id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
	}
}
