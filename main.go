package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sandeepkv93/simplegochi/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/books", handlers.GetBooksHandler)
	r.Get("/books/{id}", handlers.GetBookHandler)
	r.Post("/books", handlers.CreateBookHandler)
	r.Put("/books/{id}", handlers.UpdateBookHandler)
	r.Delete("/books/{id}", handlers.DeleteBookHandler)

	fmt.Println("Starting server on port 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Printf("Error starting server: %s", err.Error())
	}
}
