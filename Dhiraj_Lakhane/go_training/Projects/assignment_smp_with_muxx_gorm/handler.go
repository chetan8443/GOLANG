package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStudentsResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var st []StudentInfo
	Database.Find(&st)
	json.NewEncoder(w).Encode(st)

}
func GetStudentsResult2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var st1 StudentInfo
	json.NewDecoder(r.Body).Decode(&st1)
	Database.First(&st1, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(st1)
}
