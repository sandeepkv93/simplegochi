package services

import (
	"fmt"

	"github.com/sandeepkv93/simplegochi/models"
)

var books = []models.Book{
	{ID: "1", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"},
}

func GetBooks() []models.Book {
	return books
}

func GetBook(id string) (*models.Book, error) {
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("Book not found")
}

func CreateBook(book models.Book) {
	books = append(books, book)
}

func UpdateBook(id string, newBook models.Book) error {
	for i, b := range books {
		if b.ID == id {
			books[i] = newBook
			return nil
		}
	}
	return fmt.Errorf("Book not found")
}

func DeleteBook(id string) error {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Book not found")
}
