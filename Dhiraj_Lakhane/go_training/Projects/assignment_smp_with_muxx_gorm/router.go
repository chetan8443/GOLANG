package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func handlerRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/getStu", GetStudentsResult).Methods("GET")
	r.HandleFunc("/getStu/{id}", GetStudentsResult2).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}
