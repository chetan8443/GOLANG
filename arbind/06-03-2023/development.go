package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var employee []Employee

func main() {
	employee = []Employee{
		{Id: 1, Name: "Ram", Age: 23},
		{Id: 2, Name: "Akbar", Age: 25},
		{Id: 3, Name: "Josh", Age: 31},
	}
	//Defining the API route
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/get_emp", getEmployee).Methods("GET")
	myRouter.HandleFunc("/add_emp", addEmployee).Methods("POST")
	myRouter.HandleFunc("/updateemp", updateEmp).Methods("PUT")
	myRouter.HandleFunc("/getbyid/{id}", getById).Methods("GET")
	myRouter.HandleFunc("/deleteemp/{id}", deleteEmp).Methods("DELETE")

	//Assign port for the server
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
func getEmployee(w http.ResponseWriter, r *http.Request) {
	//Converting the employee slice to json
	jsonData, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func addEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	a, err := io.ReadAll(r.Body)
	json.Unmarshal(a, &emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Adding new data to Employee table
	employee = append(employee, emp)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee details have been added successfully"})
	w.Write(jsonData)
}
func getById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, v := range employee {
		if key == v.Id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	// w.Header().Set("Content-Type", "application/json")
	// jsonData, _ := json.Marshal(map[string]string{"message": "No employee found with this Id !"})
	// w.Write(jsonData)
	jsonData := map[string]string{"messege": "No employee found with this Id !"}
	json.NewEncoder(w).Encode(jsonData)
}

func updateEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for i, emp := range employee {
		if emp.Id == key {
			employee = append(employee[:i], employee[i+1:]...)
		}

	}

	var emp1 Employee
	err := json.NewDecoder(r.Body).Decode(&emp1)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}

	w.Header().Set("Content-Type", "application/json")
	employee = append(employee, emp1)
	err = json.NewEncoder(w).Encode(&emp1)
	if err != nil {
		log.Fatalln("There was an error encoding the initialized struct")
	}
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee details have been updated successfully"})
	w.Write(jsonData)
}

func deleteEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, v := range employee {
		if key == v.Id {
			employee = append(employee[:i], employee[i+1:]...)
			jsonData := map[string]string{"messege": "Employee deleted successfully"}
			json.NewEncoder(w).Encode(jsonData)
			return
		}
	}
	jsonData := map[string]string{"messege": "No employee found with this Id !"}
	json.NewEncoder(w).Encode(jsonData)
}
