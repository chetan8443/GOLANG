package main

import "fmt"
import "math/rand"
//created a struct name emp 
type emp struct {
	name  string
	age   int
	email string
	empNo int
}

func main() {

	e1 := emp{name: "abc", age: 23, email: "abc@gmail.com"}

	ptr := &e1 // declaring pointer and assining address of e1 to it 
	var name string
	var age int 
	var mail string
    fmt.Println("Register employee information") //taking all information from console
	fmt.Printf("Enter name:")
	fmt.Scanln(&name)
	fmt.Printf("Enter age :")
	fmt.Scanln(&age)
	fmt.Printf("Enter email:")
	fmt.Scanln(&mail)
	// saving info of emp using pointer in e1 
	(*ptr).name=name
	(*ptr).email=mail
	(*ptr).age = age
	(*&ptr.empNo)=rand.Int()
// showing all info of emp accessing that info using struct
	fmt.Println("Registation succesfull")
	fmt.Println("============================================")
	fmt.Println("Registation Details")
	fmt.Println("Name of Employee : ",(*&ptr.name))
	fmt.Println("Age of Employee : ",(*&ptr.age))
	fmt.Println("Email of Employee : ",(*&ptr.email))
	fmt.Println("Email of Employee : ",(*&ptr.empNo))


}
