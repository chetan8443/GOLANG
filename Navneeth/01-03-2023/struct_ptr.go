package main

import "fmt"

type student struct {
	name   string
	age    int
	rollno int
}

func main() {
	student := student{
		name:   "Nav",
		age:    21,
		rollno: 8009,
	}

	fmt.Printf("%+v \n", student)

	//changing student name using pointer

	changename(&student)
	fmt.Printf("%+v \n", student)

}

func changename(n *student) {
	n.name = "tony"
}
