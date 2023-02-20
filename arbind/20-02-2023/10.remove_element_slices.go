// How to remove a value from slices based index

package main

import "fmt"

func main() {

	var courses = []string{"java", "reactjs", "python", "Ruby", "swift"} // taking a slice and remove an indexed element

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
