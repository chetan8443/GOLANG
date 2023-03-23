package MarksProcessor

import (
	"encoding/csv"
	"fmt"
	A "new/BulkLoad"
	"os"
	"strconv"
)

type Marks struct {
	ID          string
	StudentName string
	Marks       string
}

func MarksProcessor() {

	res := []string{} //slice of strings

	var db = A.Connect()

	marks := []Marks{} //slice of struct
	//CSV file accessing
	file, err := os.Open("BulkLoad/StudentMarks.csv")
	if err != nil {
		panic(err)
	}
	//reading all the contents
	df := csv.NewReader(file)
	data, err := df.ReadAll()
	fmt.Println(data)
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		marks = append(marks, Marks{ID: value[0], StudentName: value[1], Marks: value[2]})
	}

	// _, err = db.Exec("create table result1(ID varchar(20), result varchar(20))")
	// b.CheckError(err)

	for j := 0; j < len(marks); j++ {
		mark, err := strconv.ParseInt(marks[j].Marks, 10, 64)
		if err != nil {
			panic(err)
		}
		if mark >= 70 {
			res = append(res, "Pass with Distinction")
		} else {
			if mark < 70 && mark >= 40 {
				res = append(res, "Pass")
			} else {
				res = append(res, "Fail")
			}
		}
	}

	fmt.Println(res)

	db.Query("truncate table result1")

	for i := 0; i < len(marks); i++ {
		fmt.Println(marks[i].ID, res[i])
		// db.Query("ALTER TABLE result1 ADD result varhar(200)")
		_, err = db.Exec("insert into result(ID, result) values(?, ?)", marks[i].ID, res[i])
		if err != nil {
			fmt.Println("Unable to Insert the data:", err)
		}
	}
	fmt.Println("All Records Inserted...")
}
