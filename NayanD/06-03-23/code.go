package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Students struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	RollNo int    `json:"RollNo"`
}

var students []Students

func main() {
	students = []Students{
		{Id: 1, Name: "Nayan", RollNo: 1},
		{Id: 2, Name: "Arbind", RollNo: 2},
		{Id: 3, Name: "chetan", RollNo: 3},
	}
	http.HandleFunc("/getAll", studentHandler)
	http.HandleFunc("/adddata", studentHandler)
	http.HandleFunc("/delete", studentHandler)
	http.HandleFunc("/update", updateInfo)
	http.HandleFunc("/getById", getByID)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
func getByID(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	var std Students
	InputId, _ := strconv.Atoi(keys[0])

	for i, v := range students {
		if v.Id == InputId {
			std = students[i]
			jsonData, _ := json.Marshal(std)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
	}
}
func getAll(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(students)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func updateInfo(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	var std Students
	InputId, _ := strconv.Atoi(keys[0])

	a, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(a, &std)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, v := range students {
		if v.Id == InputId {
			students = append(students[:i], students[i+1:]...)
			students = append(students, std)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "Student details have been Updated successfully"})
			w.Write(jsonData)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "student details have been added successfully"})
	w.Write(jsonData)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var std Students

	a, er := ioutil.ReadAll(r.Body)

	er = json.Unmarshal(a, &std)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
		fmt.Print(er)
		return
	}

	students = append(students, std)
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "student details have been added successfully"})
	w.Write(jsonData)
}
func studentHandler(w http.ResponseWriter, r *http.Request) {
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
func deleteById(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	InputId, _ := strconv.Atoi(keys[0])

	for i, v := range students {
		if v.Id == InputId {
			students = append(students[:i], students[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "student details have been Removed successfully"})
			w.Write(jsonData)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "student details Not Found"})
	w.Write(jsonData)
	return

}
