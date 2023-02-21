package main

import "fmt"

func main() {
	var username string = "info"
	fmt.Println(username)
	fmt.Printf("varaibles is of type: %T \n", username)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("varaibles is of type: %T \n", isLogged)

	var smallFloat float32 = 255.1234567
	fmt.Println(smallFloat)
	fmt.Printf("varaibles is of type: %T \n", smallFloat)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("varaibles is of type: %T \n", smallVal)

}
