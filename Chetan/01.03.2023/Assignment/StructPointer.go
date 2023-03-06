package main

import "fmt"

type Student struct {
	roll_no int
	name    string
	branch  string
	batch   int
}

func main() {

	var s1 Student=Student{1,"chetan","IT",32,}
	
	var ptr *Student= &s1

	//printing name before changing
	fmt.Printf("Name: %s\n", (*ptr).name)

	changeWithoutPtr(Student{})
	fmt.Printf("Calling after changeWithoutPtr function\n Name: %s\n", s1.name)

	//call function for changinby passing pointer g name
	changeName(ptr)
	s1.roll_no = 27

	s1.batch = 2019
	fmt.Printf("Calling after changeName with pointer function\n")
	fmt.Printf("Roll Number: %d\n", ptr.roll_no)
	fmt.Printf("Name: %s\n", (*ptr).name)
	fmt.Printf("Branch: %s\n", (*ptr).branch)
	fmt.Printf("Batch: %d", (*ptr).batch)


}
func changeName(pt *Student){
pt.name="Rahul"
}
func changeWithoutPtr(s Student){
s.name="Akash"
}
