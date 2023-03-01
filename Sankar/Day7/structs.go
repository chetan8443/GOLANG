package main

import "fmt"

type EmpContact struct {
	Email string
	Phone int
}

type EmpAddress struct {
	Dno int
	Town  string
	District string
}

type EmpDetails struct{
	Name string
	EmpAddress
	EmpContact
}

func main() {
	e := EmpDetails{Name:"vijay",EmpAddress: EmpAddress{Dno:512,Town:"Bangalore",District:"Bangalore"} , EmpContact: EmpContact{Email:"vijay@gmail.com",Phone:9876540321}}
	fmt.Println(e.EmpAddress)
	fmt.Printf("Employee phone := %v and email := %d",e.EmpContact.Email,e.Phone)
}
