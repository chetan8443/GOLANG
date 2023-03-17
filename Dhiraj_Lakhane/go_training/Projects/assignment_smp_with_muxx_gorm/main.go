package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:dhiraj@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	DataHandler()

	var db = getMySQLDB()
	DummySlice := []StudentInfo{}

	file, err := os.Open("StudentMarks.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range data {
		log.Println(v)
	}

	var de string
	var res string

	for _, value := range data {

		// fmt.Println(v)

		// fmt.Printf("var1 = %T\n", value[i])
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
		DummySlice = append(DummySlice, StudentInfo{Id: value[0], Name: value[1], Marks: value[2], Result: res})
	}
	for j := 0; j < len(DummySlice); j++ {
		Sid, _ := strconv.Atoi(DummySlice[j].Id)
		Marks, _ := strconv.Atoi(DummySlice[j].Marks)

		_, err := db.Exec("insert into student_infos(id, Name, Marks,Result) values(?,?,?,?)", Sid, DummySlice[j].Name, Marks, DummySlice[j].Result)
		if err != nil {
			log.Fatal(err)
		}
	}
	handlerRouter()

}
