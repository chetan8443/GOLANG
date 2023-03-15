package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HEllo World 1")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	http.Handle("/", r)
	//http.HandleFunc("/", handler)
	http.ListenAndServe(":8008", nil)
}
