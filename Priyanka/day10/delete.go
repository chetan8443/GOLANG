package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee represents a single employee
type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Our fake "database" of employees
var employees = []Employee{
	{ID: 1, FirstName: "Priya", LastName: "Panja"},
	{ID: 2, FirstName: "Bobby", LastName: "Vanga"},
	{ID: 3, FirstName: "Aishu", LastName: "Thorry"},
}

// DeleteEmployee removes an employee by their ID
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

// GetEmployees returns a list of all employees
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}
