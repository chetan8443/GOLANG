package main

import "fmt"

type students struct {
	name   string
	ID     int
	age    int
	branch branch
}

type branch struct {
	electronics string
	comp        string
	civil       string
}

func main() {
	var s students = students{
		name: "ram",
		ID:   100,
		age:  17,
		branch: branch{
			electronics: "yes",
			comp:        "no",
			civil:       "no",
		},
	}

	fmt.Println("student for electronics branch selection:", s.name, s.ID, s.branch.electronics)

}
