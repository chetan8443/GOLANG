package main

import (
	"fmt"
	"sort"
)

func main() {

	var x = []string{"sunday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	sort.Strings(x) // sort function which sort the slice of an array
	fmt.Println(x)

}
