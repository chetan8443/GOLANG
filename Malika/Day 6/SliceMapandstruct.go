// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Address struct {
	City    string
	PinCode int
	State   string
}

func MapOfStruct() {
	a1 := Address{
		City:    "Jalandhar",
		PinCode: 144001,
		State:   "Punjab",
	}
	a2 := Address{
		City:    "Noida",
		PinCode: 201301,
		State:   "UP",
	}
	a3 := Address{
		City:    "Hyderabad",
		PinCode: 500032,
		State:   "Telengana",
	}

	var nameAddressMap map[string]Address

	nameAddressMap = map[string]Address{"Ram": a1, "Shayam": a2, "Lakshman": a3}
	fmt.Println("Map is:", nameAddressMap)

}

func main() {
	MapOfStruct()
}
