//Creating a pointer to a struct and accessing its fields

package main

import "fmt"

type person struct { // declaring struct
	Name string
	age  int
}

func main() {
	person_info := person{"prathmesh", 23} // initializing the struct

	pointer_info := &person_info // storing memory address of person struct

	//Accessing elements of a struct
	fmt.Printf("Nmae: %s, Age: %d\n", pointer_info.Name, pointer_info.age)
	changeValue(&person_info)
	fmt.Println(person_info)

}

func changeValue(v *person) { // creating the pointer function which changes the value of struct element by the pointer
	v.age = 22

}
