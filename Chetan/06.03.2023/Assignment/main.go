// Create a program with a GET API that retrieves data by Id.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Name    string `json:"Name"`
	Id      int    `json:"Id"`
	SurName string `json:"SurName"`
}

var student []Student

func main() {
	student = []Student{
		{Name: "Shreya", Id: 1, SurName: "patil"},
		{Name: "swara", Id: 2, SurName: "patil"},
		{Name: "Jayanta", Id: 3, SurName: "patil"}}

	http.HandleFunc("/getStudent", getStudentById)
	http.HandleFunc("/delete", dStudentById)
	http.HandleFunc("/update", UpdateStudent)
	http.HandleFunc("/getAll", getStudents)
	log.Fatal(http.ListenAndServe(":8082", nil))

}

func getStudents(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(student)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

// func getStudent(w http.ResponseWriter, r *http.Request){

// }

func getStudentById(w http.ResponseWriter, r *http.Request) {

	// keys :=
	// var student Students
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	for i, s := range student {
		if id == s.Id {

			jsonData, _ := json.Marshal(student[i])

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
			return
		}
	}

}

func dStudentById(w http.ResponseWriter, r *http.Request) {



	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	for i, s := range student {
		if id == s.Id {
			student = append(student[:i], student[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "student details have been Removed successfully"})
			w.Write(jsonData)
			return
		}
	}

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var std Student

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	dataBytes, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(dataBytes, &std)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, s := range student {
		if id == s.Id {
			student = append(student[:i], student[i+1:]...)
			student = append(student, std)
			w.Header().Set("Content-Type", "application/json")
			jsonData, _ := json.Marshal(map[string]string{"message": "Student details have been Updated successfully"})
			w.Write(jsonData)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, _ := json.Marshal(map[string]string{"message": "student details not found"})
	w.Write(jsonData)

}
