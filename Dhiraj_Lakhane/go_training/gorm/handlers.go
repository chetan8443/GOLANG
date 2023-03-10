package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var empId Employee
	json.NewDecoder(r.Body).Decode(&empId)
	Database.First(&empId, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(empId)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

}
