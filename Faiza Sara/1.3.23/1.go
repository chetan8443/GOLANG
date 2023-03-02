// program to create a pointer to a struct and accessing its fields

package main

import "fmt"

type person struct {
	name string
	mail string
	age  int
}

func main() {

	per := person{"Hania", "hania@gmail.com", 23}
	var ptr *person = &per
	fmt.Println(ptr)
	fmt.Println(ptr.name, " ", ptr.mail, " ", ptr.age)

}
