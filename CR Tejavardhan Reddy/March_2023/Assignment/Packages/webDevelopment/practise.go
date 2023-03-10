package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var employee []Employee

func main() {
	fmt.Println("Entering to main function:")
	employee = []Employee{
		{Id: 1001, Name: "TEJA", Age: 21},
		{Id: 1002, Name: "VARDHAN", Age: 22},
		{Id: 1003, Name: "REDDY", Age: 23},
	}
	//Defining the API route
	http.HandleFunc("/getempinfo", employeeHandler)
	http.HandleFunc("/getById", getById)
	http.HandleFunc("/getAll", employeeHandler)
	http.HandleFunc("/update", updateInfo)
	http.HandleFunc("/delete", employeeHandler)

	//Assigning of the port
	log.Fatal(http.ListenAndServe(":8088", nil))
}
func getById(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["Id"]
	fmt.Println(keys)
	var std Employee
	InputId, _ := strconv.Atoi(keys[0])

	for i, v := range employee {
		if v.Id == InputId {
			std = employee[i]
			jsonData, _ := json.Marshal(std)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee details not Found"})
	w.Write(jsonData)
	return
}

func getAll(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(employee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "apllication/json")
	w.Write(jsonData)
}

func updateInfo(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	var std Employee
	InputId, _ := strconv.Atoi(keys[0])

	a, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(a, &std)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, v := range employee {
		if v.Id == InputId {
			employee = append(employee[:i], employee[i+1:]...)
			employee = append(employee, std)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "Employee  details have been Updated successfully"})
			w.Write(jsonData)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee details have been added successfully"})
	w.Write(jsonData)
}
func employeeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAll(w, r)
	case "POST":
		addStudent(w, r)
	case "DELETE":
		deleteById(w, r)
	case "PUT":
		updateInfo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	var std Employee

	a, er := ioutil.ReadAll(r.Body)

	er = json.Unmarshal(a, &std)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		fmt.Print(er)
		return
	}

	employee = append(employee, std)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee  details have been added successfully"})
	w.Write(jsonData)
}
func deleteById(w http.ResponseWriter, r *http.Request) {
	// a, er := ioutil.ReadAll(r.Body)
	keys := r.URL.Query()["id"]
	InputId, _ := strconv.Atoi(keys[0])

	for i, v := range employee {
		if v.Id == InputId {
			employee = append(employee[:i], employee[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "Employee  details have been Removed successfully"})
			w.Write(jsonData)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "Employee details Not Found"})
	w.Write(jsonData)
	return
}
