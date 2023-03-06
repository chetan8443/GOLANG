package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	http.HandleFunc("/getempdata", employeeHandler)
	//Assign port for the server
	log.Fatal(http.ListenAndServe(":8081", nil))

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
func employeeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getEmployee(w, r)
	case "POST":
		addEmployee(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
