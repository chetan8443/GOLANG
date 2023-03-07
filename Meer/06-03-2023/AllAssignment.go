package main

import (
	"encoding/json"
	"io/ioutil"
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
		{Id: 1, Name: "Meer", Age: 23},
		{Id: 2, Name: "Fahad", Age: 25},
		{Id: 3, Name: "Ali", Age: 31},
	}

	handleRequests()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/emp", getAllEmp).Methods("GET")
	myRouter.HandleFunc("/emp/{id}", getEmpById).Methods("GET")
	myRouter.HandleFunc("/emp/{id}", deleteEmpById).Methods("DELETE")
	myRouter.HandleFunc("/emp/{id}", updateEmpById).Methods("PUT")
	myRouter.HandleFunc("/emp", addEmp).Methods("POST")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

// function to get all employees
func getAllEmp(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// function to get employee by id
func getEmpById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, emp := range employee {
		if emp.Id == key {
			json.NewEncoder(w).Encode(emp)
		}

	}
}

// function to delete employee by id
func deleteEmpById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for i, emp := range employee {
		if emp.Id == key {
			employee = append(employee[:i], employee[i+1:]...)
			jsonData, _ := json.Marshal(map[string]string{"message": "Employee deleted successfully"})
			w.Write(jsonData)
		}

	}
}

// function to update employee details by id
func updateEmpById(w http.ResponseWriter, r *http.Request) {
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

// function to insert the employee
func addEmp(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	a, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(a, &emp)
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
