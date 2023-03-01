package main

import "fmt"

type Contact struct {
	Email string
	Phone string
}

type Person struct {
	Name string
	Age  int
	Contact
}

func main() {
	p := Person{Name: "priya", Age: 28, Contact: Contact{Email: "priyanka@gmail.com", Phone: "123456789"}}
	fmt.Printf("%s is %d years old and can be contacted at %s or %s.\n", p.Name, p.Age, p.Email, p.Phone)
}
