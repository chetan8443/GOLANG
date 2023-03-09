package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// creating struct for books

type Book struct {
	ID     string
	Title  string
	Author *Author
}

// creating struct for author
type Author struct {
	Firstname string
	Lastname  string
}

// function to get books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

var books []Book

// function to add books
func addBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// function to get books by id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, items := range books {
		if items.ID == params["ID"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// function to update or modify books
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)

}

// function to delete books
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, items := range books {
		if items.ID == params["ID"] {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
