package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addProfiles(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	w.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addProfiles).Methods("POST")

	log.Fatal(http.ListenAndServe(":5050", router))
}
