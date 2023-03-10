package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func handlerRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{eid}", GetEmployeeByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", r))
}
