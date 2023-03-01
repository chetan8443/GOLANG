// 2.Program to create a slice of and map of structs.

package main

import "fmt"

type fruit struct { // making a struct of fruit
	colour string
	name   string
}

func main() {

	f1 := fruit{"red", "apple"}
	f2 := fruit{"yellow", "banana"}
	f3 := fruit{"orange", "orange"}

	fruits := []fruit{f1, f2, f3} // slice of struct fruit
	fmt.Printf("The slice of fruit is : %+v \nThe type of slice is %T\n", fruits, fruits)

	fruitMap := map[int]fruit{ // create map of struct fruit
		1: f1,
		2: f2,
		3: f3,
	}

	fmt.Println(fruitMap) // map of struct fruit

	for _, v := range fruitMap {
		fmt.Println("Colour :", v.colour, "\tFruit Name :", v.name) //printing fields in the struct
	}

}
