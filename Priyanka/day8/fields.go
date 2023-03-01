package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := &Person{
		Name: "Priya",
		Age:  28,
	}

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)

}
