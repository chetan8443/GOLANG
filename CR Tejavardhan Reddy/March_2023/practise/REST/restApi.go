package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	Fname string `json:"Fname"`
	Lname string `json:"Lname"`
	Email string `json:"Email"`
}

type Profile struct {
	Dept        string `json:"dept"`
	Designation string `json:"designation"`
	Employee    User   `json:"Employee"`
}

func additem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.body).Decode(&newProfile)

	w.Header().Set("Content-Type", application/json)

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func main() {
	router := mux.NewRecorder()
	router.HandleFunc("/profiles", additem).Methods("POST")
	http.ListenAndServe(":8089", router)
}
