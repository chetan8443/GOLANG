//Creating a slice and modifying its elements using a pointer.

package main

import "fmt"

func main() {

	names := []string{"Nayan", "prathmesh", "chetan", "Arbind"}

	fmt.Println("Before change")
	fmt.Println(names)

	fmt.Println("After change")
	change(&names)
	fmt.Println(names)
}

func change(n *[]string) {
	(*n)[0] = "aniket"
	(*n)[1] = "Meer"

}
