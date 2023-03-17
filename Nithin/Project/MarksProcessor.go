package MarksProcessor

import (
	"encoding/csv"
	"fmt"
	b "new/BulkLoad"
	"os"
	"strconv"
)

type Marks struct {
	ID          string
	StudentName string
	Marks       string
}

func IngestRes() {

	res := []string{}

	var db = b.GetMySQLDB()

	marks := []Marks{}

	file, err := os.Open("StudentMarks.csv")
	b.CheckError(err)

	df := csv.NewReader(file)
	data, err := df.ReadAll()
	b.CheckError(err)

	for _, value := range data {
		marks = append(marks, Marks{ID: value[0], StudentName: value[1], Marks: value[2]})
	}

	for j := 0; j < len(marks); j++ {
		mark, err := strconv.ParseInt(marks[j].Marks, 10, 64)
		b.CheckError(err)
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
		_, err = db.Exec("insert into result1(ID, result) values(?,?)", marks[i].ID, res[i])
		b.CheckError(err)
	}
	fmt.Println("All Records Inserted...")
}
