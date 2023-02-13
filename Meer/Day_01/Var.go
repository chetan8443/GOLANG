package main

import "fmt"

// jwtToken :=300000 No Var Style is not allowed outside the functions

const LoginToken string = "ABCD" //public bcoz name starts with cap letters

func main() {
	fmt.Println("Variables ")
	fmt.Println()
	var username string = "Meer"
	fmt.Println("Hello " + username)
	fmt.Printf("Variable is of type : %T \n", username)
	fmt.Println()
	var isLoggedIn bool = true
	fmt.Println("Meer is Logged In : ", isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)
	fmt.Println()

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	fmt.Println()

	var smallFloat float32 = 255.454545455 // prints only five decimal points
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)
	fmt.Println()

	var anotherVar int // prints default value of int var i.e. 0
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type : %T \n", anotherVar)
	fmt.Println()

	//implicite type

	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)
	fmt.Println()

	//No Var style valid only inside the function
	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)
	fmt.Println()

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
