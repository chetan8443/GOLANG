package main

import "fmt"

func main() {
	//Write a program that declares a map of strings to integers and prints out the values in reverse order
	a := map[int]string{90: "Dog", 91: "Cat", 92: "Cow", 93: "Bird", 94: "Rabbit"}
	fmt.Println(a)
	var v [10]string
	var k int = 0
	for i, j := range a {
		fmt.Println(i, j)
		v[k] = j
		k = k + 1
	}
	fmt.Println(v)
	//priting the elements in reverse order
	var r [10]string
	l := 0
	var i int
	for i = len(a) - 1; i >= 0; i-- {
		r[l] = v[i]
		l = l + 1
	}
	fmt.Println(r)
	//revrse the values of keys in dictionary
	duplicate := a
	var d int = 0
	for m, _ := range a {
		duplicate[m] = r[d]
		d = d + 1

	}
	fmt.Println(duplicate)

}
