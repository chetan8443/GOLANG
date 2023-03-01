package main

import "fmt"

type things struct {
	color string
	name  string
}

func main() {
	b1 := things{"blue", "bat"}
	b2 := things{"red", "ball"}
	b3 := things{"yellow", "cup"}

	thing := []things{b1, b2, b3} //slice of struct things
	fmt.Printf("The slice of things is : %+v \nThe type of slice is %T\n", thing, thing)

	thingsMap := map[int]things{ // craeting map of struct things
		1: b1,
		2: b2,
		3: b3,
	}
	fmt.Println(thingsMap)

}
