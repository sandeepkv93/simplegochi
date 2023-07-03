package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sandeepkv93/simplegochi/models"
	"github.com/sandeepkv93/simplegochi/services"

	"github.com/go-chi/chi"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.GetBooks())
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := services.GetBook(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(book)
	}
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	services.CreateBook(book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	err := services.UpdateBook(id, book)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(book)
	}
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := services.DeleteBook(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Book deleted"))
	}
}
