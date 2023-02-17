package main

import "fmt"

func main() {

	// Area of different entities

	circle := Circle{r: 23}
	rect := Rectangle{h: 25, w: 45}

	fmt.Println("The area of circle:", getArea(circle))
	fmt.Println("The area of rectangle:", getArea(rect))

	// Error Handling in Go

	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
