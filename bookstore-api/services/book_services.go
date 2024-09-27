package services

import (
	"bookstore-api/models"
	"errors"
)

var books = []models.Book{
	{ID: "1", Title: "1984", Author: "George Orwell", Price: 9.99},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Price: 7.99},
}

func GetBooks() ([]models.Book, error) {
	return books, nil
}

func CreateBook(book models.Book) error {
	books = append(books, book)
	return nil
}

func UpdateBook(id string, newBook models.Book) error {
	for i, book := range books {
		if book.ID == id {
			books[i] = newBook
			return nil
		}
	}
	return errors.New("Book not found")
}

func DeleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("Book not found")
}
