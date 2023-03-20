package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	Sid    int
	Sname  string
	Marks  int
	Result string
}

func main() {
	// Connect
	db, err := sql.Open("mysql", "root:Faiza@22@(127.0.0.1:3306)/student")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/students/sid", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["Sid"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var student Student
		err = db.QueryRow("SELECT Sid, Sname, Marks, result FROM studentinfo WHERE sid = ?", id).Scan(&student.Sid, &student.Sname, &student.Marks, &student.Result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}
