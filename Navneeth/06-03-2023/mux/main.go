package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//adding values to books
	books = append(books, Book{
		ID:    "1",
		Title: "Book One",
		Author: &Author{
			Firstname: "Navneeth",
			Lastname:  "Prasanth",
		},
	})

	books = append(books, Book{
		ID:    "2",
		Title: "Book Two",
		Author: &Author{
			Firstname: "go",
			Lastname:  "lang",
		},
	})

	r := mux.NewRouter() //new router

	//route handlers
	r.HandleFunc("/api/Books", GetBooks).Methods("GET")
	r.HandleFunc("/api/Book/{ID}", GetBook).Methods("GET")
	r.HandleFunc("/api/Books", addBooks).Methods("POST")
	r.HandleFunc("/api/Book/{ID}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/Book/{ID}", deleteBooks).Methods("DELETE")

	http.ListenAndServe(":8000", r)

}
