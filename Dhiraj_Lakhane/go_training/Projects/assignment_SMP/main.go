package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:dhiraj@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Dummy struct {
	Sid    string
	Name   string
	Marks  string
	Result string
}

var db = getMySQLDB()

func main() {
	// courses := []model.StudentDetails{}
	courses := []Dummy{}

	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var de string
	var res string

	for _, value := range data {
		de = value[2]
		val, _ := strconv.Atoi(de)

		if val > 75 {
			res = "passed with distinction"
		} else if val < 75 && val > 60 {
			res = "passed with grade A"
		} else if val < 60 && val > 50 {
			res = "passed with grade B"
		} else if val < 50 && val > 40 {
			res = "passed with grade C"
		} else {
			res = "failed"
		}
		courses = append(courses, Dummy{Sid: value[0], Name: value[1], Marks: value[2], Result: res})

	}
	for j := 0; j < len(courses); j++ {
		Sid, _ := strconv.Atoi(courses[j].Sid)
		Marks, _ := strconv.Atoi(courses[j].Marks)

		_, err := db.Exec("insert into studentdetail(id, Name, Marks,Result) values(?,?,?,?)", Sid, courses[j].Name, Marks, courses[j].Result)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("all data inserted")

	http.HandleFunc("/get", getDet)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func getDet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("in if")
		keys := r.URL.Query()["Sid"]

		InputId, _ := strconv.Atoi(keys[0])

		var s Dummy
		res1 := fmt.Sprintf("SELECT Name,Marks,Result FROM studentdetail WHERE id=%d", InputId)

		rows, err := db.Query(res1)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&s.Name, &s.Marks, &s.Result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(s)
		}

		jsonD, err := json.Marshal(s)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonD)
	}

}
