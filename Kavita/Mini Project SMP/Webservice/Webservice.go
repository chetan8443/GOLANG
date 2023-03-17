package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Student struct {
	Studentname string
	Result      string
}

func main() {
	db, err := sql.Open("sqlite3", "../db/mydatabase.db")
	if err != nil {
		log.Fatal("DB error:", err)
	}
	defer db.Close()

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {

		studentid := r.URL.Query().Get("id")
		if studentid == "" {
			http.Error(w, "Missing student id", http.StatusBadRequest)
		}

		tk := Student{}
		row := db.QueryRow("SELECT studentname, result FROM results_table where studentid=?", studentid)
		err = row.Scan(&tk.Studentname, &tk.Result)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "DB Error", http.StatusBadRequest)
		}

		jsonBytes, err := json.Marshal(tk)
		if err != nil {
			http.Error(w, "Json error", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})

	fmt.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
