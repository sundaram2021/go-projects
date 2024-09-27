package controller

import (
	"bookstore-api/models"
	"bookstore-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	books, err := services.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Create a book
// @Description Add a new book to the store
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 201 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update book details by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Updated Book Data"
// @Success 200 {object} models.Book
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := services.UpdateBook(id, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update book"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Param id path string true "Book ID"
// @Success 200
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete book"})
		return
	}
	c.Status(http.StatusOK)
}
