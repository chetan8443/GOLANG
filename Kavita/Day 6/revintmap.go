package main

import "fmt"

func main() {

	mymap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	var values []int

	for _, value := range mymap {
		values = append(values, value)
	}

	for i := len(values) - 1; i >= 0; i-- {
		fmt.Println(values[i])
	}
}
