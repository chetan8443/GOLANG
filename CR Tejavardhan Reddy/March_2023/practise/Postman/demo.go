package main

import (
	"log"
	"net/http"
)

type server struct{}

func main() {
	S := &server{}
	http.Handle("/", S)
	log.Fatal(http.ListenAndServe(":8887", nil))

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliaction/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"HEllo World":"First API"}`))
}
