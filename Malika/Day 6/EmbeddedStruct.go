// embedded struct example
package main

import "fmt"

type Address struct {
	City    string
	PinCode int
	State   string
}

type UserDetails struct {
	Name      string
	Email     string
	Mobile    string
	MyAddress Address
}

func EmbeddedStruct() {
	var address Address
	var userDetails UserDetails

	address = Address{
		City:    "Jalandhar",
		PinCode: 144001,
		State:   "Punjab",
	}

	userDetails = UserDetails{
		Name:      "Malika",
		Email:     "malika@gmail.com",
		Mobile:    "9999900000",
		MyAddress: address,
	}

	fmt.Println("UserDetails are: ", userDetails)
}

func main() {
	EmbeddedStruct()
}
